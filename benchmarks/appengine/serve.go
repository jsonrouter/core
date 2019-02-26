package server
import (
	"fmt"
	ht "net/http"
	//
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tests/common"
	"github.com/jsonrouter/validation"
	"github.com/jsonrouter/platforms/appengine"
)

type testStruct struct {
	name string
	data int
}

func returnStruct() interface{} {
	return &testStruct {
		name : "steve",
		data : 10,
	}
}

type App struct {
	*common.TestHTTPStruct
}

func (app *App) ApiGET(req http.Request) *http.Status {
	return req.Respond(
		req.Param("x").(int),
	)
}

func (app *App) ApiPOST(req http.Request) *http.Status {
	return req.Respond(
		req.Param("x").(int),
	)
}

func Start() {
	app := &App{}

	s := openapiv2.New(common.CONST_SPEC_HOST, common.CONST_SPEC_TITLE)
	s.BasePath = common.CONST_SPEC_BASEPATH
	s.Info.Contact.URL = common.CONST_SPEC_URL
	s.Info.Contact.Email = common.CONST_SPEC_EMAIL
	s.Info.License.URL = common.CONST_SPEC_URL

	service, err := jsonrouter.New(s)
	if err != nil {
		panic(err)
	}

	endpoint := service.Node.Add("/endpoint").Param(validation.Int(), "x")
	endpoint.GET(app.ApiGET)
	endpoint.POST(app.ApiPOST).Required(
		validation.Payload{
			"hello": validation.String(1, 100),
		},
	)

	fmt.Println("Serving:", common.CONST_PORT_APPENGINE)

	panic(
			ht.ListenAndServe(
				fmt.Sprintf(
					":%d",
					common.CONST_PORT_APPENGINE,
				),
				service,
			),
	)
}
