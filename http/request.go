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
	Write([]byte)
	WriteString(string)
	Fail() *Status
	Respond(args ...interface{}) *Status
	// logging
	Log() logging.Logger
	//
	Res() http.ResponseWriter
	R() interface{}
}

// returns a standard 500 http error status
func Fail() *Status {
	return Respond(500, "UNEXPECTED APPLICATION ERROR")
}

func Respond(args ...interface{}) *Status {

	var ok bool
	s := &Status{
		Code: 200,
	}

	l := len(args)

	switch l {

		case 1:

			switch args[0].(type) {

			case nil:
				s.Value = args[0]
			case string:
				s.Value = args[0]
			case []byte:
				s.Value = args[0]
			case [][]byte:
				s.Value = args[0]

			default:

				b, err := json.Marshal(args[0])
				if err != nil {
					return Respond(500, CONST_HTTP_STATUS_MSG_500)
				}

				s.Value = b

			}

			return s

		case 2, 3:

			s.Code, ok = args[0].(int); if !ok {
				return &Status{nil, 501, "Respond(...) METHOD HAS 2 ARGS; UNEXPECTED ARG 0 TYPE: " + reflect.TypeOf(args[0]).String()}
			}

			// argument 1 is now an interface, so we can handle errors

			if args[1] == nil {
				panic("2nd ARGUEMENT TO RESPOND IS NIL")
			}

			s.Message = args[1]

			if l == 3 {
				s.Value = args[2]
			}

			return s

		default:

			return &Status{nil, 400, "INVALID STATUS ARGS LENGTH: "+strconv.Itoa(len(args))}

	}

	return nil // Unreachable code warning
}
