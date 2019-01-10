package tree

import (
	"fmt"
	"strings"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/openapi/v3"
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
		def := handler.Ref(spec.BasePath)
		ref := fmt.Sprintf("#/definitions/%s", def)

		if spec.Definitions[def] == nil {
			spec.Definitions[def] = &openapiv2.Definition{
				Type: "object",
				Properties: map[string]openapiv2.Parameter{},
			}
		}

		for k, v := range payload {
			handler.updateSpecParam(required, spec.Definitions[def], k, v)
		}

		// only create the definition ONCE if it has contents
		if !handler.spec.addedBodyDefinition {
			if len(spec.Definitions[def].Properties) > 0 {
				pathMethod.Parameters = append(
					pathMethod.Parameters,
					&openapiv2.Parameter{
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

	case *openapiv3.Spec:
/*
		path := spec.Paths[handler.Node.FullPath()]
		pathMethod := path[strings.ToLower(handler.Method)]
		ref := fmt.Sprintf("#/components/requestBodies/%s", pathMethod)

		if spec.Components.RequestBodies[path] == nil {
			spec.Components.RequestBodies[path] = &openapiv3.RequestBody{
				Required: required,
				Description: handler.Descr,
				Content: map[string]*MediaType{},
			}
		}

		for k, v := range payload {
			handler.updateSpecParam(required, spec.Components.BodyObjects[ref], k, v)
		}

		// only create the definition ONCE if it has contents
		if !handler.spec.addedBodyDefinition {
			if len(spec.Components.RequestBodies[ref].Content) > 0 {
				pathMethod.Parameters = append(
					pathMethod.Parameters,
					&openapiv3.Parameter{
						Name: "body",
						In: "body",
						Description: handler.Descr,
						Schema: &openapiv3.Schema{
							Ref: ref,
						},
					},
				)
			}
			handler.spec.addedBodyDefinition = true
		}
*/
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
		param.Type = openapiv3.Type(cfg.Model)
//		param.Required = required

		if required == true {
			definition.Required = append(
				definition.Required,
				key,
			)
		}

		definition.Properties[key] = param

	case *openapiv3.RequestBody:
/*
		param := openapiv3.Parameter{}

		param.Description = cfg.DescriptionValue
		//param.Minimum = pointerFloat64(cfg.Min)
		//param.Maximum = pointerFloat64(cfg.Max)
		param.Default = cfg.DefaultValue
		param.Format = cfg.Type
		param.Type = openapiv3.Type(cfg.Model)
		param.Required = required

		definition.Content[key] = param
*/
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
	case *openapiv3.Spec:

		path := spec.Paths[handler.Node.FullPath()]
		operation := path[strings.ToLower(handler.Method)]

		for _, cfg := range handler.Node.Validations {
			param := &openapiv3.Parameter{}
			param.In = "path"
			param.Name = cfg.Keys[0]
			param.Description = cfg.DescriptionValue
			minLength := int(cfg.Min)
			maxLength := int(cfg.Max)
			param.Schema = &openapiv3.Schema{}
			param.Schema.Type = cfg.Type
			param.Schema.Minimum = minLength
			param.Schema.Maximum = maxLength
			
			param.Required = true

			operation.Parameters = append(operation.Parameters, param)
		}
		default: panic("INVALID SPEC TYPE")
	}
}
