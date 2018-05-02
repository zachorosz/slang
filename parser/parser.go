package parser

import (
	"fmt"
	"strconv"

	"github.com/zachorosz/slang"
)

type ParseError struct {
	name    string
	tok     token
	message string
}

func (err ParseError) Error() string {
	return fmt.Sprintf("%s: %s (%d, %d)", err.name, err.message, err.tok.line, err.tok.pos)
}

type parser struct {
	lexer   *lexer
	current *token
}

func (p *parser) next() *token {
	tok := p.lexer.nextToken()
	p.current = &tok
	return p.current
}

func (p *parser) peek() *token {
	if p.current == nil {
		return p.next()
	}
	return p.current
}

func parseNumber(p *parser) (slang.LangType, error) {
	tok := p.peek()
	if n, err := strconv.ParseInt(tok.literal, 0, 64); err == nil {
		return slang.Number(float64(n)), nil
	}
	n, err := strconv.ParseFloat(tok.literal, 64)
	if err != nil {
		return nil, err
	}
	return slang.Number(n), nil
}

func parseQuote(p *parser) (slang.List, error) {
	p.next() // throw away quote
	quoted, err := parse(p)
	if err != nil {
		return slang.List{}, err
	}
	return slang.MakeList(slang.Symbol("quote"), quoted), nil
}

func parseSequence(p *parser, close tokenType, seq slang.Sequence) (slang.Sequence, error) {
	for tok := p.next(); tok.typ != close; tok = p.next() {
		form, err := parse(p)
		if err != nil {
			return nil, err
		}
		seq = seq.Append(form)
	}
	return seq, nil
}

func parseSymbol(p *parser) slang.LangType {
	tok := p.peek()
	switch tok.literal {
	case "true":
		return true
	case "false":
		return false
	case "nil":
		return nil
	default:
		return slang.Symbol(tok.literal)
	}
}

func parse(p *parser) (slang.LangType, error) {
	tok := p.peek()
	switch tok.typ {
	case tokenEOF:
		return nil, ParseError{p.lexer.name, *tok, "Unexpected EOF"}
	case tokenError:
		return nil, ParseError{p.lexer.name, *tok, tok.literal}
	case tokenLeftParen:
		return parseSequence(p, tokenRightParen, slang.List{})
	case tokenRightParen:
		return nil, ParseError{p.lexer.name, *tok, fmt.Sprintf("Unexpected '%s'", tok.literal)}
	case tokenLeftBracket:
		return parseSequence(p, tokenRightBracket, slang.Vector{})
	case tokenRightBracket:
		return nil, ParseError{p.lexer.name, *tok, fmt.Sprintf("Unexpected '%s'", tok.literal)}
	case tokenQuote:
		return parseQuote(p)
	case tokenNumber, tokenComplexNumber:
		return parseNumber(p)
	case tokenString:
		return slang.Str(tok.literal), nil
	case tokenSymbol:
		return parseSymbol(p), nil
	default:
		return nil, ParseError{p.lexer.name, *tok, fmt.Sprintf("Encountered token %s with unknown type", tok.literal)}
	}
}

func newParser(l *lexer) *parser {
	return &parser{
		lexer: l,
	}
}

// Parse lexes and parses a slang input string
func Parse(name, input string) ([]slang.LangType, error) {
	p := newParser(lex(name, input))
	program := make([]slang.LangType, 0)
	for p.next().typ != tokenEOF {
		form, err := parse(p)
		if err != nil {
			return nil, err
		}
		program = append(program, form)
	}
	return program, nil
}
