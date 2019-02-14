package main

import (
	"fmt"
	"encoding/json"
	"time"
	"testing"
	//
	"github.com/go-resty/resty"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/metrics"
	"github.com/jsonrouter/platforms/fasthttp"
	"github.com/jsonrouter/logging/testing"
	
	fasthttptest "github.com/jsonrouter/core/tests/fasthttp"
	standardtest "github.com/jsonrouter/core/tests/standard"
	//appenginetest "github.com/jsonrouter/core/tests/appengine"
	"github.com/jsonrouter/core/tests/common"
)

type App struct {
	*common.TestHTTPStruct
	metrics *metrics.Metrics
}

func (app *App) ApiGET(req http.Request) *http.Status {
	x := req.Param("x").(int)
	req.SetParam("y", x)
	return nil
}

func (app *App) ApiPOST(req http.Request) *http.Status {
	x := req.Param("x").(int)
	req.SetParam("y", x)

	return nil
}

func TestFastHttpMetrics(t *testing.T) {

	app := &App {
		metrics: &metrics.Metrics{
			BenchMarks: map[string]*metrics.BenchMark{
				"RequestMethods": &metrics.BenchMark {
					Name: "RequestMethods",
					Timers: map[string]*metrics.Timer{
						"GET": &metrics.Timer{
							Name : "GET",
							BufferSize : 1000,
						},
						"POST": &metrics.Timer{
							Name : "POST",
							BufferSize : 1000,
						},
					},
				},
			},
		Results: map[string]interface{}{},
		},
	}


	// create routing structure
	root, _ := jsonrouter.New(logs.NewClient().NewLogger("Server"), openapiv2.New("localhost", "test"))
	
	endpoint := root.Add("/endpoint").Param(validation.Int(), "x")
	endpoint.GET(app.ApiGET)
	endpoint.POST(app.ApiPOST)

	a := map[string]func(t *testing.T, node *tree.Node) *common.TestHTTPStruct{
		"fasthttp": fasthttptest.TestServer,
		//"appengine": appenginetest.TestServer,
		"standard": standardtest.TestServer,
	}

	for name, fnc := range a {

		t.Run(
			"RUNNING TEST FOR PLATFORM - " + name,
			func (t *testing.T) {

				app.TestHTTPStruct = fnc(t, root)
				app.Met = app.TestHTTPStruct.Met

				url := fmt.Sprintf("http://localhost:%d/endpoint/%d", app.TestHTTPStruct.Port, x)

				for x := 0; x < 1000; x+=1 {

					app.metrics.BenchMarks["RequestMethods"].StartTimer("GET")
					resp, err := resty.R().Get(url)
					app.metrics.BenchMarks["RequestMethods"].StopTimer("GET")
					
					if err != nil || resp.StatusCode() == 500 {
						t.Error(resp.String())
						t.Fail()
						return
					}
				}

				for x = 0; x < 1000; x+=1 {
					app.metrics.BenchMarks["RequestMethods"].StartTimer("POST")
					resp, err = resty.R().Post(url)
					app.metrics.BenchMarks["RequestMethods"].StopTimer("POST")
					
					if err != nil || resp.StatusCode() == 500 {
						t.Error(resp.String())
						t.Fail()
						return
					}
				}



				time.Sleep(3 * time.Second)
				app.metrics.BenchMarks["RequestMethods"].Update(&app.metrics.Results)

				res, err := json.Marshal(app.metrics.Results)
				if err != nil {
					t.Log(err)
					t.Fail()
				} 
				fmt.Println(string(res))
			},
		)

	}

}
