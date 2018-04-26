package slang

import (
	"testing"
)

func TestGt(t *testing.T) {
	cases := []struct {
		lhs, rhs Comparable
		want     bool
	}{
		{Number(1), Number(0), true},
		{Number(0), Number(1), false},
		{Number(1), Number(1), false},
	}

	for _, c := range cases {
		got, _ := Gt(c.lhs, c.rhs)
		if got != c.want {
			t.Errorf("Gt(%q, %q) == %t, want %t", c.lhs, c.rhs, got.(bool), c.want)
		}
	}
}

func TestLt(t *testing.T) {
	cases := []struct {
		lhs, rhs Comparable
		want     bool
	}{
		{Number(1), Number(0), false},
		{Number(0), Number(1), true},
		{Number(1), Number(1), false},
	}

	for _, c := range cases {
		got, _ := Lt(c.lhs, c.rhs)
		if got != c.want {
			t.Errorf("Lt(%q, %q) == %t, want %t", c.lhs, c.rhs, got.(bool), c.want)
		}
	}
}

func TestGte(t *testing.T) {
	cases := []struct {
		lhs, rhs Comparable
		want     bool
	}{
		{Number(1), Number(0), true},
		{Number(0), Number(1), false},
		{Number(1), Number(1), true},
	}

	for _, c := range cases {
		got, _ := Gte(c.lhs, c.rhs)
		if got != c.want {
			t.Errorf("Gte(%q, %q) == %t, want %t", c.lhs, c.rhs, got.(bool), c.want)
		}
	}
}

func TestLte(t *testing.T) {
	cases := []struct {
		lhs, rhs Comparable
		want     bool
	}{
		{Number(1), Number(0), false},
		{Number(0), Number(1), true},
		{Number(1), Number(1), true},
	}

	for _, c := range cases {
		got, _ := Lte(c.lhs, c.rhs)
		if got != c.want {
			t.Errorf("Lte(%q, %q) == %t, want %t", c.lhs, c.rhs, got.(bool), c.want)
		}
	}
}

func TestAddNumbers(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Number
	}{
		{Number(1), Number(0), Number(1)},
		{Number(-1), Number(1), Number(0)},
		{Number(0.5), Number(0.25), Number(0.75)},
		{Number(1.0), Number(1), Number(2)},
	}

	for _, c := range cases {
		got, _ := Add(c.x, c.y)
		if got != c.want {
			t.Errorf("Add(%s, %s) == %s, want %s", c.x, c.y, got.(Number), c.want)
		}
	}
}

func TestAddStrings(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Str
	}{
		{Str("Hello"), Str(", World"), Str("Hello, World")},
		{Str("Wooden Roller Coaster "), Number(1), Str("Wooden Roller Coaster 1")},
		{Number(2), Str(" Cool"), Str("2 Cool")},
	}

	for _, c := range cases {
		got, _ := Add(c.x, c.y)
		if got != c.want {
			t.Errorf("Add(%s, %s) == %s, want %s", c.x, c.y, got.(Str), c.want)
		}
	}
}

func TestSubNumbers(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Number
	}{
		{Number(1), Number(0), Number(1)},
		{Number(-1), Number(1), Number(-2)},
		{Number(0.5), Number(0.25), Number(0.25)},
		{Number(1.0), Number(1), Number(0)},
	}

	for _, c := range cases {
		got, _ := Sub(c.x, c.y)
		if got != c.want {
			t.Errorf("Sub(%s, %s) == %s, want %s", c.x, c.y, got.(Number), c.want)
		}
	}
}

func TestMulNumbers(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Number
	}{
		{Number(1), Number(0), Number(0)},
		{Number(-1), Number(1), Number(-1)},
		{Number(0.5), Number(0.25), Number(0.125)},
		{Number(1.0), Number(1), Number(1)},
	}

	for _, c := range cases {
		got, _ := Mul(c.x, c.y)
		if got != c.want {
			t.Errorf("Mul(%s, %s) == %s, want %s", c.x, c.y, got.(Number), c.want)
		}
	}
}

func TestMulStrings(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Str
	}{
		{Str("!"), Number(1), Str("!")},
		{Str("!"), Number(5), Str("!!!!!")},
		{Str("!"), Number(-5), Str("!!!!!")},
		{Str("!"), Number(0), Str("")},
	}

	for _, c := range cases {
		got, _ := Mul(c.x, c.y)
		if got != c.want {
			t.Errorf("Mul(%s, %s) == %s, want %s", c.x, c.y, got.(Str), c.want)
		}
	}
}

func TestDivNumbers(t *testing.T) {
	cases := []struct {
		x, y Algebraic
		want Number
	}{
		{Number(1), Number(1), Number(1)},
		{Number(5), Number(2), Number(2.5)},
		{Number(0.5), Number(0.25), Number(2)},
	}

	for _, c := range cases {
		got, _ := Div(c.x, c.y)
		if got != c.want {
			t.Errorf("Div(%s, %s) == %s, want %s", c.x, c.y, got.(Number), c.want)
		}
	}
}

func TestModNumbers(t *testing.T) {
	cases := []struct {
		x, y Number
		want Number
	}{
		{Number(4), Number(2), Number(0)},
		{Number(4.2), Number(2.23213123), Number(0)},
		{Number(4.9), Number(2.77777), Number(0)},
		{Number(5), Number(2), Number(1)},
	}

	for _, c := range cases {
		got, _ := Mod(c.x, c.y)
		if got != c.want {
			t.Errorf("Mod(%s, %s) == %s, want %s", c.x, c.y, got.(Number), c.want)
		}
	}
}
