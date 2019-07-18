package http

import 	(
	"io"
	"strconv"
	"reflect"
	"net/http"
	//
	"github.com/jsonrouter/logging"
	json "github.com/json-iterator/go"
)

type Request interface {
	Testing() bool
	UID() (string, error)
	FullPath() string
	IsTLS() bool
	Method() string
	Device() string
	Body(string) interface{}
	// accesses the request params of the payload
	Param(string) interface{}
	Params() map[string]interface{}
	SetParam(string, interface{})
	SetParams(map[string]interface{})
	BodyParam(string) interface{}
	BodyParams() map[string]interface{}
	SetBodyParam(string, interface{})
	SetBodyParams(map[string]interface{})
	SetResponseHeader(string, string)
	GetResponseHeader(string) string
	SetRequestHeader(string, string)
	GetRequestHeader(string) string
	RawBody() (*Status, []byte)
	ReadBodyObject() *Status
	ReadBodyArray() *Status
	BodyArray() []interface{}
	BodyObject() map[string]interface{}
	Redirect(string, int) *Status
	ServeFile(string)
	HttpError(string, int)
	Writer() io.Writer
	Write([]byte) (int, error)
	WriteString(string) (int, error)
	Fail() *Status
	Respond(args ...interface{}) *Status
	// logging
	Log() logging.Logger
	//
	Res() http.ResponseWriter
	R() interface{}
}

// Fail returns a standard 500 http error status
func Fail() *Status {
	status, _ := Respond(500, "UNEXPECTED APPLICATION ERROR")
	return status
}

// Respond creates a Status object that can be served to the client.
func Respond(args ...interface{}) (status *Status, contentType string) {

	contentType = "application/json"

	var ok bool
	status = &Status{
		Code: 200,
	}

	l := len(args)

	switch l {

		case 1:

			switch args[0].(type) {

			case nil:
				status.Value = args[0]
			case string:
				status.Value = args[0]
			case []byte:
				status.Value = args[0]
			case [][]byte:
				status.Value = args[0]

			default:

				b, err := json.Marshal(args[0])
				if err != nil {
					return Respond(500, CONST_HTTP_STATUS_MSG_500)
				}

				status.Value = b

			}

			return

		case 2, 3:

			status.Code, ok = args[0].(int); if !ok {
				return &Status{
					nil,
					501,
					"Respond(...) METHOD HAS 2 ARGS; UNEXPECTED ARG 0 TYPE: " + reflect.TypeOf(args[0]).String(),
				}, contentType
			}

			// argument 1 is now an interface, so we can handle errors

			if args[1] == nil {
				panic("2nd ARGUEMENT TO RESPOND IS NIL")
			}

			status.Message = args[1]

			if l == 3 {
				status.Value = args[2]
			}

			return

		default:

			return &Status{nil, 400, "INVALID STATUS ARGS LENGTH: "+strconv.Itoa(len(args))}, contentType

	}

	return // Unreachable code warning
}
