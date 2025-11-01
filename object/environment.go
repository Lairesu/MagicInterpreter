// --------------------------
// Creation of Environment
// --------------------------

// At its heart the environment is a hash map, associates strings with objects
// Environment, a thin wrapper around a map which is use to keep track of value
// by associating them with a name

package object


func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func NewEnvironment() *Environment {
	s := make (map[string]Object)
	return &Environment{store: s, outer:nil}
}

type Environment struct {
	store map[string]Object

	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return  obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}



