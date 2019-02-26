package benchmarks

import (
	"fmt"
	"testing"
	//
	"github.com/jsonrouter/core/tests/common"
	"github.com/go-resty/resty"
	//
	"github.com/jsonrouter/core/benchmarks/appengine"
)

func init() {
	go server.Start()
}

func BenchmarkAppengineGET(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_APPENGINE)
	//fmt.Println(url)

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().Get(url)
		if err != nil {
			b.Error(err)
			return
		}
		if resp.StatusCode() != 200 {
			b.Error(resp.Status())
			return
		}
	}
}

func BenchmarkAppenginePOST(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_APPENGINE)
	//fmt.Println(url)

	payload := map[string]interface{}{
		"hello": "world",
	}

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		resp, err := resty.R().SetBody(payload).Post(url)
		if err != nil {
			b.Error(err)
			return
		}
		if resp.StatusCode() != 200 {
			b.Error(resp.Status())
			return
		}
	}
}
