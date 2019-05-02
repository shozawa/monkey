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
		return evalStatements(node.Statements, env)
	case *ast.LetStatement:
		env.Set(node.Name.Value, Eval(node.Value, env))
		return nil
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.Identifier:
		value, ok := env.Get(node.Value)
		if !ok {
			return nil
		}
		return value
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
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Body: body, Env: env}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		args := evalExpressions(node.Arguments, env)
		return applyFunction(function, args)
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

func evalStatements(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)
	}

	return result
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	// TODO: Error hadnling
	var result []object.Object
	for _, exp := range exps {
		evaluated := Eval(exp, env)
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return nil
	}
	extendEnv := extendFunctionEnv(function, args)
	return Eval(function.Body, extendEnv)
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}
