package tests

import (
	"fmt"
	"time"
	"testing"
	//
	"github.com/go-resty/resty"
	//
	//"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/logging/testing"
	//"github.com/jsonrouter/platforms/standard"
	fasthttptest "github.com/jsonrouter/core/tests/fasthttp"
	//
	"github.com/jsonrouter/core/tests/common"
)

type App struct {
	logger logging.Logger
	*common.TestHTTPStruct
}

func TestMetrics(t *testing.T) {

	app := &App{
		logger: logs.NewClient().NewLogger("Server"),
	}

	t.Log("Serving:")

	app.TestHTTPStruct = fasthttptest.TestServer(t, nil)

	time.Sleep(time.Second)

	url := fmt.Sprintf("http://localhost:%d/metrics", common.CONST_PORT_FASTHTTP)

	for i := 0; i < 10; i++ {
		resp, err := resty.R().Get(url)
		if (resp != nil) {
			t.Log(resp.String())
		}
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
	}

	time.Sleep(10 * time.Second)

}
