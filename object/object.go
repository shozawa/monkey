package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
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
