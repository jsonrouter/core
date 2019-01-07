package tree

import (
	"strings"
	//
	"github.com/jsonrouter/core/openapi/v2"
	//"github.com/jsonrouter/core/openapi/v3"
)

func (handler *Handler) SetHeaders(headers map[string]interface{}) *Handler {

	for k, v := range headers {
		handler.Headers[k] = v
	}

	handler.updateSpecHeaders()

	return handler
}

func (handler *Handler) updateSpecHeaders() {

	switch spec := handler.Node.Config.Spec.(type) {
	case *openapiv2.Spec:

		path := spec.Paths[handler.Path(spec.BasePath)]
		pathMethod := path[strings.ToLower(handler.Method)]

		x := 200
		for k, v := range handler.Headers {
			response := pathMethod.Response(x)
			response.Headers[k] = &openapiv2.Header{
				Type: openapiv2.Type(v),
				Default: v,
			}
		}

	/*case *openapiv3.Spec:

		path := spec.Paths[handler.Path(spec.BasePath)]
		pathMethod := path[strings.ToLower(handler.Method)]

		x := 200
		for k, v := range handler.Headers {
			response := pathMethod.Response(x)
			response.Headers[k] = &openapiv3.Header{
				Type: openapiv3.Type(v),
				Default: v,
			}
		}*/

	}
}
