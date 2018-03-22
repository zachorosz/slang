package slang

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []struct {
		name string
		expr string
		want []string
	}{
		{"Single Character Tokens", "()[]'", []string{"(", ")", "[", "]", "'"}},
		{"Comment Ignored", ";this is a comment", []string{}},
		{"String", `"My name is Zach."`, []string{`"My name is Zach."`}},
		{"String with escaped quotes", "\"My name is \\\"Zach\\\".\"", []string{"\"My name is \\\"Zach\\\".\""}},
		{
			"Multi-line String",
			`"this is
a multi
line
string"`,
			[]string{"\"this is\na multi\nline\nstring\""}},
		{
			"Symbols and atoms",
			"this-is-A-symbol! 1 123 -123 +123 0.123 true false nil symbolHere",
			[]string{"this-is-A-symbol!", "1", "123", "-123", "+123", "0.123", "true", "false", "nil", "symbolHere"},
		},
		{
			"Program",
			`(define cons [car cdr]
				(lambda [x] (if x car cdr)))`,
			[]string{"(", "define", "cons", "[", "car", "cdr", "]", "(", "lambda", "[", "x", "]", "(", "if", "x", "car", "cdr", ")", ")", ")"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := tokenize(c.expr); !reflect.DeepEqual(got, c.want) {
				t.Errorf("tokenize(%q) = %v, want %v", c.expr, got, c.want)
			}
		})
	}
}
