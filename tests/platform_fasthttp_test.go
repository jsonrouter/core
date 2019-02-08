package tests

import (
	"fmt"
	"time"
	"testing"
	//
	"github.com/go-resty/resty"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/fasthttp"
)

const (
	CONST_TEST_PORT = 999
)

func TestFastHttp(t *testing.T) {

	s := openapiv2.New(CONST_SPEC_HOST, CONST_SPEC_TITLE)
	s.BasePath = CONST_SPEC_BASEPATH
	s.Info.Contact.URL = CONST_SPEC_URL
	s.Info.Contact.Email = CONST_SPEC_EMAIL
	s.Info.License.URL = CONST_SPEC_URL

	log := logs.NewClient().NewLogger()

	root, service := jsonrouter.New(log, s)

	met := root.Config.Metrics

	endpoint := root.Add("/endpoint")

	endpoint.GET(
		func (req http.Request) *http.Status {

			defer func() {

				val := met.Counters["requestCount"].GetValue()
				if val != 1 {
					req.Log().Debugf("CORRECT VALUE IS %v", val)
					t.Fail()
				}

			}()

			return nil
		},
	)

	endpoint.POST(
		func (req http.Request) *http.Status {

			defer func() {

				val := met.Counters["requestCount"].GetValue()
				if val != 1 {
					req.Log().Debugf("CORRECT VALUE IS %v", val)
					t.Fail()
				}

			}()

			return nil
		},
	)

	go func() {
		if err := service.Serve(CONST_TEST_PORT); err != nil {
			t.Error(err)
			t.Fail()
		}
	}()

	time.Sleep(time.Second)

	url := fmt.Sprintf("http://localhost:%d/endpoint", CONST_PORT)

	resp, err := resty.R().Get(url)
	if log.Error(err) || resp.StatusCode() == 500 {
		log.NewError(resp.String())
		t.Fail()
		return
	}

	resp, err = resty.R().Post(url)
	if log.Error(err) || resp.StatusCode() == 500 {
		log.NewError(resp.String())
		t.Fail()
		return
	}

	time.Sleep(3 * time.Second)

}
