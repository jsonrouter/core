package http

import (
	"strconv"
	"reflect"
)

type Status struct {
	Value interface{} `json:"value,omitempty"`
	Code int `json:"code"`
	Message interface{} `json:"message"`
}

// MessageString serialises the status to a string.
func (status *Status) MessageString() string {
	switch v := status.Message.(type) {
		case nil:
			return "null"
		case error:
			return v.Error()
		case string:
			return v
	}
	return "INVALID STATUS MESSAGE TYPE: "+reflect.TypeOf(status.Message).String()
}

// Respond executes the status on a HTTP level.
func (status *Status) Respond(req Request) {

	// return with no action if handler returns nil
	if status == nil { return }

	switch v := status.Value.(type) {

		case nil:

		case string:

			req.WriteString(v)

		case []byte:

			req.Write(v)

		case [][]byte:

			for _, b := range v {
				// if writing part of the response fails, then return a HTTP 500 status.
				_, err := req.Write(b)
				if req.Log().Error(err) {
					status = req.Fail()
				}
			}

		default:

			panic("THIS CODE SHOULD BE UNREACHABLE")

	}

	if status.Code >= 200 && status.Code < 300 {
		return
	}

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.MessageString()

	req.HttpError(statusMessage, status.Code)
}
