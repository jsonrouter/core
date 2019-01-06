package tree

import (
	"fmt"
	"strings"
	//
	"github.com/jsonrouter/core/openapi/v2"
//	"github.com/jsonrouter/core/openapi/v3"
)

func (node *Node) addHandler(method string, handler *Handler) {

	handler.Method = method
	handler.Config = node.Config
	handler.Node = node

	switch spec := node.Config.Spec.(type) {
	case *openapiv2.Spec:

		if spec.Paths == nil {
			spec.Paths = make(map[string]openapiv2.Path)
		}

		pathMethod := &openapiv2.PathMethod{
			Produces: []string{
				"application/json",
			},
			Description: "Serves the OpenAPI spec JSON",
			Responses: openapiv2.Responses{
				Code200: &openapiv2.StatusCode{
					Description: "Done OK",
					Schema: openapiv2.StatusSchema{
						Type: "object",
					},
				},
			},
		}

		path := handler.Path(spec.BasePath)
		fmt.Println(method, path)
		if spec.Paths[path] == nil {
			spec.Paths[path] = openapiv2.Path{}
		}
		spec.Paths[path][strings.ToLower(method)] = pathMethod

		default:
			panic("INVALID TYPE FOR HTTP METHOD SWITCH")

	}

	handler.updateParameters()

	node.Lock()
	defer node.Unlock()
	node.Methods[method] = handler
}
