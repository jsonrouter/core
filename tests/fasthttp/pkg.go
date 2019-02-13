package fasthttptest

import (
	"time"
	"testing"
	//
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/fasthttp"
	//
	"github.com/jsonrouter/core/tests/common"
)

func TestServer(t *testing.T, node *tree.Node) *common.TestHTTPStruct {

	s := openapiv2.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	log := logs.NewClient().NewLogger()

	root, service := jsonrouter.New(log, s)

	self := &common.TestHTTPStruct{
		T: t,
		Met: &root.Config.Metrics,
		Port: common.CONST_PORT_FASTHTTP,
	}

	// make the supplied routing work on this root node
	if node != nil {
		root.Use(node)
	}

	go func() {
		panic(
			service.Serve(common.CONST_PORT_FASTHTTP),
		)
	}()

	// wait for router to be serving
	time.Sleep(time.Second / 10)

	return self
}
