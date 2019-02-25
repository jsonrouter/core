package benchmarks

import (
	"fmt"
	"testing"	
	"github.com/jsonrouter/core/tests/common"
	"github.com/go-resty/resty"

)

func BenchmarkFasthttp(b *testing.B) {

	url := fmt.Sprintf("http://localhost:%d/endpoint/0", common.CONST_PORT_FASTHTTP)

	// Actual benchmark starts here
	for n := 0; n < b.N; n++ {
		for x := 0; x < 1000; x+=1 {
			resp, err := resty.R().Get(url)
			if err != nil || resp.StatusCode() == 500 {
				b.Error(resp.String())
				return
			}
		}
	}
}