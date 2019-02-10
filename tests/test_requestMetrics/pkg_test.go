package main

import (
	"fmt"
	"time"
	"testing"
	//
	"github.com/go-resty/resty"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/http"
	//
	fasthttptest "github.com/jsonrouter/core/tests/fasthttp"
	standardtest "github.com/jsonrouter/core/tests/standard"
	appenginetest "github.com/jsonrouter/core/tests/appengine"
	"github.com/jsonrouter/core/tests/common"
)

type App struct {
	*common.TestHTTPStruct
}

func (app *App) ApiGET(req http.Request) *http.Status {

	req.Log().Debug("GET")

	//defer func() {
		x := req.Param("x").(int)
		val := app.Met.Counters["requestCount"].GetValue()
		if int(val) != (x + 1) {
			req.Log().Debugf("GET: CORRECT VALUE IS %v NOT %v", x, int(val))
			app.T.Fail()
		}
	//}()

	return nil
}

func (app *App) ApiPOST(req http.Request) *http.Status {

	req.Log().Debug("POST")

	x := req.Param("x").(int)
	val := app.Met.Counters["requestCount"].GetValue()
	if int(val) != (x + 2) {
		req.Log().Debugf("POST: CORRECT VALUE IS %v", x)
		return req.Fail()
	}

	return nil
}

func TestFastHttpMetrics(t *testing.T) {

	app := &App{}

	// create routing structure
	config := &tree.Config{
		Spec: openapiv2.New("localhost", "test"),
	}
	root := tree.NewNode(config)
	endpoint := root.Add("/endpoint").Param(validation.Int(), "x")
	endpoint.GET(app.ApiGET)
	endpoint.POST(app.ApiPOST)

	a := map[string]func(t *testing.T, node *tree.Node) *common.TestHTTPStruct{
		"fasthttp": fasthttptest.TestServer,
		"appengine": appenginetest.TestServer,
		"standard": standardtest.TestServer,
	}

	for name, fnc := range a {

		t.Run(
			"RUNNING TEST FOR PLATFORM - " + name,
			func (t *testing.T) {

				app.TestHTTPStruct = fnc(t, root)

				url := fmt.Sprintf("http://localhost:%d/openapi.spec.json", common.CONST_PORT_FASTHTTP)

				resp, err := resty.R().Get(url)
				if err != nil || resp.StatusCode() == 500 {
					t.Error(resp.String())
					t.Fail()
					return
				}

				for x := 0; x < 100; x+=2 {

					url = fmt.Sprintf("http://localhost:%d/endpoint/%d", common.CONST_PORT_FASTHTTP, x)

					resp, err := resty.R().Get(url)
					if err != nil || resp.StatusCode() == 500 {
						t.Error(resp.String())
						t.Fail()
						return
					}

					resp, err = resty.R().Post(url)
					if err != nil || resp.StatusCode() == 500 {
						t.Error(resp.String())
						t.Fail()
						return
					}

				}

				time.Sleep(3 * time.Second)

			},
		)

	}

}
