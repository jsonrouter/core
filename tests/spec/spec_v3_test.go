package tests

import (
	"fmt"
	//"errors"
	"testing"
	"encoding/json"
	//
//	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v3"
	"github.com/jsonrouter/platforms/standard"
	"github.com/jsonrouter/core/tests/common"
)

func TestSpecV3(t *testing.T) {

	s := openapiv3.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	//s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	root, err := jsonrouter.New(
		logs.NewClient().NewLogger(),
		s,
	)
	if err != nil {
		panic(err)
	}

	// load the demo routing info
	common.TestTree(root.Node)

	//req := http.NewMockRequest("", "")
	spec := root.Config.Spec.(*openapiv3.Spec)
/*
	t.Run(
		"Test the spec",
		func (t *testing.T) {

			if spec.Paths["/openapi.spec.v2.json"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			if spec.Paths["/openapi.spec.v3.json"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			if spec.Host != common.CONST_SPEC_HOST {
				t.Error(errors.New("SPEC HAS INVALID HOST!"))
			}

			if spec.Info.Title != common.CONST_SPEC_TITLE {
				t.Error(errors.New("SPEC HAS INVALID TITLE!"))
			}

			if spec.Info.Contact.URL != common.CONST_SPEC_URL {
				t.Error(errors.New("SPEC HAS INVALID CONTACT URL!"))
			}

			if len(spec.Paths) != 4 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUM OF PATHS! %v", len(spec.Paths)))
			}

			if spec.Paths["/test/resource/{id}"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			pl := len(spec.Paths["/test/resource/{id}"]["get"].Parameters)
			if pl != 1 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUMBER OF PARAMETERS! %v", pl))
			}

			if len(spec.Definitions) != 1 {
				t.Error(fmt.Errorf("SPEC HAS INVALID NUMBER OF DEFINITIONS! %v", len(spec.Definitions)))
			}

		},
	)
*/
	b, _ := json.Marshal(spec)
	fmt.Println(string(b))
}
