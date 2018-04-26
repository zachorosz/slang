package parser

import "testing"

var (
	eofToken          = token{typ: tokenEOF, literal: ""}
	leftParenToken    = token{typ: tokenLeftParen, literal: "("}
	rightParenToken   = token{typ: tokenRightParen, literal: ")"}
	leftBracketToken  = token{typ: tokenLeftBracket, literal: "["}
	rightBracketToken = token{typ: tokenRightBracket, literal: "]"}
	quoteToken        = token{typ: tokenQuote, literal: "'"}
)

var lexTests = []struct {
	name   string
	input  string
	tokens []token
}{
	{"text with left paren", "(", []token{leftParenToken, eofToken}},
	{"text with right paren", ")", []token{rightParenToken, eofToken}},
	{"text with left and right paren", "()", []token{leftParenToken, rightParenToken, eofToken}},
	{"text with left bracket", "[", []token{leftBracketToken, eofToken}},
	{"text with right bracket", "]", []token{rightBracketToken, eofToken}},
	{"text with quote", "'", []token{quoteToken, eofToken}},
	{"text with string", "\"hello world\"", []token{token{typ: tokenString, literal: "hello world"}, eofToken}},
	{"text with empty string", "\"\"", []token{token{typ: tokenString, literal: ""}, eofToken}},
	{"text with string with escaped chars", "\"hello\\n\\\"world\\\"\"", []token{token{typ: tokenString, literal: "hello\\n\\\"world\\\""}, eofToken}},
	{"text with comment", "; this is a comment!", []token{eofToken}},
	{"text with comment delimited by newline", "; an empty list\n'()", []token{quoteToken, leftParenToken, rightParenToken, eofToken}},
	{"text with all whitespace", "    \t   \t\t", []token{eofToken}},
}

func equal(expected, got []token) bool {
	if len(expected) != len(got) {
		return false
	}
	for i := range expected {
		if expected[i].typ != got[i].typ {
			return false
		}
		if expected[i].literal != got[i].literal {
			return false
		}
	}
	return true
}

func collect(l *lexer) []token {
	tokens := make([]token, 0)
	for {
		tok := l.nextToken()
		tokens = append(tokens, tok)
		if tok.typ == tokenEOF {
			break
		}
	}
	return tokens
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		l := lex(test.name, test.input)
		got := collect(l)
		if !equal(test.tokens, got) {
			t.Errorf("%s: got %+v expected %v", test.name, got, test.tokens)
		}
	}
}
