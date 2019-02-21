package tree

import 	(
	"sync"
	//"time"
	"encoding/json"
	//"fmt"
	"github.com/jsonrouter/logging"
	"github.com/jsonrouter/core/http"
	openapiv2 "github.com/jsonrouter/core/openapi/v2"
	openapiv3 "github.com/jsonrouter/core/openapi/v3"
	"github.com/jsonrouter/core/metrics"
)

type Config struct {
	Log logging.Logger
	Spec interface{}
	// ProjectName is for App Engine apps
	ProjectName string
	CacheFiles bool
	ForcedTLS bool
	sync.RWMutex
	MetResults map[string]interface{}
	Metrics metrics.Metrics
}

// SpecV2 exposes the spec object as a V2
func (config *Config) SpecV2() *openapiv2.Spec {
	return config.Spec.(*openapiv2.Spec)
}

// SpecV3 exposes the spec object as a V3
func (config *Config) SpecV3() *openapiv3.Spec {
	return config.Spec.(*openapiv3.Spec)
}

// SpecHandler serves the raw spec to a HTTP request.
func (config *Config) SpecHandler(req http.Request) *http.Status {
	return req.Respond(config.Spec)
}

// MetricsHandler serves the raw metrics to a HTTP request.
func (config *Config) MetricsHandler(req http.Request) *http.Status {

	res, err := json.Marshal(config.Metrics.Results)
	if err != nil {
		return req.Fail()
	}

	return req.Respond(res)
}

// NoCache sets the no-cache part of the router config to true. This stops any static files being cached.
func (config *Config) NoCache() {
	config.Lock()
	defer config.Unlock()
	config.CacheFiles = false
}

// ForceTLS blocks all non-https requests
func (config *Config) ForceTLS() {
	config.Lock()
	defer config.Unlock()
	config.ForcedTLS = true
}
