package tests

import (
	"fmt"
	"time"
	"testing"
	//
	"github.com/go-resty/resty"
	"github.com/chrysmore/metrics"
	//
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/fasthttp"
)

type TestHTTPStruct struct {
	t *testing.T
	met metrics.Metrics
}

func (self *TestHTTPStruct) ApiGET(req http.Request) *http.Status {

	req.Log().Debug("GET")

	defer func() {
		x := req.Param("x").(int)
		val := self.met.Counters["requestCount"].GetValue()
		if int(val) != (x + 1) {
			req.Log().Debugf("GET: CORRECT VALUE IS %v NOT %v", x, int(val))
			self.t.Fail()
		}
	}()


	req.Log().Debug("GET2")


	return nil
}

func (self *TestHTTPStruct) ApiPOST(req http.Request) *http.Status {

	req.Log().Debug("POST")

	x := req.Param("x").(int)
	val := self.met.Counters["requestCount"].GetValue()
	if int(val) != (x + 2) {
		req.Log().Debugf("POST: CORRECT VALUE IS %v", x)
		return req.Fail()
	}

	return nil
}

func TestFastHttp(t *testing.T) {

	s := openapiv2.New(CONST_SPEC_HOST, CONST_SPEC_TITLE)
	s.BasePath = CONST_SPEC_BASEPATH
	s.Info.Contact.URL = CONST_SPEC_URL
	s.Info.Contact.Email = CONST_SPEC_EMAIL
	s.Info.License.URL = CONST_SPEC_URL

	log := logs.NewClient().NewLogger()

	root, service := jsonrouter.New(log, s)

	self := &TestHTTPStruct{
		met: root.Config.Metrics,
	}

	endpoint := root.Add("/endpoint").Param(validation.Int(), "x")

	endpoint.GET(self.ApiGET)
	endpoint.POST(self.ApiPOST)

	go func() {

		//spec := root.Config.Spec.(*openapiv2.Spec)
		//log.DebugJSON(spec)

		if err := service.Serve(CONST_PORT); err != nil {
			t.Error(err)
			t.Fail()
		}
	}()

	time.Sleep(time.Second)

	url := fmt.Sprintf("http://localhost:%d/openapi.spec.json", CONST_PORT)

	resp, err := resty.R().Get(url)
	if log.Error(err) || resp.StatusCode() == 500 {
		log.NewError(resp.String())
		t.Fail()
		return
	}

	log.Debug(resp.String())

	for x := 0; x < 10; x+=2 {

		url = fmt.Sprintf("http://localhost:%d/endpoint/%d", CONST_PORT, x)

		log.Debug(url)

		resp, err := resty.R().Get(url)
		if log.Error(err) || resp.StatusCode() == 500 {
			log.NewError(resp.String())
			t.Fail()
			return
		}

/*
		resp, err = resty.R().Post(url)
		if log.Error(err) || resp.StatusCode() == 500 {
			log.NewError(resp.String())
			t.Fail()
			return
		}
*/
	}

	time.Sleep(3 * time.Second)

}
