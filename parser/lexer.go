package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// enumerated constants for token types
const (
	tokenEOF          tokenType = iota
	tokenLeftParen              // left paren (, open list
	tokenRightParen             // right paren ), closing list
	tokenLeftBracket            // left bracket [, open vector
	tokenRightBracket           // right bracket ], closing vector
	tokenQuote                  // quote '
	tokenNumber                 // number
	tokenString                 // string
	tokenSymbol                 // symbol
)

const eof = -1

// Recursive definition of a state function. Executes actions and returns the next state function.
type stateFn func(*lexer) stateFn

type tokenType int

type token struct {
	typ     tokenType
	literal string
	pos     int // starting position within input string
	line    int // line number
}

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	}
	return fmt.Sprintf("%q", t.literal)
}

type lexer struct {
	name   string     // arbitrary name used for debugging and/or error reporting
	input  string     // string being scanned
	start  int        // start position of a token within input string
	pos    int        // current position in the input string
	width  int        // width (size) of last rune read from input
	line   int        // line number (number of newlines seen)
	tokens chan token // channel with scanned tokens
}

// ignore skips over pending input.
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// NOTE(zachorosz): Can only be called once per call of next(); lexer width state is only the size
// of last rune.
func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, size := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = size
	l.pos += l.width

	return r
}

// peek returns, but does not eat, the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept eats next rune iff it is from the string of valid runes.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun eats a run of valid runes until it encounters an invalid rune.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// emit sends the token currently being anaylzed through the token channel. The state calling emit
// specifies the type.
func (l *lexer) emit(t tokenType) {
	current := l.input[l.start:l.pos]
	l.tokens <- token{t, current, l.pos, l.line}
	l.start = l.pos
}

// run scans and lexes input by executing state functions until the next state is nil.
// The initial state is lexText.
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	// done emitting tokens
	close(l.tokens)
}

func (l *lexer) nextToken() token {
	return <-l.tokens
}

// lex creates a new scanner for the input string.
func lex(name, input string) *lexer {
	l := &lexer{
		name:   name,
		input:  input,
		line:   1,
		tokens: make(chan token),
	}
	go l.run()
	return l
}

func lexLeftParen(l *lexer) stateFn {
	l.emit(tokenLeftParen)
	return lexText
}

func lexRightParen(l *lexer) stateFn {
	l.emit(tokenRightParen)
	return lexText
}

func lexLeftBracket(l *lexer) stateFn {
	l.emit(tokenLeftBracket)
	return lexText
}

func lexRightBracket(l *lexer) stateFn {
	l.emit(tokenRightBracket)
	return lexText
}

func lexQuote(l *lexer) stateFn {
	l.emit(tokenQuote)
	return lexText
}

// lexString accepts a run of characters between two double-quotes. The first quote is assumed to be
// seen already. Multi-line, or raw, strings are allowed
func lexString(l *lexer) stateFn {
	l.ignore() // ignore leading '"'
Loop:
	for {
		switch r := l.next(); {
		case r == '\\':
			// get escaped rune, make sure it is not a EOF or newline
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case r == eof:
			panic("unterminated string")
		case isLineEnding(r):
			l.line++
		case r == '"':
			// backup to not emit string with ending '"'
			l.backup()
			break Loop
		}
	}
	l.emit(tokenString)
	// call next to get to ending '"' again and ignore it
	l.next()
	l.ignore()
	return lexText
}

// lexComment lexes a comment beginning with ';'. Tokens from this position to the next newline are
// ultimately ignored. Comment delimiter ';' is assumed to be seen already.
func lexComment(l *lexer) stateFn {
	comment := l.input[l.pos:]
	if index := strings.Index(comment, "\n"); index > -1 {
		l.pos += index // ';' is at pos, offset to next newline
	} else {
		// if no newline, skip to EOF
		l.pos = len(comment) + 1
	}
	l.ignore() // ignore comment
	return lexText
}

func lexWhiteSpace(l *lexer) stateFn {
	for isWhiteSpace(l.peek()) {
		l.next()
	}
	l.ignore()
	return lexText
}

func lexText(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(tokenEOF)
			return nil // stop state machine
		case isLineEnding(r):
			l.ignore()
			l.line++
			return lexText
		case isWhiteSpace(r):
			return lexWhiteSpace
		case r == '(':
			return lexLeftParen
		case r == ')':
			return lexRightParen
		case r == '[':
			return lexLeftBracket
		case r == ']':
			return lexRightBracket
		case r == '\'':
			return lexQuote
		case r == '"':
			return lexString
		case r == ';':
			return lexComment
		default:
			panic(fmt.Sprintf("unknown rune %q", r))
		}
	}
}

func isLineEnding(r rune) bool {
	return r == '\r' || r == '\n'
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
