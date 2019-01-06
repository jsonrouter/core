package main

import (
	"fmt"
	"errors"
	"testing"
	"encoding/json"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/standard"
)

const (
	CONST_SPEC_HOST = "dummyhost"
	CONST_SPEC_TITLE = "my title"
	CONST_SPEC_URL = "https://example.com"
	CONST_SPEC_EMAIL = "address@example.com"
	CONST_SPEC_LICENSE = "MIT"
	CONST_SPEC_BASEPATH = "/api"
)

type TestObject struct {
	Hello string
	World int
}

func TestMain(t *testing.T) {

	s := openapiv2.NewV2(CONST_SPEC_HOST, CONST_SPEC_TITLE)
	s.BasePath = CONST_SPEC_BASEPATH
	s.Info.Contact.URL = CONST_SPEC_URL
	s.Info.Contact.Email = CONST_SPEC_EMAIL
	s.Info.License.URL = CONST_SPEC_URL

	root, err := jsonrouter.New(
		logs.NewClient().NewLogger(),
		s,
	)
	if err != nil {
		panic(err)
	}

	api := root.Add(CONST_SPEC_BASEPATH).Add("test")

	api.GET(
		dummyHandler,
	).Description(
		"This is a GET endpoint!",
	).Response(
		TestObject{},
	)

	api.POST(
		dummyHandler,
	).Description(
		"This is a POST endpoint!",
	).Required(
		validation.Payload{
			"hello": validation.String(10, 20).Description("The hellos!").Default("Helloy"),
		},
	).Optional(
		validation.Payload{
			"world": validation.Int().Description("The worlds!").Default(2),
		},
	).Response(
		TestObject{},
	)

	apiResource := api.Add("/resource").Param(
		validation.String(1, 64).Description("The id of the user."),
		"id",
	)

	root.Config.Log.DebugJSON(apiResource.Validations)

		apiResource.GET(
			dummyHandler,
		).Description(
			"Handles access to the resource",
		).Response(
			TestObject{},
		)

	req := http.NewMockRequest("", "")
	spec := root.Config.Spec.(*openapiv2.Spec)
//	openapi.BuildV2(spec, root.Config.Handlers)

	req.Log().DebugJSON(spec.Paths)

	t.Run(
		"Test the spec",
		func (t *testing.T) {

			if spec.Paths["/openapi.spec.v2.json"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			if spec.Paths["/openapi.spec.v3.json"] == nil {
				t.Error(fmt.Errorf("SPEC HAS INVALID PATHS! %v", len(spec.Paths)))
			}

			if spec.Host != CONST_SPEC_HOST {
				t.Error(errors.New("SPEC HAS INVALID HOST!"))
			}

			if spec.Info.Title != CONST_SPEC_TITLE {
				t.Error(errors.New("SPEC HAS INVALID TITLE!"))
			}

			if spec.Info.Contact.URL != CONST_SPEC_URL {
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

	b, _ := json.Marshal(spec)
	fmt.Println(string(b))
}

func dummyHandler(req http.Request) *http.Status {

	return nil
}
