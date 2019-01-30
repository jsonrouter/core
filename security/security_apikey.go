package security

import (
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
//	"github.com/jsonrouter/core/openapi/v3"
)

// API key implementation

type Security_ApiKey struct {
}

func (apiKey *Security_ApiKey) Name() string {
	return "apiKey"
}

func (apiKey *Security_ApiKey) Validate(req http.Request) *http.Status {
	return nil
}

func (apiKey *Security_ApiKey) DefinitionV2() *openapiv2.SecurityDefinition {
	return &openapiv2.SecurityDefinition{
		Type: "apiKey",
		Name: "apikey",
		In: "header",
	}
}

/*
func (apiKey *SecurityApiKey) DefinitionV3() *openapiv3.SecurityDefinition {
	return &openapi.SecurityDefinition{
		Type: "apikey",
		Name: "apikey",
		In: "header",
	}
}
*/