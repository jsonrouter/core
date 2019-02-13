package tests

import (
	"testing"
	//
//	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/appengine"
	"github.com/jsonrouter/core/tests/common"
)

func TestAppEngine(t *testing.T) {

	s := openapiv2.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	_, err := jsonrouter.New(
		s,
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
