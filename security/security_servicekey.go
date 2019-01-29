package security

import (
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
//	"github.com/jsonrouter/core/openapi/v3"
)

// API key implementation

type ServiceKey struct {
	Key string
}

func (self *ServiceKey) Name() string {
	return "serviceKey"
}

func (self *ServiceKey) Validate(req http.Request) *http.Status {
	if req.GetRequestHeader("Authorization") == self.Key {
		return nil
	}
	return req.Respond(403, "serviceKey: FAILED TO VALIDATE, ACCESS DENIED!")
}

func (self *ServiceKey) DefinitionV2() *openapiv2.SecurityDefinition {
	return &openapiv2.SecurityDefinition{
		Type: self.Name(),
		Name: self.Name(),
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
