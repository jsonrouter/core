package tree

import (
	"fmt"
	"reflect"
	"strings"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/openapi/v3"
)

// Required applies model which describes required request payload fields
func (handler *Handler) Required(objects ...validation.Payload) *Handler {
	for _, object := range objects {
		handler.updateSpecParams(true, object)
	}
	return handler
}

// Optional applies model which describes optional request payload fields
func (handler *Handler) Optional(objects ...validation.Payload) *Handler {
	for _, object := range objects {
		handler.updateSpecParams(false, object)
	}
	return handler
}

// updateSpecParams triggers an update of the spec parameters
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

		pathString := handler.Node.FullPath()
		path := spec.Paths[pathString]
		operation := path[strings.ToLower(handler.Method)]
		//ref := fmt.Sprintf("#/components/requestBodies/%v", pathString[1:])

		if spec.Components.RequestBodies[pathString] == nil {
			spec.Components.RequestBodies[pathString] = &openapiv3.RequestBody{
				Required: required,
				Description: handler.Descr,
				Content: map[string]*openapiv3.MediaType{
					"application/json": &openapiv3.MediaType{
						Schema: &openapiv3.Schema{
							Properties: map[string]*openapiv3.Schema{},
						},
					},
				},
			}
		}

		for k, v := range payload {
			handler.updateSpecParam(required, spec.Components.RequestBodies[pathString], k, v)
		}

		// only create the definition ONCE if it has contents
		if !handler.spec.addedBodyDefinition {
			if len(spec.Components.RequestBodies[pathString].Content["application/json"].Schema.Properties) > 0 {
				operation.RequestBody = spec.Components.RequestBodies[pathString]
				//operation.RequestBody.Ref = ref
			}
			handler.spec.addedBodyDefinition = true
		}

	default: panic("INVALID SPEC TYPE")
	}

}

// updateSpecParam triggers and updates the spec parameter.
func (handler *Handler) updateSpecParam(required bool, def interface{}, key string, cfg *validation.Config) {

	switch v := handler.payloadSchema.(type) {
		case nil:

			m := validation.Payload{}
			m[key] = cfg
			handler.payloadSchema = m

		case validation.Payload:

			v[key] = cfg
			handler.payloadSchema = v

		default:

			panic("INVALID PAYLOAD TYPE: "+reflect.TypeOf(handler.payloadSchema).String())

	}

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
		if param.Type == "array" {
			param.Items = map[string]string{
				"type": "string",
			}
		}
//		param.Required = required

		if required == true {
			definition.Required = append(
				definition.Required,
				key,
			)
		}

		definition.Properties[key] = param

	case *openapiv3.RequestBody:

		obj := &openapiv3.Schema{}

		obj.Description = cfg.DescriptionValue
		obj.Minimum = int(cfg.Min)
		obj.Maximum = int(cfg.Max)
		obj.Default = fmt.Sprintf("%v", cfg.DefaultValue)
		obj.Format = cfg.Type
		obj.Type = openapiv3.Type(cfg.Model)
		obj.Required = required

		definition.Content["application/json"].Schema.Properties[key] = obj

		default: panic("INVALID SPEC TYPE")
	}

}

// updateParameters adds route params to the spec
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
			param.Type = openapiv2.Type(cfg.Type)
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
