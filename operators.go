package slang

// Algebraic is an interface for algebraic operations. Types that implement this interface can
// overload the behavior of an "algebraic" operator (+, -, *, /). Only Algebraic types can be added,
// subtracted, multiplied, and divided by other Algrebraic types.
type Algebraic interface {
	Plus(obj Algebraic) (Algebraic, error)
	Minus(obj Algebraic) (Algebraic, error)
	Multiply(obj Algebraic) (Algebraic, error)
	Divide(obj Algebraic) (Algebraic, error)
}

// Comparable is an interface for comparision operations. Types that implement this interface can
// overload the behavior of a comparision operator (>, <, >=, <=). Only Comparable types can be
// compared against other Comparable types.
type Comparable interface {
	GreaterThan(obj Comparable) (bool, error)
	LessThan(obj Comparable) (bool, error)
	GreaterThanOrEqualTo(obj Comparable) (bool, error)
	LessThanOrEqualTo(obj Comparable) (bool, error)
}

// Gt is a conditional operator that returns true if lhs is greater than rhs.
// Usage: `(> x y)`
func Gt(lhs Comparable, rhs Comparable) (LangType, error) {
	return lhs.GreaterThan(rhs)
}

// Lt is a conditional operator that returns true if lhs is less than rhs.
// Usage: `(< x y)`
func Lt(lhs, rhs Comparable) (LangType, error) {
	return lhs.LessThan(rhs)
}

// Gte is a conditional operator that returns true if lhs is greater than or equal to rhs.
// Usage: `(>= x y)`
func Gte(lhs, rhs Comparable) (LangType, error) {
	return lhs.GreaterThanOrEqualTo(rhs)
}

// Lte is a conditional operator that returns true if lhs is less than or equal to than rhs.
// Usage: `(<= x y)`
func Lte(lhs, rhs Comparable) (LangType, error) {
	return lhs.LessThanOrEqualTo(rhs)
}

// Add is the addition operator.
// Usage: `(+ x y)`
func Add(x, y Algebraic) (LangType, error) {
	return x.Plus(y)
}

// Sub is the subtraction operator.
// Usage: `(- x y)`
func Sub(x, y Algebraic) (LangType, error) {
	return x.Minus(y)
}

// Mul is the multiplication operator.
// Usage: `(* x y)`
func Mul(x, y Algebraic) (LangType, error) {
	return x.Multiply(y)
}

// Div is the division operator.
// Usage: `(/ x y)`
func Div(x, y Algebraic) (LangType, error) {
	return x.Divide(y)
}

// Mod is the modulus operator. Modulo of two numbers returns the remainder of the quotient. This
// operator cannot be overloaded.
// Usage: `(% x y)`
func Mod(x, y Number) (LangType, error) {
	return Number(int(x) % int(y)), nil
}
