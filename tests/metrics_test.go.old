package tests

import (
	//ht "net/http"
	"fmt"
	"time"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/platforms/fasthttp"
	"github.com/jsonrouter/validation"
	"testing"
	"github.com/go-resty/resty"
)

type App struct {
	logger logging.Logger
}

func TestMetrics(t *testing.T) {

	app := &App{
		logger: logs.NewClient().NewLogger("Server"),
	}

	router, service := jsonrouter.New(app.logger, openapiv2.New("localhost", "test"))

	api := router.Add("/api")

	api.Add("/test").GET(
		app.api_test,
	).Description(
		"test",
	).Response(
		validation.Object{},
	)

	t.Log("Serving:")

	go func() {
		if err := service.Serve(CONST_PORT); err != nil {
			t.Error(err)
			t.Fail()
		}
	}()

	time.Sleep(time.Second)

	url := fmt.Sprintf("http://127.0.0.1:%d/metrics", CONST_PORT)

	for i := 0; i < 10; i++ {
		resp, err := resty.R().Get(url)
		
		if (resp != nil) {t.Log(resp.String())}
			
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
	}

	time.Sleep(10 * time.Second)
	
}

func (app *App) api_test(req http.Request) *http.Status {
	time.Sleep(10 * time.Millisecond)
	return req.Respond(200)
}
