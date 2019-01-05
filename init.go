package core

import 	(
		"fmt"
		//
		"github.com/jsonrouter/core/config"
		"github.com/jsonrouter/core/tree"
		"github.com/jsonrouter/validation"
		)

// for internal debugging use
func debug(s string) { fmt.Println(s) }

var globalNode *tree.Node

func rootNode() *tree.Node {
	return &tree.Node{
		Headers: map[string]string{},
		Routes:	map[string]*tree.Node{},
		Methods: map[string]*tree.Handler{},
		RequestParams: map[string]interface{}{},
		Modules: []*tree.Module{},
		Validations: []*validation.Config{},
	}
}

func init() {

	globalNode = rootNode()
	globalNode.Config = &config.Config{
		CacheFiles: true,
	}

}

func Root() *tree.Node {

	return globalNode
}
