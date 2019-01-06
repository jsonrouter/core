package tree

import	(
	"fmt"
	"sync"
	"path"
	"mime"
	"reflect"
	"strings"
	www "net/http"
	"io/ioutil"
	"encoding/json"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/validation"
)

type HandlerFunction func (req http.Request) *http.Status

type Handler struct {
	Config *Config
	Node *Node
	Method string
	Descr string
	Headers map[string]string
	Function func (req http.Request) *http.Status
	File *File
	responseSchema interface{}
	payloadSchema []interface{}
	patchSchema []interface{}
	spec interface{}
	sync.RWMutex
}

func (handler *Handler) Path(removePrefix ...string) string {
	return strings.Replace(handler.Node.FullPath(), removePrefix[0], "", 1)
}

func (handler *Handler) DetectContentType(req http.Request, filePath string) *http.Status {

	if handler.File.Cache == nil || !handler.Node.Config.CacheFiles {

		// handle potential trailing slash on folder path declaration
		filePath := strings.Replace(filePath, "//", "/", -1)

		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return req.Respond(404, err.Error())
		}

		handler.File.Cache = b
		handler.File.MimeType = mime.TypeByExtension(path.Ext(filePath))
		if handler.File.MimeType == "" {
			handler.File.MimeType = www.DetectContentType(b)
		}
	}

	return nil
}

func (handler *Handler) ApiUrl() string {

	var name string

	parts := strings.Split(handler.Node.FullPath(), "/")

	for _, part := range parts {

		if len(part) == 0 { continue }

		if string(part[0]) == ":" {

			part = "'+" + part[1:] + "+'"

		}

		name += "/" + part

	}

	return "'" + name + "'"
}

// Describes the function via the spec JSON
func (handler *Handler) Description(descr string) *Handler {

	handler.Descr = descr

	return handler
}

// Applies model which describes request payload
func (handler *Handler) Patch(objects ...*validation.Patch) *Handler {

	for _, object := range objects {
		handler.patchSchema = append(
			handler.payloadSchema,
			object,
		)
	}

	return handler
}

// Applies model which describes request payload
func (handler *Handler) BodyIsObject() *Handler {

	handler.payloadSchema = append(
		handler.payloadSchema,
		&validation.Object{},
	)

	return handler
}

// Applies model which describes request payload
func (handler *Handler) BodyIsArray() *Handler {

	handler.payloadSchema = append(
		handler.payloadSchema,
		&validation.Array{},
	)

	return handler
}

// Applys model which describes response schema
func (handler *Handler) Response(schema ...interface{}) *Handler {

	handler.responseSchema = schema[0]

	return handler
}

// Validates any payload present in the request body, according to the payloadSchema
func (handler *Handler) ReadPayload(req http.Request) *http.Status {

	// handle payload

	var paramCount int
	var optionalCount int
	var readBodyObject bool
	bodyParams := map[string]interface{}{}
	statusMessages := map[string]*http.Status{}

	for _, schema := range handler.payloadSchema {

		switch params := schema.(type) {

			case nil:

				// do nothing

			case []byte:

			// do nothing

			case map[string]interface{}:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

			case []interface{}:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

			case *validation.Array:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				array := *params

				switch len(array) {

					case 1:

						return req.Respond(400, "INVALID TYPE FOR ARRAY PAYLOAD SCHEMA, EXPECTS 0 OR 2 ARGS (*ValidationConfig, paramKey)")

					case 2:

						vc, ok := array[0].(*validation.Config); if !ok { return req.Respond(500, "INVALID ARRAY PAYLOAD SCHEMA VALIDATION CONFIG") }

						paramKey, ok := array[1].(string); if !ok { return req.Respond(500, "INVALID ARRAY PAYLOAD SCHEMA PARAM KEY") }

						status, array := vc.BodyFunction(req, req.BodyArray()); if status != nil {

							req.Log().DebugJSON(req.BodyArray())
							//return req.Respond(400, "INVALID TYPE FOR ARRAY PAYLOAD ITEM, EXPECTED: "+vc.Type())

							return status
						}

						req.SetParam(paramKey, array)

				}

			case *validation.Object:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

			case *validation.Payload:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				object := *params

				for key, vc := range object {
					paramCount++
					status, x := vc.BodyFunction(
						req,
						req.Body(key),
					)
					if status != nil {
						// dont leak data to logs
						//status.Value = req.Body(key)
						status.Message = fmt.Sprintf("%s KEY '%s'", status.MessageString(), key)
						statusMessages[key] = status
					} else {
						bodyParams[key] = x
					}
				}

			case *validation.Optional:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				object := *params

				for key, vc := range object {
					optionalCount++
					status, x := vc.BodyFunction(
						req,
						req.Body(key),
					)
					if status == nil {
						bodyParams[key] = x
					}
				}

			default:

				return req.Respond(500, "INVALID OPTIONAL PAYLOAD SCHEMA CONFIG TYPE: "+reflect.TypeOf(params).String())

		}

	}

	if len(statusMessages) > 0 {
		b, _ := json.Marshal(statusMessages)
		return req.Respond(500, string(b))
	}

	lp := len(bodyParams)
	if len(bodyParams) < paramCount {
		return req.Respond(
			400,
			fmt.Sprintf(
				"INVALID PAYLOAD FIELD COUNT %v EXPECTED %v/%v",
				lp,
				paramCount,
				paramCount+optionalCount,
			),
		)
	}

	req.SetBodyParams(bodyParams)

	return nil
}

func (handler *Handler) UseFunction(f interface{}) {

	switch v := f.(type) {

		case func(http.Request) *http.Status:

		  handler.Function = v

		default:

		  panic("INVALID ARGUMENT TYPE FOR HANDLER METHOD FUNCTION DECLARATION")

	}
}
