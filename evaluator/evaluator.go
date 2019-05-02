package evaluator

import (
	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/object"
)

var (
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
	NULL  = &object.Null{}
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
		return strToBoolObject(node.Value)
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
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node.Operator, left, right)
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

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	default:
		return NULL
	}
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

func strToBoolObject(str string) *object.Bool {
	if str == "true" {
		return TRUE
	} else {
		return FALSE
	}
}
