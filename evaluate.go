package slang

import (
	"fmt"
)

func evaluateListItems(lst List, env Env) ([]LangType, error) {
	lstLen := int(lst.Len())
	values := make([]LangType, lstLen)

	head := lst.head
	for n := 0; n < lstLen; n++ {
		item, err := Evaluate(head.value, env)
		if err != nil {
			return nil, err
		}
		values[n] = item
		head = head.next
	}

	return values, nil
}

func evaluateVectorItems(vec Vector, env Env) (Vector, error) {
	values := make(Vector, len(vec))
	for i, item := range vec {
		value, err := Evaluate(item, env)
		if err != nil {
			return Vector{}, err
		}
		values[i] = value
	}
	return values, nil
}

// evaluateSequence evaluates each expression in the given list except the last expression. The tail
// expression is returned to be evaluated.
func evaluateBodyTCO(lst List, env Env) (LangType, error) {
	tailN := lst.len - 1
	node := lst.head
	for n := 0; n < tailN; n++ {
		_, err := Evaluate(node.value, env)
		if err != nil {
			return nil, err
		}
		node = node.next
	}
	return lst.tail.value, nil
}

func evaluateExpr(expr LangType, env Env) (LangType, error) {
	switch t := expr.(type) {
	case Vector:
		return evaluateVectorItems(t, env)
	case Symbol:
		return env.Get(t)
	default:
		return t, nil
	}
}

// Evaluate evaluates an expression
func Evaluate(expr LangType, env Env) (LangType, error) {
	for {
		form, isList := expr.(List)
		if !isList {
			return evaluateExpr(expr, env)
		}

		if form.Len() == 0 {
			return expr, nil
		}

		operator := form.First()

		// peek at first object in list to determine evaluation path.
		// if first is a symbol, it can match against any of the special form cases below. A list
		// should be evaluated in the default case and any other type is not able to be applied.
		var first Symbol
		switch t := operator.(type) {
		case List:
			// just break. will eval and apply in default case
			break
		case Symbol:
			first = t
		default:
			return nil, fmt.Errorf("'%s' is not applicable", operator)
		}

		switch first {
		case "define":
			operands := form.Rest()

			if operands.Len() < 1 {
				return nil, fmt.Errorf("Invalid form for define")
			}

			defsym, isSymbol := operands.First().(Symbol)
			if !isSymbol {
				return nil, fmt.Errorf("First argument must be a symbol")
			}

			var defval LangType
			var err error
			if operands.Len() >= 3 {
				// syntactic sugar for a procedure definition
				// Usage: `(define <procedureName> [params...] body...)
				procdef := operands.Rest()
				params := procdef.First()
				body := procdef.Rest().(List)

				paramsVec, isVec := params.(Vector)
				if !isVec {
					return nil, fmt.Errorf("Second argument must be a vector for procedure definition")
				}

				defval, err = MakeLambda(&env, paramsVec, body)
			} else {
				defval, err = Evaluate(operands.Nth(1), env)
			}

			if err != nil {
				return nil, err
			}

			env.Define(defsym, defval)

			return defval, nil
		case "lambda":
			operands := form.Rest()

			// Usage: `(lambda [...params] ...body)`
			if operands.Len() < 2 {
				return nil, fmt.Errorf("Invalid number of arguments - expected at least 2 arguments")
			}

			params, isList := operands.First().(Vector)
			if !isList {
				return nil, fmt.Errorf("First argument to lambda must be a vector")
			}

			body := operands.Rest().(List)

			return MakeLambda(&env, params, body)
		case "quote":
			operands := form.Rest()

			if operands.Len() != 1 {
				return nil, fmt.Errorf("Invalid number of arguments - expected 1 argument")
			}
			return operands.First(), nil
		// Tail-call optimized paths
		case "begin":
			operands := form.Rest()

			if operands.Len() < 1 {
				return nil, fmt.Errorf("Invalid form for begin")
			}

			body := operands.Rest().(List)

			tail, err := evaluateBodyTCO(body, env)
			if err != nil {
				return nil, err
			}

			expr = tail
		case "if":
			operands := form.Rest()

			if operands.Len() < 2 || operands.Len() > 3 {
				return nil, fmt.Errorf("Invalid form for if")
			}

			predicate, err := Evaluate(operands.First(), env)
			if err != nil {
				return nil, err
			}

			result, isBool := predicate.(bool)
			if !isBool {
				return nil, fmt.Errorf(
					"If predicate must evaluate to either %t or %t", true, false)
			}

			if result == true {
				// TCO loop to evaluate consequent
				expr = operands.Nth(1)
			} else if operands.Len() == 3 {
				// TCO loop to evaluate alternative if provided
				expr = operands.Nth(2)
			} else {
				return false, nil
			}
			// loop to evaluate consequent or alternative
		default:
			// evaluate all items in form (list); first should be an applicable procedure/lambda
			values, err := evaluateListItems(form, env)
			if err != nil {
				return nil, err
			}

			procedure := values[0]
			args := values[1:]

			// apply lambda or subroutine
			if lambda, isLambda := procedure.(Lambda); isLambda {
				if len(args) != len(lambda.params) {
					return nil, fmt.Errorf("Incorrect number of arguments to apply lambda")
				}

				// modify env to reference the lambda's closure
				env = lambda.closure

				// perform left to right bindings of lambda arguments
				for i, bindValue := range args {
					bindSymbol := lambda.params[i].(Symbol)
					env.Define(bindSymbol, bindValue)
				}

				// evaluate all items in body except the final expression
				// last expression is set in for next TCO loop iteration to evaluate it
				expr, err = evaluateBodyTCO(lambda.body, env)
				if err != nil {
					return nil, err
				}
			} else if subr, isSubroutine := procedure.(Subroutine); isSubroutine {
				return subr.Apply(args...)
			} else {
				return nil, fmt.Errorf("'%s' is not applicable", procedure)
			}
		}
	}
}
