package slang

import (
	"fmt"
)

// Env environment with scopes and reference to enclosing frame
type Env struct {
	enclosingFrame *Env
	frame          map[Symbol]LangType
}

// Get performs a symbol lookup. If the symbol key is not present in the current
// or any enclosing frame, an undefined symbol error is returned.
func (env *Env) Get(symbol Symbol) (LangType, error) {
	currentEnv := env
	for currentEnv != nil {
		if value, exists := currentEnv.frame[symbol]; exists {
			return value, nil
		}
		currentEnv = currentEnv.enclosingFrame
	}
	return nil, fmt.Errorf("Symbol '%s' is undefined", symbol)
}

// Define adds a new symbol definition to the environment. If symbol is already defined, an error is
// returned; use the Mutate method to change the definition of a symbol.
func (env *Env) Define(symbol Symbol, value LangType) error {
	if _, exists := env.frame[symbol]; !exists {
		env.frame[symbol] = value
		return nil
	}
	return fmt.Errorf("Symbol '%s' is already defined", symbol)
}

// Mutate mutates a symbol definition. Mutating an undefined symbol is not allowed; use the Set
// method to add a new symbol definition.
func (env *Env) Mutate(symbol Symbol, value LangType) error {
	if _, exists := env.frame[symbol]; exists {
		env.frame[symbol] = value
		return nil
	}
	return fmt.Errorf("Symbol '%s' is undefined", symbol)
}

// UseSubrPackage loads a
func (env *Env) UseSubrPackage(pkgName string,
	pkg map[string]func(...LangType) (LangType, error)) error {

	for k, v := range pkg {
		if err := env.Define(Symbol(k), Subroutine{v}); err != nil {
			return err
		}
	}
	return nil
}

// MakeEnv constructs an empty environment.
func MakeEnv() Env {
	frame := map[Symbol]LangType{}
	return Env{frame: frame}
}

// MakeClosure encloses an environment over a new environment.
func MakeClosure(enclosingEnv *Env) Env {
	env := MakeEnv()
	env.enclosingFrame = enclosingEnv
	return env
}
