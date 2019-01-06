package tree

import (
	"fmt"
	"strings"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/openapi/v2"
//	"github.com/jsonrouter/core/openapi/v3"
)

// Applies model which describes required request payload fields
func (handler *Handler) Required(objects ...validation.Payload) *Handler {
	for _, object := range objects {
		handler.updateSpecParams(true, object)
	}
	return handler
}

// Applies model which describes optional request payload fields
func (handler *Handler) Optional(objects ...validation.Payload) *Handler {
	for _, object := range objects {
		handler.updateSpecParams(false, object)
	}
	return handler
}

func (handler *Handler) updateSpecParams(required bool, payload validation.Payload) {

	switch spec := handler.Node.Config.Spec.(type) {
	case *openapiv2.Spec:

		path := spec.Paths[handler.Path(spec.BasePath)]
		pathMethod := path[strings.ToLower(handler.Method)]
		ref := handler.Path(spec.BasePath)
		ref = strings.Replace(ref, "/", "-", -1)
		ref = strings.Replace(ref, "{", "", -1)
		ref = strings.Replace(ref, "}", "", -1)
		ref = fmt.Sprintf(
			"#/definitions/%s-%s",
			handler.Method,
			ref,
		)

		if spec.Definitions[ref] == nil {
			spec.Definitions[ref] = &openapiv2.Definition{
				Type: "object",
				Properties: map[string]openapiv2.Parameter{},
			}
		}

		for k, v := range payload {
			handler.updateSpecParam(required, spec.Definitions[ref], k, v)
		}

		// only create the definition ONCE if it has contents
		if handler.spec.addedBodyDefinition {
			if len(spec.Definitions[ref].Properties) > 0 {
				pathMethod.Parameters = append(
					pathMethod.Parameters,
					&openapiv2.Parameter{
	//					Required: true,
						Name: "body",
						In: "body",
						Description: handler.Descr,
						Schema: &openapiv2.Schema{
							Ref: ref,
						},
					},
				)
			}
			handler.spec.addedBodyDefinition = true
		}

	default: panic("INVALID SPEC TYPE")
	}

}

func (handler *Handler) updateSpecParam(required bool, def interface{}, key string, cfg *validation.Config) {

	pointerFloat64 := func(f float64) *float64 {
		if f == 0 {
			return nil
		}
		return &f
	}

	switch definition := def.(type) {
	case *openapiv2.Definition:

		param := openapiv2.Parameter{}
		param.Description = cfg.DescriptionValue
		param.Minimum = pointerFloat64(cfg.Min)
		param.Maximum = pointerFloat64(cfg.Max)
		param.Default = cfg.DefaultValue
		param.Format = cfg.Type
//		param.Required = required

		if required == true {
			definition.Required = append(
				definition.Required,
				key,
			)
		}

		switch param.Format {
		case "bool":
			param.Type = "boolean"
		case "string":
			param.Type = "string"
		case "int64", "int":
			param.Type = "integer"
		case "float32", "float64":
			param.Type = "number"
		}

		definition.Properties[key] = param

	default: panic("INVALID SPEC TYPE")
	}

}

// Adds route params to the spec
func (handler * Handler) updateParameters() {

	switch spec := handler.Node.Config.Spec.(type) {
	case *openapiv2.Spec:

		path := spec.Paths[handler.Path(spec.BasePath)]
		pathMethod := path[strings.ToLower(handler.Method)]

		for _, cfg := range handler.Node.Validations {
			param := &openapiv2.Parameter{}
			param.In = "path"
			param.Name = cfg.Keys[0]
			param.Description = cfg.DescriptionValue
			param.Type = cfg.Type
			minLength := int64(cfg.Min)
			maxLength := int64(cfg.Max)
			param.MinLength = &minLength
			param.MaxLength = &maxLength
			param.Required = true

			pathMethod.Parameters = append(pathMethod.Parameters, param)
		}
	}
}
