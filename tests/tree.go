package tests

import (
	"github.com/jsonrouter/core/tree"
	"github.com/jsonrouter/validation"
)

func testTree(root *tree.Node) {

	api := root.Add(
		CONST_SPEC_BASEPATH,
	).Add(
		"test",
	).SetHeaders(
		map[string]interface{}{
			"Authorisation": "hello",
		},
	)

	api.GET(
		dummyHandler,
	).Description(
		"This is a GET endpoint!",
	).SetHeaders(
		map[string]interface{}{
			"Authorisation2": "hello",
		},
	).Response(
		TestObject{},
	)

	api.POST(
		dummyHandler,
	).Description(
		"This is a POST endpoint!",
	).Required(
		validation.Payload{
			"hello": validation.String(10, 20).Description("The hellos!").Default("Helloy"),
		},
	).Optional(
		validation.Payload{
			"world": validation.Int().Description("The worlds!").Default(2),
		},
	).Response(
		TestObject{},
	)

	apiResource := api.Add("/resource").Param(
		validation.String(1, 64).Description("The id of the user."),
		"id",
	)

		apiResource.GET(
			dummyHandler,
		).Description(
			"Handles access to the resource",
		).Response(
			TestObject{},
		)

}
