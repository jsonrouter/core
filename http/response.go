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
				req.Write(b)
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
