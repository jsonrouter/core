package fasthttptest

import (
	"fmt"
	"time"
	"testing"
	ht "net/http"
	//
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/appengine"
	//
	"github.com/jsonrouter/core/tests/common"
)

func TestServer(t *testing.T, node *tree.Node) *common.TestHTTPStruct {

	s := openapiv2.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	service, err := jsonrouter.New(s)
	if err != nil {
		t.Error(err)
		t.Fail()
		return nil
	}

	self := &common.TestHTTPStruct{
		T: t,
		Met: service.Node.Config.Metrics,
	}

	// make the supplied routing work on this root node
	service.Node.Use(node)

	go func() {
		panic(
			ht.ListenAndServe(
				fmt.Sprintf(
					":%s",
					common.CONST_PORT_APPENGINE,
				),
				service,
			),
		)
	}()

	// wait for router to be serving
	time.Sleep(time.Second / 10)

	return self
}
