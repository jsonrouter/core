package benchmarks

import (
	"fmt"
	"testing"
	//
	"github.com/jsonrouter/core/tests/common"
	"github.com/go-resty/resty"
	//
	"github.com/jsonrouter/core/benchmarks/fasthttp"

)

func init() {
	go server.Start()
}

func BenchmarkFasthttpGET(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_FASTHTTP)

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().Get(url)
		if err != nil || resp.StatusCode() != 200 {
			b.Error(err)
			return
		}
	}
}

func BenchmarkFasthttpPOST(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_FASTHTTP)

	payload := map[string]interface{}{
		"hello": "world",
	}

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().SetBody(payload).Post(url)
		if err != nil || resp.StatusCode() != 200 {
			b.Error(resp.String())
			return
		}
	}
}
