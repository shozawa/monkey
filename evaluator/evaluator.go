package evaluator

import (
	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		var result object.Object
		for _, stmt := range node.Statements {
			result = Eval(stmt, env)
		}
		return result
	case *ast.LetStatement:
		env.Set(node.Name.Value, Eval(node.Value, env))
		return nil
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.Identifier:
		return env.Get(node.Value)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
