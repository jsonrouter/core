package benchmarks

import (
	"fmt"
	"testing"
	//
	"github.com/jsonrouter/core/tests/common"
	"github.com/go-resty/resty"
	"github.com/jsonrouter/core/benchmarks/standard"
)

func init() {
	go server.Start()
}

func BenchmarkStandardGET(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_STANDARD)

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().Get(url)
		if err != nil || resp.StatusCode() != 200 {
			b.Error(err)
			return
		}
	}
}

func BenchmarkStandardPOST(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_STANDARD)

	payload := map[string]interface{}{
		"hello": "world",
	}

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().SetBody(payload).Post(url)
		if err != nil || resp.StatusCode() != 200 {
			b.Error(err)
			return
		}
	}
}
