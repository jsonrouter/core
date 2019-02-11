package common

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/metrics"
)

type TestHTTPStruct struct {
	T *testing.T
	Met metrics.Metrics
}

func dummyHandler(req http.Request) *http.Status {

	return nil
}
