package tree

import 	(
	"sync"
	//
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core/http"
	openapiv2 "github.com/jsonrouter/core/openapi/v2"
	openapiv3 "github.com/jsonrouter/core/openapi/v3"
)

type Config struct {
	Log logging.Logger
	Spec interface{}
	// ProjectName is for App Engine apps
	ProjectName string
	CacheFiles bool
	ForcedTLS bool
	sync.RWMutex
}

func (config *Config) SpecV2() *openapiv2.Spec {
	return config.Spec.(*openapiv2.Spec)
}

func (config *Config) SpecV3() *openapiv3.Spec {
	return config.Spec.(*openapiv3.Spec)
}

func (config *Config) ServeSpec(req http.Request) *http.Status {
	return req.Respond(config.Spec)
}

func (config *Config) NoCache() {
	config.Lock()
	defer config.Unlock()
	config.CacheFiles = false
}

// block all non-https requests
func (config *Config) ForceTLS() {
	config.Lock()
	defer config.Unlock()
	config.ForcedTLS = true
}
