package fasthttptest

import (
	"fmt"
	"time"
	"testing"
	ht "net/http"
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/standard"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/tests/common"
)

func TestServer(t *testing.T, node *tree.Node) *common.TestHTTPStruct {

	s := openapiv2.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	log := logs.NewClient().NewLogger()

	service, err := jsonrouter.New(log, s)
	if err != nil {
		t.Error(err)
		t.Fail()
		return nil
	}

	self := &common.TestHTTPStruct{
		T: t,
		Met: &service.Node.Config.Metrics,
		Port: common.CONST_PORT_STANDARD,
	}

	// make the supplied routing work on this root node
	if node != nil {
		service.Node.Use(node)
	}
	
	go func() {
		panic(
			ht.ListenAndServe(
				fmt.Sprintf(
					":%d",
					common.CONST_PORT_STANDARD,
				),
				service,
			),
		)
	}()

	// wait for router to be serving
	time.Sleep(time.Second / 10)

	return self
}
