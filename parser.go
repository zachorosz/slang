package slang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type reader interface {
	peek() *string
	next() *string
}

type tokenReader struct {
	tokens   []string
	position int
}

func (rdr *tokenReader) peek() *string {
	if rdr.position >= len(rdr.tokens) {
		return nil
	}
	return &rdr.tokens[rdr.position]
}

func (rdr *tokenReader) next() *string {
	if rdr.position >= len(rdr.tokens) {
		return nil
	}
	token := rdr.tokens[rdr.position]
	rdr.position++
	return &token
}

func tokenize(expr string) []string {
	r := regexp.MustCompile(`[()\[\]']|;.*|"(?:\\.|[^\\"])*"|[^\s()']*`)
	tokens := make([]string, 0)
	for _, token := range r.FindAllString(strings.TrimSpace(expr), -1) {
		if !strings.HasPrefix(token, ";") {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func parseAtom(rdr reader) (LangType, error) {
	token := rdr.next()

	if nAtom, err := strconv.ParseFloat(*token, 64); err == nil {
		return Number(nAtom), nil
	}

	if strings.HasPrefix(*token, "\"") {
		if strings.HasSuffix(*token, "\"") {
			return (*token)[1:len(*token)], nil
		}
		return nil, fmt.Errorf(`Expected string to end with non-escaped double quotes '"'`)
	}

	if *token == "true" {
		return true, nil
	}

	if *token == "false" {
		return false, nil
	}

	if *token == "nil" {
		return nil, nil
	}

	return Symbol(*token), nil
}

func parseSequence(rdr reader, open string, close string, seq Sequence) (Sequence, error) {
	token := rdr.next()
	if token == nil || *token != open {
		return nil, fmt.Errorf("Expected sequence to begin with '%s'", open)
	}

	for token = rdr.peek(); true; token = rdr.peek() {
		if token == nil {
			return nil, fmt.Errorf("Unbalanced '%s%s'", open, close)
		}
		if *token == close {
			break
		}
		form, err := parseForm(rdr)
		if err != nil {
			return nil, err
		}
		seq = seq.Append(form)
	}
	// skip over closing delim
	rdr.next()

	return seq, nil
}

func parseQuote(rdr reader) (List, error) {
	rdr.next()

	quoted, err := parseForm(rdr)
	if err != nil {
		return List{}, err
	}
	//quoteForm := List{Symbol("quote"), quoted}
	quoteForm := MakeList(Symbol("quote"), quoted)
	return quoteForm, nil
}

func parseForm(rdr reader) (LangType, error) {
	token := rdr.peek()
	if token == nil {
		return nil, nil
	}

	switch *token {
	case "(":
		lst, err := parseSequence(rdr, "(", ")", List{})
		return lst, err
	case ")":
		return nil, fmt.Errorf("Unexpected ')'")
	case "[":
		vec, err := parseSequence(rdr, "[", "]", Vector{})
		return vec, err
	case "]":
		return nil, fmt.Errorf("Unexpected ']'")
	case "'":
		return parseQuote(rdr)
	default:
		return parseAtom(rdr)
	}
}

// Read tokenizes and parses an expression string into a domain language type.
func Read(expr string) (LangType, error) {
	tokens := tokenize(expr)
	AST, err := parseForm(&tokenReader{tokens, 0})
	if err != nil {
		return nil, err
	}
	return AST, nil
}
