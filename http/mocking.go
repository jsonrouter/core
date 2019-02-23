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

// NewMockRequest creates an implementation of the Request interface for testing or other.
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

// UID returns a UUID that has been generated randoomly for this request.
func (ti *MockRequest) UID() (string, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

// Testing
func (req *MockRequest) Testing() bool {
	return true
}

// FullPath
func (ti *MockRequest) FullPath() string {
	return ti.path
}

// IsTls
func (ti *MockRequest) IsTLS() bool {
	return false
}

// Method
func (ti *MockRequest) Method() string {
	return ti.method
}

// Device
func (ti *MockRequest) Device() string {
	return ti.method
}

// Body
func (ti *MockRequest) Body(s string) interface{} {
	return 0
}

// Param
func (ti *MockRequest) Param(k string) interface{} {
	return ti.params[k]
}

// Params
func (ti *MockRequest) Params() map[string]interface{} {
	return ti.params
}

// SetParam
func (ti *MockRequest) SetParam(k string, i interface{}) {
	ti.params[k] = i
}

// SetParams
func (ti *MockRequest) SetParams(m map[string]interface{}) {
	ti.params = m
}

// BodyParam
func (ti *MockRequest) BodyParam(k string) interface{} {
	return ti.bodyParams[k]
}

// BodyParams
func (ti *MockRequest) BodyParams() map[string]interface{} {
	return ti.bodyParams
}


// SetBodyParam
func (ti *MockRequest) SetBodyParam(k string, i interface{}) {
	ti.bodyParams[k] = i
}

// SetBodyParams
func (ti *MockRequest) SetBodyParams(m map[string]interface{}) {
	ti.bodyParams = m
}

// SetRequestHeader
func (ti *MockRequest) SetRequestHeader(k, v string) {
	ti.requestHeaders[k] = v
}

// GetRequestHeader
func (ti *MockRequest) GetRequestHeader(k string) string {
	return ti.requestHeaders[k]
}

// SetResponseHeader
func (ti *MockRequest) SetResponseHeader(k, v string) {
	ti.responseHeaders[k] = v
}

// GetResponseHeader
func (ti *MockRequest) GetResponseHeader(k string) string {
	return ti.responseHeaders[k]
}

// RawBody
func (ti *MockRequest) RawBody() (*Status, []byte) {
	return nil, []byte{}

}

// ReadBodyObject
func (ti *MockRequest) ReadBodyObject() *Status {
	return nil
}

// ReadBodyArray
func (ti *MockRequest) ReadBodyArray() *Status {
	return nil
}

// BodyObject
func (ti *MockRequest) BodyObject() map[string]interface{} {
	return map[string]interface{}{}
}

// BodyArray
func (ti *MockRequest) BodyArray() []interface{} {
	return []interface{}{}
}

// Redirect
func (ti *MockRequest) Redirect(s string, x int) *Status {
	return nil
}

// ServeFile
func (ti *MockRequest) ServeFile(s string) {

}

// HttrError
func (ti *MockRequest) HttpError(s string, x int) {

}

// Writer
func (ti *MockRequest) Writer() io.Writer {
	return &rW{}
}

// WriteString
func (ti *MockRequest) WriteString(s string) {

}

// Write
func (ti *MockRequest) Write(b []byte) (int, error) {
	return 0, nil
}

// Fail
func (ti *MockRequest) Fail() *Status {
	return Fail()
}

// Respond
func (ti *MockRequest) Respond(args ...interface{}) *Status {
	return Respond(args...)
}

// Log
func (ti *MockRequest) Log() logging.Logger {
	return ti.log
}

// Res
func (ti *MockRequest) Res() http.ResponseWriter {
	return &rW{}
}

// R
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

// Status
func (w *rW) Status() int {
	return w.status
}

// Size
func (w *rW) Size() int {
	return w.size
}

// Header
func (w *rW) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write
func (w *rW) Write(data []byte) (int, error) {

	written, err := os.Stdin.Write(data)
	w.size += written

	return written, err
}

// WriteHeader
func (w *rW) WriteHeader(statusCode int) {

	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
