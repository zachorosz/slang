package slang

import "fmt"

// Env environment with scopes and reference to enclosing frame
type Env struct {
	EnclosingFrame *Env
	Frame          map[Symbol]LangType
}

// Get performs a symbol lookup. If the symbol key is not present in the current
// or any enclosing frame, an undefined symbol error is returned.
func (env *Env) Get(symbol Symbol) (LangType, error) {
	currentEnv := env
	for currentEnv != nil {
		if value, exists := currentEnv.Frame[symbol]; exists {
			return value, nil
		}
		currentEnv = currentEnv.EnclosingFrame
	}
	return nil, fmt.Errorf("Symbol '%s' is undefined", symbol)
}

// Define adds a new symbol definition to the environment. If symbol is already defined, an error is
// returned; use the Mutate method to change the definition of a symbol.
func (env *Env) Define(symbol Symbol, value LangType) error {
	if _, exists := env.Frame[symbol]; !exists {
		env.Frame[symbol] = value
		return nil
	}
	return fmt.Errorf("Symbol '%s' is already defined", symbol)
}

// Mutate mutates a symbol definition. Mutating an undefined symbol is not allowed; use the Set
// method to add a new symbol definition.
func (env *Env) Mutate(symbol Symbol, value LangType) error {
	if _, exists := env.Frame[symbol]; exists {
		env.Frame[symbol] = value
		return nil
	}
	return fmt.Errorf("Symbol '%s' is undefined", symbol)
}

var env = map[Symbol]LangType{
	// Predicates
	Symbol("list?"): Subroutine{
		Name: "list?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return ListP(args[0]), nil
		},
	},
	Symbol("nil?"): Subroutine{
		Name: "nil?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return NilP(args[0]), nil
		},
	},
	Symbol("number?"): Subroutine{
		Name: "number?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return NumberP(args[0]), nil
		},
	},
	Symbol("procedure?"): Subroutine{
		Name: "procedure?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return ProcedureP(args[0]), nil
		},
	},
	Symbol("seq?"): Subroutine{
		Name: "seq?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return SequenceP(args[0]), nil
		},
	},
	Symbol("string?"): Subroutine{
		Name: "string?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return StringP(args[0]), nil
		},
	},
	Symbol("symbol?"): Subroutine{
		Name: "number?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return SymbolP(args[0]), nil
		},
	},
	Symbol("vec?"): Subroutine{
		Name: "vec?",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			return VectorP(args[0]), nil
		},
	},
	Symbol("not"): Subroutine{
		Name: "not",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			pred, isBool := args[0].(bool)
			if !isBool {
				return nil, fmt.Errorf("Attempt to negate non-bool - %s", args[0])
			}
			return Not(pred), nil
		},
	},
	// Operators
	Symbol(">"): Subroutine{
		Name: ">",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			lhs, ok := args[0].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Greater than equal to operator is not defined on %T", args[0])
			}
			rhs, ok := args[1].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Greater than operator is not defined on %T", args[1])
			}

			return Gt(lhs, rhs)
		},
	},
	Symbol("<"): Subroutine{
		Name: "<",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			lhs, ok := args[0].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Less than operator is not defined on %T", args[0])
			}
			rhs, ok := args[1].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Less thanoperator is not defined on %T", args[1])
			}

			return Lt(lhs, rhs)
		},
	},
	Symbol(">="): Subroutine{
		Name: ">=",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			lhs, ok := args[0].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Greater than or equal to operator is not defined on %T", args[0])
			}
			rhs, ok := args[1].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Greater than or equal to operator is not defined on %T", args[1])
			}

			return Gte(lhs, rhs)
		},
	},
	Symbol("<="): Subroutine{
		Name: "<=",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			lhs, ok := args[0].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Less than or equal to operator is not defined on %T", args[0])
			}
			rhs, ok := args[1].(Comparable)
			if !ok {
				return nil, fmt.Errorf("Less than or equal to operator is not defined on %T", args[1])
			}

			return Lte(lhs, rhs)
		},
	},
	Symbol("+"): Subroutine{
		Name: "+",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			x, ok := args[0].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Addition operator is not defined on %T", args[0])
			}
			y, ok := args[1].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Addition operator is not defined on %T", args[1])
			}

			return Add(x, y)
		},
	},
	Symbol("-"): Subroutine{
		Name: "-",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			x, ok := args[0].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Subtraction operator is not defined on %T", args[0])
			}
			y, ok := args[1].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Subtraction operator is not defined on %T", args[1])
			}

			return Sub(x, y)
		},
	},
	Symbol("*"): Subroutine{
		Name: "*",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			x, ok := args[0].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Multiplication operator is not defined on %T", args[0])
			}
			y, ok := args[1].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Multiplication operator is not defined on %T", args[1])
			}

			return Mul(x, y)
		},
	},
	Symbol("/"): Subroutine{
		Name: "/",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			x, ok := args[0].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Division operator is not defined on %T", args[0])
			}
			y, ok := args[1].(Algebraic)
			if !ok {
				return nil, fmt.Errorf("Division operator is not defined on %T", args[1])
			}

			return Div(x, y)
		},
	},
	Symbol("%"): Subroutine{
		Name: "%",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			x, ok := args[0].(Number)
			if !ok {
				return nil, fmt.Errorf("Modulo operator is not defined on %T", args[0])
			}
			y, ok := args[1].(Number)
			if !ok {
				return nil, fmt.Errorf("Modulo operator is not defined on %T", args[1])
			}

			return Mod(x, y)
		},
	},
	// Sequence functions
	Symbol("append"): Subroutine{
		Name: "append",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			seq, isSeq := args[0].(Sequence)
			if !isSeq {
				return nil, fmt.Errorf("%s is not a sequence", args[0])
			}
			return seq.Append(args[1]), nil
		},
	},
	Symbol("nth"): Subroutine{
		Name: "nth",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
			}
			seq, isSeq := args[0].(Sequence)
			if !isSeq {
				return nil, fmt.Errorf("%s is not a sequence", args[0])
			}
			n, isNumber := args[1].(Number)
			if !isNumber {
				return nil, fmt.Errorf("%s is not a valid number", args[1])
			}
			return Nth(seq, n)
		},
	},
	Symbol("len"): Subroutine{
		Name: "len",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
			}
			seq, isSeq := args[0].(Sequence)
			if !isSeq {
				return nil, fmt.Errorf("%s is not a sequence", args[0])
			}
			return seq.Len(), nil
		},
	},
	Symbol("list"): Subroutine{
		Name: "list",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) < 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected at least 1 argument")
			}
			return MakeList(args[0], args[1:]), nil
		},
	},
	Symbol("vec"): Subroutine{
		Name: "vec",
		Func: func(args ...LangType) (LangType, error) {
			if len(args) < 1 {
				return nil, fmt.Errorf("Incorrect number of arguments - expected at least 1 argument")
			}
			return MakeVector(args[0], args[1:]), nil
		},
	},
}

// StandardEnv is an environment with builtin subroutines
var StandardEnv = Env{Frame: env}
