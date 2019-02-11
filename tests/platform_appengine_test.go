package tests

import (
	"testing"
	//
//	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/appengine"
)

func TestAppEngine(t *testing.T) {

	s := openapiv2.New(CONST_SPEC_HOST, CONST_SPEC_TITLE)
	s.BasePath = CONST_SPEC_BASEPATH
	s.Info.Contact.URL = CONST_SPEC_URL
	s.Info.Contact.Email = CONST_SPEC_EMAIL
	s.Info.License.URL = CONST_SPEC_URL

	_, err := jsonrouter.New(
		s,
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
