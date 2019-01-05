package openapi

import (
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
)

type Spec interface {
	Build(handlerMap map[string]*tree.Handler) *http.Status
}
