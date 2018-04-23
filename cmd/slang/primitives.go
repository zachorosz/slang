package main

import (
	"fmt"

	"github.com/zachorosz/slang"
)

// Primitives is a map with applications of slang primitives.
//
// This is the core package for slang environments.
var Primitives = map[string]func(...slang.LangType) (slang.LangType, error){
	"list?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.ListP(args[0]), nil
	},
	"nil?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.NilP(args[0]), nil
	},
	"number?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.NumberP(args[0]), nil
	},
	"procedure?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.ProcedureP(args[0]), nil
	},
	"seq?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.SequenceP(args[0]), nil
	},
	"string?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.StringP(args[0]), nil
	},
	"symbol?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.SymbolP(args[0]), nil
	},
	"vec?": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		return slang.VectorP(args[0]), nil
	},
	">": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		lhs, ok := args[0].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Greater than equal to operator is not defined on %T", args[0])
		}
		rhs, ok := args[1].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Greater than operator is not defined on %T", args[1])
		}

		return slang.Gt(lhs, rhs)
	},
	"<": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		lhs, ok := args[0].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Less than operator is not defined on %T", args[0])
		}
		rhs, ok := args[1].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Less thanoperator is not defined on %T", args[1])
		}

		return slang.Lt(lhs, rhs)
	},
	">=": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		lhs, ok := args[0].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Greater than or equal to operator is not defined on %T", args[0])
		}
		rhs, ok := args[1].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Greater than or equal to operator is not defined on %T", args[1])
		}

		return slang.Gte(lhs, rhs)
	},
	"<=": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		lhs, ok := args[0].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Less than or equal to operator is not defined on %T", args[0])
		}
		rhs, ok := args[1].(slang.Comparable)
		if !ok {
			return nil, fmt.Errorf("Less than or equal to operator is not defined on %T", args[1])
		}

		return slang.Lte(lhs, rhs)
	},
	"+": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		x, ok := args[0].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Addition operator is not defined on %T", args[0])
		}
		y, ok := args[1].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Addition operator is not defined on %T", args[1])
		}

		return slang.Add(x, y)
	},
	"-": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		x, ok := args[0].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Subtraction operator is not defined on %T", args[0])
		}
		y, ok := args[1].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Subtraction operator is not defined on %T", args[1])
		}

		return slang.Sub(x, y)
	},
	"*": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		x, ok := args[0].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Multiplication operator is not defined on %T", args[0])
		}
		y, ok := args[1].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Multiplication operator is not defined on %T", args[1])
		}

		return slang.Mul(x, y)
	},
	"/": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		x, ok := args[0].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Division operator is not defined on %T", args[0])
		}
		y, ok := args[1].(slang.Algebraic)
		if !ok {
			return nil, fmt.Errorf("Division operator is not defined on %T", args[1])
		}

		return slang.Div(x, y)
	},
	"%": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		x, ok := args[0].(slang.Number)
		if !ok {
			return nil, fmt.Errorf("Modulo operator is not defined on %T", args[0])
		}
		y, ok := args[1].(slang.Number)
		if !ok {
			return nil, fmt.Errorf("Modulo operator is not defined on %T", args[1])
		}

		return slang.Mod(x, y)
	},
	"append": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		seq, isSeq := args[0].(slang.Sequence)
		if !isSeq {
			return nil, fmt.Errorf("%s is not a sequence", args[0])
		}
		return seq.Append(args[1]), nil
	},
	"first": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 arguments")
		}
		seq, isSeq := args[0].(slang.Sequence)
		if !isSeq {
			return nil, fmt.Errorf("%s is not a sequence", args[0])
		}
		return seq.First(), nil
	},
	"rest": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 arguments")
		}
		seq, isSeq := args[0].(slang.Sequence)
		if !isSeq {
			return nil, fmt.Errorf("%s is not a sequence", args[0])
		}
		return seq.Rest(), nil
	},
	"nth": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 2 arguments")
		}
		seq, isSeq := args[0].(slang.Sequence)
		if !isSeq {
			return nil, fmt.Errorf("%s is not a sequence", args[0])
		}
		n, isNumber := args[1].(slang.Number)
		if !isNumber {
			return nil, fmt.Errorf("%s is not a valid number", args[1])
		}
		return slang.Nth(seq, n)
	},
	"len": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected 1 argument")
		}
		seq, isSeq := args[0].(slang.Sequence)
		if !isSeq {
			return nil, fmt.Errorf("%s is not a sequence", args[0])
		}
		return seq.Len(), nil
	},
	"list": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected at least 1 argument")
		}
		return slang.MakeList(args[0], args[1:]...), nil
	},
	"vec": func(args ...slang.LangType) (slang.LangType, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("Incorrect number of arguments - expected at least 1 argument")
		}
		return slang.MakeVector(args[0], args[1:]...), nil
	},
}
