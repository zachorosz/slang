package parser

import (
	"testing"

	"github.com/zachorosz/slang"
)

var numberTests = []struct {
	input    string
	expected slang.Number
}{
	{"1", slang.Number(1)},
	{"0xD3ADB33F", slang.Number(0xD3ADB33F)},
	{"1e3", slang.Number(1e3)},
	{"+1.2E-4", slang.Number(+1.2E-4)},
	{"0.1234", slang.Number(0.1234)},
}

func TestParseNumber(t *testing.T) {
	for _, test := range numberTests {
		got, err := Parse("numbers", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseNumber", err)
			return
		}
		n, ok := got[0].(slang.Number)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp Number", "TestParseNumber", got[0])
		} else if n != test.expected {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseNumber", got[0], test.expected)
		}
	}
}

var symbolTests = []struct {
	input    string
	expected slang.Symbol
}{
	{"thisIs-a_symbol!", slang.Symbol("thisIs-a_symbol!")},
}

func TestParseSymbol(t *testing.T) {
	for _, test := range symbolTests {
		got, err := Parse("TestParseSymbol", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseSymbol", err)
			return
		}
		sym, ok := got[0].(slang.Symbol)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp Symbol", "TestParseSymbol", got[0])
		} else if sym != test.expected {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseSymbol", got[0], test.expected)
		}
	}
}

var boolTests = []struct {
	input    string
	expected bool
}{
	{"true", true},
	{"false", false},
}

func TestParseBool(t *testing.T) {
	for _, test := range boolTests {
		got, err := Parse("TestParseBool", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseBool", err)
			return
		}
		b, ok := got[0].(bool)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp bool", "TestParseBool", got[0])
		} else if b != test.expected {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseBool", got[0], test.expected)
		}
	}
}

var stringTests = []struct {
	input    string
	expected slang.Str
}{
	{"\"hello world\"", slang.Str("hello world")},
}

func TestParseStr(t *testing.T) {
	for _, test := range stringTests {
		got, err := Parse("TestParseStr", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseStr", err)
			return
		}
		str, ok := got[0].(slang.Str)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp Str", "TestParseStr", got[0])
		} else if str != test.expected {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseStr", got[0], test.expected)
		}
	}
}

var listTests = []struct {
	input    string
	expected slang.List
}{
	{"()", slang.List{}},
	{"(1 2 3 4 5)", slang.MakeList(slang.Number(1), slang.Number(2), slang.Number(3), slang.Number(4), slang.Number(5))},
	{"(1 (2 3))", slang.MakeList(slang.Number(1), slang.MakeList(slang.Number(2), slang.Number(3)))},
}

func TestParseList(t *testing.T) {
	for _, test := range listTests {
		got, err := Parse("TestParseList", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseList", err)
			return
		}
		lst, ok := got[0].(slang.List)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp List", "TestParseList", got[0])
		} else if !slang.Eq(lst, test.expected) {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseList", got[0], test.expected)
		}
	}
}

var vectorTests = []struct {
	input    string
	expected slang.Vector
}{
	{"[]", slang.Vector{}},
	{"[1 2 3 4 5]", slang.MakeVector(slang.Number(1), slang.Number(2), slang.Number(3), slang.Number(4), slang.Number(5))},
	{"[[1 2] [3 4]]", slang.MakeVector(slang.MakeVector(slang.Number(1), slang.Number(2)), slang.MakeVector(slang.Number(3), slang.Number(4)))},
}

func TestParseVector(t *testing.T) {
	for _, test := range vectorTests {
		got, err := Parse("TestParseVector", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseVector", err)
			return
		}
		vec, ok := got[0].(slang.Vector)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp Vector", "TestParseVector", got[0])
		} else if !slang.Eq(vec, test.expected) {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseVector", got[0], test.expected)
		}
	}
}

var quoteTests = []struct {
	input    string
	expected slang.List
}{
	{"'a", slang.MakeList(slang.Symbol("quote"), slang.Symbol("a"))},
	{"'1", slang.MakeList(slang.Symbol("quote"), slang.Number(1))},
	{"'(a b c)", slang.MakeList(slang.Symbol("quote"), slang.MakeList(slang.Symbol("a"), slang.Symbol("b"), slang.Symbol("c")))},
}

func TestParseQuote(t *testing.T) {
	for _, test := range quoteTests {
		got, err := Parse("TestParseQuote", test.input)
		if err != nil {
			t.Errorf("\n%s:\n\tgot unexpected error %+v", "TestParseQuote", err)
			return
		}
		vec, ok := got[0].(slang.List)
		if !ok {
			t.Errorf("\n%s:\n\tgot %T\n\texp List", "TestParseQuote", got[0])
		} else if !slang.Eq(vec, test.expected) {
			t.Errorf("\n%s:\n\tgot %v\n\texp %v", "TestParseQuote", got[0], test.expected)
		}
	}
}
