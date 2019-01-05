package tree

import 	(
	"github.com/jsonrouter/core/config"
	"github.com/jsonrouter/core/http"
)

type ModuleFunction func (req http.Request, arg interface{}) *http.Status

type Module struct {
	config *config.Config
	function ModuleFunction
	arg interface{}
}

func (mod *Module) Run(req http.Request) *http.Status {

	return mod.function(req, mod.arg)
}
