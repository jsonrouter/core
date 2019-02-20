package tree

import (
	"github.com/jsonrouter/core/http"
)

// HEAD allows POST requests to the node's handler
func (node *Node) HEAD(functions ...interface{}) *Handler {

	handler := node.newHandler("HEAD")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("HEAD", handler)

	return handler
}

// GET allows GET requests to the node's handler
func (node *Node) GET(functions ...interface{}) *Handler {

	handler := node.newHandler("GET")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("GET", handler)

	return handler
}

// POST allows POST requests to the node's handler
func (node *Node) POST(functions ...interface{}) *Handler {

	handler := node.newHandler("POST")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("POST", handler)

	return handler
}

// PUT allows PUT requests to the node's handler
func (node *Node) PUT(functions ...interface{}) *Handler {

	handler := node.newHandler("PUT")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("PUT", handler)

	return handler
}

// DELETE allows DELETE requests to the node's handler
func (node *Node) DELETE(functions ...interface{}) *Handler {

	handler := node.newHandler("DELETE")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("DELETE", handler)

	return handler
}

// PATCH allows PATCH requests to the node's handler
func (node *Node) PATCH(functions ...interface{}) *Handler {

	handler := node.newHandler("PATCH")

	if len(functions) > 0 { handler.Function = functions[0].(func(http.Request) *http.Status) }

	node.addHandler("PATCH", handler)

	return handler
}
