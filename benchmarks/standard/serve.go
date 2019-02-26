package server
import (
	"fmt"
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/platforms/standard"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tests/common"
	"github.com/jsonrouter/validation"
	ht "net/http"
	//"expvar"
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

	log := logs.NewClient().NewLogger()

	service, err := jsonrouter.New(log, s)
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

	//expvar.Publish("METRIC", expvar.Func(returnStruct))
	fmt.Println("Serving:", common.CONST_PORT_STANDARD)

	panic(
			ht.ListenAndServe(
				fmt.Sprintf(
					":%d",
					common.CONST_PORT_STANDARD,
				),
				service,
			),
	)
}
