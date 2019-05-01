package object

import (
	"fmt"

	"github.com/shozawa/monkey/ast"
)

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	BOOL_OBJ     = "BOOL"
	FUNCTION_OBJ = "FUNCTION"
)

type Environment struct {
	store map[string]Object
}

func (e *Environment) Set(k string, v Object) {
	e.store[k] = v
}

func (e *Environment) Get(k string) Object {
	return e.store[k]
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
