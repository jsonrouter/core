package openapiv2

import (
	"reflect"
)

// New returns a new spec object
func New(host, serviceName string) *Spec {
	return &Spec{
		Swagger: "2.0",
		Info: &Info{
			Title: serviceName,
		},
		Host: host,
		Paths: map[string]Path{},
		Schemes: []string{"http"},
		Consumes: []string{"application/json"},
		Produces: []string{"application/json"},
		Definitions: map[string]*Definition{},
		SecurityDefinitions: map[string]*SecurityDefinition{},
	}
}

func Type(x interface{}) string {
	switch reflect.TypeOf(x).String() {
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "int64", "int":
		return "integer"
	case "float32", "float64":
		return "number"
	case "map[string]interface {}":
		return "object"
	case "[]string":
		return "array"
	case "[]interface {}":
		return "array"
	}
	panic("openapi: UNMAPPED TYPE "+reflect.TypeOf(x).String())
	return ""
}
