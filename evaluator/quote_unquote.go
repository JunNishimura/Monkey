package evaluator

import (
	"github.com/JunNishimura/Monkey/ast"
	"github.com/JunNishimura/Monkey/object"
)

func quote(node ast.Node) *object.Quote {
	return &object.Quote{Node: node}
}
