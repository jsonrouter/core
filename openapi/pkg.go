package openapi

import (
	"fmt"
	"reflect"
	//
	"github.com/jsonrouter/core/openapi/v2"
	"github.com/jsonrouter/core/openapi/v3"
)

func ValidSpec(x interface{}) error {

	switch x.(type) {

	case *openapiv2.Spec:
		return nil

	case *openapiv3.Spec:
		return nil

	}

	return fmt.Errorf("OpenAPI spec provided was not correct type! Consider using a type other than: %s", reflect.TypeOf(x).String())
}

func pointerFloat64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}
