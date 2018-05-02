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
	{"parens", "()", []token{leftParenToken, rightParenToken, eofToken}},
	{"brackets", "[]", []token{leftBracketToken, rightBracketToken, eofToken}},
	{"quote", "'()", []token{quoteToken, leftParenToken, rightParenToken, eofToken}},
	{"string", "\"hello world\"", []token{
		token{typ: tokenString, literal: "hello world"},
		eofToken,
	}},
	{"empty string", "\"\"", []token{token{typ: tokenString, literal: ""}, eofToken}},
	{"string with escaped chars", "\"hello\\n\\\"world\\\"\"", []token{
		token{typ: tokenString, literal: "hello\\n\\\"world\\\""},
		eofToken,
	}},
	{"comment", "; this is a comment!", []token{eofToken}},
	{"comment delimited by newline", "; an empty list\n'()", []token{
		quoteToken,
		leftParenToken,
		rightParenToken,
		eofToken,
	}},
	{"whitespace", " \t", []token{eofToken}},
	{"symbols", "! $ % & * _ + - = < > ? / thisIs-a_symbol!", []token{
		token{typ: tokenSymbol, literal: "!"},
		token{typ: tokenSymbol, literal: "$"},
		token{typ: tokenSymbol, literal: "%"},
		token{typ: tokenSymbol, literal: "&"},
		token{typ: tokenSymbol, literal: "*"},
		token{typ: tokenSymbol, literal: "_"},
		token{typ: tokenSymbol, literal: "+"},
		token{typ: tokenSymbol, literal: "-"},
		token{typ: tokenSymbol, literal: "="},
		token{typ: tokenSymbol, literal: "<"},
		token{typ: tokenSymbol, literal: ">"},
		token{typ: tokenSymbol, literal: "?"},
		token{typ: tokenSymbol, literal: "/"},
		token{typ: tokenSymbol, literal: "thisIs-a_symbol!"},
		eofToken,
	}},
	{"numbers", "1 0xd3ADb33f -1.2i 1.2i 1e3 +1.2e-4 0.1234 1+2i", []token{
		token{typ: tokenNumber, literal: "1"},
		token{typ: tokenNumber, literal: "0xd3ADb33f"},
		token{typ: tokenNumber, literal: "-1.2i"},
		token{typ: tokenNumber, literal: "1.2i"},
		token{typ: tokenNumber, literal: "1e3"},
		token{typ: tokenNumber, literal: "+1.2e-4"},
		token{typ: tokenNumber, literal: "0.1234"},
		token{typ: tokenComplexNumber, literal: "1+2i"},
		eofToken,
	}},
}

func equal(t1, t2 []token) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := range t1 {
		if t1[i].typ != t2[i].typ {
			return false
		}
		if t1[i].literal != t2[i].literal {
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
			t.Errorf("\n%s:\n\tgot %+v\n\texp %v", test.name, got, test.tokens)
		}
	}
}
