package config

import 	(
	"sync"
	//
	"github.com/jsonrouter/logging"
	openapitwo "github.com/jsonrouter/core/openapi/v2"
	openapithree "github.com/jsonrouter/core/openapi/v3"
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

func (config *Config) SpecV2() *openapitwo.APISpec {
	return config.Spec.(*openapitwo.APISpec)
}

func (config *Config) SpecV3() *openapithree.APISpec {
	return config.Spec.(*openapithree.APISpec)
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
