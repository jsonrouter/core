package tree

import (
	"reflect"
	"strings"
	"strconv"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/openapi/v3"
)

// opCounter makes sure there are no duplicate OperationIDs
var opCounter int

// addHandler literally adds a handler to the node object under a http method
func (node *Node) addHandler(method string, handler *Handler) {

	switch spec := node.Config.Spec.(type) {

	case nil:

		panic("THE API SPEC OBJECT IS NIL!")

	case *openapiv2.Spec:

		if spec.Paths == nil {
			spec.Paths = make(map[string]openapiv2.Path)
		}

		pathMethod := &openapiv2.PathMethod{
			OperationID: "?"+strconv.Itoa(opCounter),
			Produces: []string{
				"application/json",
			},
			Description: "Serves the OpenAPI spec JSON.",
			Responses: map[int]*openapiv2.Response{},
		}

		opCounter++

		// http 500 status
		s500 := pathMethod.Response(500)
		s500.Description = http.CONST_HTTP_STATUS_MSG_500

		if handler.Method != "GET" {
			// http 400 status
			s400 := pathMethod.Response(400)
			s400.Description = http.CONST_HTTP_STATUS_MSG_400
		}

		// http 200 status
		s200 := pathMethod.Response(200)
		s200.Description = http.CONST_HTTP_STATUS_MSG_200
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

		operation := &openapiv3.Operation{
			Summary: "",
			Description: "Serves the OpenAPI spec JSON",
			Parameters: []*openapiv3.Parameter{},
			Responses:  map[int]*openapiv3.Response{
				200: &openapiv3.Response{
					Description: "OK!",
				},
				500: &openapiv3.Response{
					Description: "Unknown Server Error",
				},
				400: &openapiv3.Response{
 					Description: "Bad Request",
				},
			},
		}



		path := handler.Path("")
		if spec.Paths[path] == nil {
			spec.Paths[path] = openapiv3.Path{}
		}
		spec.Paths[path][strings.ToLower(method)] = operation

	default:
		panic("INVALID TYPE FOR THE PROVIDED SPEC: "+reflect.TypeOf(node.Config.Spec).String())

	}

	// update this handler's spec
	handler.updateSpecHeaders()
	handler.updateParameters()

	node.Lock()
	defer node.Unlock()
	node.Methods[method] = handler
}
