package openapiv3

import (
	"reflect"
)

func New(host, serviceName string) *Spec {
	return &Spec{
		OpenAPI: "3.0.0",
		Info: &Info{
			Title: serviceName,
		},
		Host: host,
		Paths: map[string]Path{},
		Schemes: []string{"http"},
		Consumes: []string{"application/json"},
		Produces: []string{"application/json"},
		Definitions: map[string]*Definition{},
		Components: &Components{
			SecuritySchemes: map[string]*SecuritySchemeObject{},
		},
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
	}
	panic("openapi: UNMAPPED TYPE "+reflect.TypeOf(x).String())
	return ""
}
