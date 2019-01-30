package http

import (
	"io"
	"os"
	"net/url"
	"net/http"
	//
	"github.com/golangdaddy/go.uuid"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/logging/testing"
)

func NewMockRequest(method, path string) Request {

	return &MockRequest{
		method: method,
		path: path,
		device: "Mobile",
		params: map[string]interface{}{},
		bodyParams: map[string]interface{}{},
		requestHeaders: map[string]string{},
		responseHeaders: map[string]string{},
		log: logs.NewClient().NewLogger("MockRequest"),
	}
}

type MockRequest struct {
	method string
	device string
	path string
	params map[string]interface{}
	bodyParams map[string]interface{}
	requestHeaders map[string]string
	responseHeaders map[string]string
	log logging.Logger
}

func (ti *MockRequest) UID() (string, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (req *MockRequest) Testing() bool {
	return true
}

func (ti *MockRequest) FullPath() string { return ti.path }

func (ti *MockRequest) IsTLS() bool { return false }

func (ti *MockRequest) Method() string { return ti.method }

func (ti *MockRequest) Device() string { return ti.method }

func (ti *MockRequest) Body(s string) interface{} { return 0 }

func (ti *MockRequest) Param(k string) interface{} { return ti.params[k] }
func (ti *MockRequest) Params() map[string]interface{} { return ti.params }
func (ti *MockRequest) SetParam(k string, i interface{}) { ti.params[k] = i }
func (ti *MockRequest) SetParams(m map[string]interface{}) { ti.params = m }

func (ti *MockRequest) BodyParam(k string) interface{} { return ti.bodyParams[k] }
func (ti *MockRequest) BodyParams() map[string]interface{} { return ti.bodyParams }
func (ti *MockRequest) SetBodyParam(k string, i interface{}) { ti.bodyParams[k] = i }
func (ti *MockRequest) SetBodyParams(m map[string]interface{}) { ti.bodyParams = m }

func (ti *MockRequest) SetRequestHeader(k, v string) { ti.requestHeaders[k] = v }
func (ti *MockRequest) GetRequestHeader(k string) string { return ti.requestHeaders[k] }
func (ti *MockRequest) SetResponseHeader(k, v string) { ti.responseHeaders[k] = v }
func (ti *MockRequest) GetResponseHeader(k string) string { return ti.responseHeaders[k] }

func (ti *MockRequest) RawBody() (*Status, []byte) { return nil, []byte{} }

func (ti *MockRequest) ReadBodyObject() *Status { return nil }
func (ti *MockRequest) ReadBodyArray() *Status { return nil }

func (ti *MockRequest) BodyObject() map[string]interface{} { return map[string]interface{}{} }
func (ti *MockRequest) BodyArray() []interface{} { return []interface{}{} }

func (ti *MockRequest) Redirect(s string, x int) *Status { return nil }

func (ti *MockRequest) ServeFile(s string) { }

func (ti *MockRequest) HttpError(s string, x int) { }

func (ti *MockRequest) Writer() io.Writer { return &rW{} }
func (ti *MockRequest) Write(b []byte) { }

func (ti *MockRequest) Fail() *Status { return Fail() }

func (ti *MockRequest) Respond(args ...interface{}) *Status { return Respond(args...) }

func (ti *MockRequest) Log() logging.Logger { return ti.log }

func (ti *MockRequest) Res() http.ResponseWriter { return &rW{} }

func (ti *MockRequest) R() interface{} {
	return &http.Request{
		URL: &url.URL{
			Path: ti.path,
		},
	}
}

type rW struct {
	status int
	size   int
	http.ResponseWriter
}

func (w *rW) Status() int {
	return w.status
}

func (w *rW) Size() int {
	return w.size
}

func (w *rW) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *rW) Write(data []byte) (int, error) {

	written, err := os.Stdin.Write(data)
	w.size += written

	return written, err
}

func (w *rW) WriteHeader(statusCode int) {

	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
