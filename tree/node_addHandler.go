package tree

import (
	"strings"
	//
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/openapi/v3"
)

func (node *Node) addHandler(method string, handler *Handler) {

	switch spec := node.Config.Spec.(type) {
	case *openapiv2.Spec:

		if spec.Paths == nil {
			spec.Paths = make(map[string]openapiv2.Path)
		}

		pathMethod := &openapiv2.PathMethod{
			OperationID: "?",
			Produces: []string{
				"application/json",
			},
			Description: "Serves the OpenAPI spec JSON.",
			Responses: map[int]*openapiv2.Response{},
		}

		// http 500 status
		s500 := pathMethod.Response(500)
		s500.Description = CONST_HTTP_STATUS_MSG_500

		if handler.Method != "GET" {
			// http 400 status
			s400 := pathMethod.Response(400)
			s400.Description = CONST_HTTP_STATUS_MSG_400
		}

		// http 200 status
		s200 := pathMethod.Response(200)
		s200.Description = CONST_HTTP_STATUS_MSG_200
		s200.Schema = &openapiv2.StatusSchema{
			Type: "object",
		}

		path := handler.Path(spec.BasePath)
		if spec.Paths[path] == nil {
			spec.Paths[path] = openapiv2.Path{}
		}
		spec.Paths[path][strings.ToLower(method)] = pathMethod

	case *openapiv3.Spec:

		if spec.Paths == nil {
			spec.Paths = make(map[string]openapiv3.Path)
		}

		pathMethod := &openapiv3.PathMethod{
			Produces: []string{
				"application/json",
			},
			Description: "Serves the OpenAPI spec JSON",
			Responses: openapiv3.Responses{
				Code200: &openapiv3.StatusCode{
					Description: "Done OK",
					Schema: openapiv3.StatusSchema{
						Type: "object",
					},
				},
			},
		}

		path := handler.Path(spec.BasePath)
		if spec.Paths[path] == nil {
			spec.Paths[path] = openapiv3.Path{}
		}
		spec.Paths[path][strings.ToLower(method)] = pathMethod

		default:
			panic("INVALID TYPE FOR HTTP METHOD SWITCH")

	}

	// update this handler's spec
	handler.updateSpecHeaders()
	handler.updateParameters()

	node.Lock()
	defer node.Unlock()
	node.Methods[method] = handler
}
