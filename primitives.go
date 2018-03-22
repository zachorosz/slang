package slang

import (
	"fmt"
	"strings"
)

// LangType base type to interface with the language types of slang.
type LangType interface{}

// Symbol slang symbol type
type Symbol string

// Number is a slang number type.
type Number float64

func (n Number) String() string {
	return fmt.Sprintf("%v", float64(n))
}

// Plus returns the sum of two numbers or the concatenation of the number and a string.
func (n Number) Plus(obj Algebraic) (Algebraic, error) {
	switch t := obj.(type) {
	case Number:
		return n + t, nil
	case Str:
		return Str(fmt.Sprint(n)) + t, nil
	default:
		return nil, fmt.Errorf("Cannot add number and %T", t)
	}
}

// Minus returns the difference of two numbers.
func (n Number) Minus(obj Algebraic) (Algebraic, error) {
	switch t := obj.(type) {
	case Number:
		return n - t, nil
	default:
		return nil, fmt.Errorf("Cannot subtract number and %T", t)
	}
}

// Multiply returns the product of two numbers.
func (n Number) Multiply(obj Algebraic) (Algebraic, error) {
	switch t := obj.(type) {
	case Number:
		return n * t, nil
	default:
		return nil, fmt.Errorf("Cannot multiply number and %T", t)
	}
}

// Divide returns the quotient of two numbers.
func (n Number) Divide(obj Algebraic) (Algebraic, error) {
	switch t := obj.(type) {
	case Number:
		return n / t, nil
	default:
		return nil, fmt.Errorf("Cannot divide number and %T", t)
	}
}

// GreaterThan returns true if number is greater than the given number.
func (n Number) GreaterThan(obj Comparable) (bool, error) {
	switch t := obj.(type) {
	case Number:
		return n > t, nil
	default:
		return false, fmt.Errorf("Cannot compare number and %T", t)
	}
}

// LessThan returns true if number is less than the given number.
func (n Number) LessThan(obj Comparable) (bool, error) {
	switch t := obj.(type) {
	case Number:
		return n < t, nil
	default:
		return false, fmt.Errorf("Cannot compare number and %T", t)
	}
}

// GreaterThanOrEqualTo returns true if number is greater than or equal to the given number.
func (n Number) GreaterThanOrEqualTo(obj Comparable) (bool, error) {
	switch t := obj.(type) {
	case Number:
		return n >= t, nil
	default:
		return false, fmt.Errorf("Cannot compare number and %T", t)
	}
}

// LessThanOrEqualTo returns true if number is less than or equal to the given number.
func (n Number) LessThanOrEqualTo(obj Comparable) (bool, error) {
	switch t := obj.(type) {
	case Number:
		return n <= t, nil
	default:
		return false, fmt.Errorf("Cannot compare number and %T", t)
	}
}

// Str is a slang string type.
type Str string

func (s Str) String() string {
	return string("\"" + s + "\"")
}

// Plus returns a new, concatenated string.
// Concatenating a non-string uses the default format verb from the fmt package.
func (s Str) Plus(obj Algebraic) (Algebraic, error) {
	switch t := obj.(type) {
	case Str:
		return s + t, nil
	default:
		return s + Str(fmt.Sprint(t)), nil
	}
}

// Minus returns an invalid operation error when an attempt to subtract a string occurs.
func (s Str) Minus(obj Algebraic) (Algebraic, error) {
	return nil, fmt.Errorf("Subtraction operator is not defined on string")
}

// Multiply returns a new string with repeat count copies.
// If the repeat count is a negative number the absolute value is used.
func (s Str) Multiply(obj Algebraic) (Algebraic, error) {
	if obj == nil {
		return nil, fmt.Errorf("Repeat count expected")
	}
	switch t := obj.(type) {
	case Number:
		if t < 0 {
			t *= -1
		}
		repeated := strings.Repeat(string(s), int(t))
		return Str(repeated), nil
	default:
		return nil, fmt.Errorf("Repeat expects a number")
	}
}

// Divide returns an invalid operation error when an attempt to divide a string occurs.
func (s Str) Divide(obj Algebraic) (Algebraic, error) {
	return nil, fmt.Errorf("Division operator is not defined on string")
}

// Modulo returns an invalid operation error when an attempt to mod a string occurs.
func (s Str) Modulo(obj Algebraic) (Algebraic, error) {
	return nil, fmt.Errorf("Modulo operator is not defined on string")
}

// Subroutine a slang function that is implemented in the host language, Go!
type Subroutine struct {
	Func func(...LangType) (LangType, error)
}

// Apply applies arguments to the subroutine and returns the evaluation.
func (subr Subroutine) Apply(args ...LangType) (LangType, error) {
	return subr.Func(args...)
}

// Lambda a slang function type. Use MakeLambda to construct a Lambda.
type Lambda struct {
	params Vector
	body   List
	env    Env
}

func (lambda Lambda) String() string {
	return fmt.Sprintf("<procedure>")
}

// NilP is a predicate that returns true if object is nil or an empty list.
// Usage: `(nil? x)`
func NilP(x LangType) bool {
	if x == nil {
		return true
	}
	if list, isList := x.(List); isList {
		return list.Len() == 0
	}
	return false
}

// NumberP returns true if object is a number.
// Usage: `(procedure? x)`
func NumberP(x LangType) bool {
	_, isNumber := x.(Number)
	return isNumber
}

// ProcedureP returns true if object is a slang lambda or Go subroutine.
// Usage: `(procedure? x)`
func ProcedureP(x LangType) bool {
	switch x.(type) {
	case Lambda, Subroutine:
		return true
	default:
		return false
	}
}

// StringP returns true if object is a string.
// Usage: `(string? x)`
func StringP(x LangType) bool {
	_, isString := x.(string)
	return isString
}

// SymbolP returns true if object is a Symbol.
// Usage: `(symbol? x)`
func SymbolP(x LangType) bool {
	_, isSymbol := x.(Symbol)
	return isSymbol
}

// MakeLambda makes a new Lambda function with N-arity. When applied, arguments are bound to its
// environment frame (A.K.A. closure) and the body is evaluated. The evaluation of the final, or
// only, expression in the body is used as the return value.
// Usage: `(lambda [params...] body...)`
func MakeLambda(env Env, params Vector, body List) (Lambda, error) {
	if body.Len() == 0 {
		return Lambda{}, fmt.Errorf("Lambda body expected")
	}

	return Lambda{
		params: params,
		body:   body,
		env:    env,
	}, nil
}
