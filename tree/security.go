package tree

import (
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/openapi/v2"
//	"github.com/jsonrouter/core/openapi/v3"
)

type SecurityModule interface {
	Name() string
	Validate(http.Request) *http.Status
	DefinitionV2() *openapiv2.SecurityDefinition
//	DefinitionV3() *openapiv3.SecurityDefinition
}
