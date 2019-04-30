package evaluator

import (
	"strconv"

	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/object"
	"github.com/shozawa/monkey/token"
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
	case *ast.BoolLiteral:
		b, err := strconv.ParseBool(node.Value)
		if err != nil {
			return nil
		}
		return &object.Bool{Value: b}
	case *ast.BlockStatement:
		var result object.Object
		for _, stmt := range node.Statements {
			result = Eval(stmt, env)
		}
		return result
	case *ast.IfExpression:
		obj := Eval(node.Condition, env)
		b, ok := obj.(*object.Bool)
		if !ok {
			// TODO: report error if condition is not bool
			return nil
		}
		if b.Value {
			return Eval(node.Consequence, env)
		} else {
			return Eval(node.Alternative, env)
		}
	case *ast.Infix:
		switch node.Token.Type {
		case token.PLUS:
			integer := &object.Integer{}
			left, ok := Eval(node.Left, env).(*object.Integer)
			right, ok := Eval(node.Right, env).(*object.Integer)
			if !ok {
				// TODO: report error
			}
			integer.Value = left.Value + right.Value
			return integer
		case token.ASTERISK:
			integer := &object.Integer{}
			left, ok := Eval(node.Left, env).(*object.Integer)
			right, ok := Eval(node.Right, env).(*object.Integer)
			if !ok {
				// TODO: report error
			}
			integer.Value = left.Value * right.Value
			return integer
		}
	}
	return nil
}
