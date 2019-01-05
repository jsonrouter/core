package openapi

import (
	"fmt"
	"strings"
	"net/url"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/openapi/v2"
)

// Builds the handler documentation object.
func BuildV2(spec *openapi.APISpec, handlerMap map[string]*tree.Handler) {

	handlers := map[*tree.Handler]*tree.HandlerSpec{}

	// do the legacy spec operation
	for _, handler := range handlerMap {
		handlers[handler] = handler.Spec()
	}

	swaggerPath := "/swagger.json"
	_, ok := spec.Paths[swaggerPath]
	if !ok {
		spec.Paths[swaggerPath] = &openapi.Path{}
	}
	spec.Paths[swaggerPath].GET = &openapi.PathMethod{
		Produces: []string{
			"application/json",
		},
		Description: "Serves the OpenAPI spec JSON",
		Responses: openapi.Responses{
			Code200: &openapi.StatusCode{
				Description: "Done OK",
				Schema: openapi.StatusSchema{
					Type: "object",
				},
			},
		},
	}

	for handler, handlerSpec := range handlers {

		// paths need to be RFC 3986 path encoded
		k := url.PathEscape(
			strings.Replace(
				fmt.Sprintf("%s-%s", handler.Node.FullPath(), handler.Method),
				"/",
				"_",
				-1,
			),
		)

		// create the object that holds the handler's definition
		pathMethod := &openapi.PathMethod{
			OperationID: k,
			Produces: []string{
				"application/json",
			},
			Parameters: []*openapi.Parameter{},
			Description: handler.Description,
			Responses: openapi.Responses{
				Code200: &openapi.StatusCode{
					Description: "Done OK",
					Schema: openapi.StatusSchema{
						Type: "object",
					},
				},
			},
		}

		definition := &openapi.Definition{
			Type: "object",
			Properties: map[string]openapi.Parameter{},
		}

		// check the type of payload
		switch t := handlerSpec.PayloadSchema.(type) {
			case map[string]*validation.Config:

				// dont make a param object for GET routes etc
				if len(t) == 0 { continue }

				pathMethod.Parameters = append(
					pathMethod.Parameters,
					&openapi.Parameter{
//						Required: true,
						Name: "body",
						In: "body",
						Description: handler.Description,
						Schema: &openapi.Schema{
							Ref: "#/definitions/" + k,
						},
					},
				)

				requiredBodyParams := []string{}

				// add body params to the definition properties
				for key, cfg := range t {

					requiredBodyParams = append(
						requiredBodyParams,
						key,
					)

					param := openapi.Parameter{}
					param.Description = cfg.DescriptionValue
					param.Minimum = pointerFloat64(cfg.Min)
					param.Maximum = pointerFloat64(cfg.Max)
					param.Default = cfg.DefaultValue
					param.Format = cfg.Type
//					param.Required = cfg.RequiredValue

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

				}

				definition.Required = requiredBodyParams

		}

		// only create the definition if it has contents
		if len(definition.Properties) > 0 {
			spec.Definitions[k] = definition
		}

		// route params
		for name, cfg := range handlerSpec.RouteParams {
			param := &openapi.Parameter{}
			param.In = "path"
			param.Name = name
			param.Description = cfg.DescriptionValue
			param.Type = cfg.Type
			minLength := int64(cfg.Min)
			maxLength := int64(cfg.Max)
			param.MinLength = &minLength
			param.MaxLength = &maxLength

			pathMethod.Parameters = append(pathMethod.Parameters, param)
		}

		fullPath := handler.Node.FullPath()
		j := strings.Split(fullPath, "/")
		if len(j) >= 3 {
			fullPath = "/" + strings.Join(j[2:], "/")
		}
		_, ok := spec.Paths[fullPath]
		if !ok {
			spec.Paths[fullPath] = &openapi.Path{}
		}

		var f *openapi.Path

		switch handler.Method {
		case "GET":
			f = spec.Paths[fullPath]
			f.GET = pathMethod
		case "POST":
			f = spec.Paths[fullPath]
			f.POST = pathMethod
		case "PUT":
			f = spec.Paths[fullPath]
			f.PUT = pathMethod
		case "PATCH":
			f = spec.Paths[fullPath]
			f.PATCH = pathMethod
		case "DELETE":
			f = spec.Paths[fullPath]
			f.DELETE = pathMethod
		case "HEAD":
			f = spec.Paths[fullPath]
			f.HEAD = pathMethod
		case "OPTIONS":
			f = spec.Paths[fullPath]
			f.OPTIONS = pathMethod
		default:
			panic("wtf")
		}

		fmt.Println(fullPath, handler.Method, f)
	}

}
