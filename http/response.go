package http

import (
	"errors"
	"strconv"
	"reflect"
	"encoding/json"
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

func (status *Status) Respond(req Request) error {

	// return with no action if handler returns nil
	if status == nil { return nil }

	switch v := status.Value.(type) {

		case nil:

		case string:

			req.Write([]byte(v))

		case []byte:

			req.Write(v)

		case [][]byte:

			for _, b := range v {
				req.Write(b)
			}

		default:

			req.SetHeader("Content-Type", "application/json")
			b, err := json.Marshal(status.Value)
			if err != nil {
				return err
			}
			req.Write(b)

	}

	if status.Code >= 200 && status.Code < 300 {
		return nil
	}

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.MessageString()

	req.HttpError(statusMessage, status.Code)

	return errors.New(statusMessage)
}
