package tree

import 	(
	"github.com/jsonrouter/core/http"
)

type ModuleFunction func (req http.Request, arg interface{}) *http.Status

type Module struct {
	config *Config
	function ModuleFunction
	arg interface{}
}

// Run will run the module function 
func (mod *Module) Run(req http.Request) *http.Status {

	return mod.function(req, mod.arg)
}
