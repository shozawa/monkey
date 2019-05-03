package object

import (
	"fmt"

	"github.com/shozawa/monkey/ast"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOL_OBJ         = "BOOLEAN"
	FUNCTION_OBJ     = "FUNCTION"
	NULL_OBJ         = "NULL"
	ERROR_OBJ        = "ERROR"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnv()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Set(k string, v Object) {
	e.store[k] = v
}

func (e *Environment) Get(k string) (Object, bool) {
	obj, ok := e.store[k]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(k)
	}
	return obj, ok
}

func NewEnv() *Environment {
	store := make(map[string]Object)
	return &Environment{store: store}
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Bool struct {
	Value bool
}

func (b *Bool) Type() ObjectType { return BOOL_OBJ }
func (b *Bool) Inspect() string  { return fmt.Sprintf("%v", b.Value) }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string  { return "TODO" }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (r *ReturnValue) Inspect() string  { return r.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
