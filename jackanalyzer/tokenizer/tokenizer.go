package tokenizer

import (
	"bufio"
	"io"
	"jackanalyzer/token"
	"strconv"
	"unicode"
)

type Tokenizer struct {
	re *bufio.Reader
}

func New(r io.Reader) *Tokenizer {
	re := bufio.NewReader(r)
	tz := &Tokenizer{
		re: re,
	}
	return tz
}

func (tz *Tokenizer) Tokenize() *token.Token {
	head := token.Token{
		Next: nil,
	}
	cur := &head

	// tokenize until EOF comes out
	for {
		// call startsWithKeyword
		kw := tz.startsWithKeyword()
		if kw != "" {
			cur = newToken(
				cur, token.KEYWORD, kw, "", "", 0, "",
			)
			continue
		}

		c, _, err := tz.re.ReadRune()
		if err != nil {
			// TODO return err
		}
		if err == io.EOF {
			break
		}

		// skip white space
		if unicode.IsSpace(c) {
			continue
		}

		// isComment
		if ok, ct := tz.isComment(c); ok {
			tz.skipComment(ct)
			continue
		}

		// IsSymbol?
		// TODO: if unicode.IsPunct() == true
		if token.IsSymbol(c) {
			cur = newToken(
				cur, token.SYMBOL, "", string(c), "", 0, "",
			)
			continue
		}

		// Identifier
		if isAlpherUnder(c) {
			id := tz.startsWithIdentifier(c)
			cur = newToken(
				cur, token.IDENTIFIER, "", "", id, 0, "",
			)
			continue
		}

		// IntegerConstant
		if unicode.IsNumber(c) {
			iv := tz.startsWithIntegerConstant(c)
			cur = newToken(
				cur, token.INT_CONST, "", "", "", iv, "",
			)
			continue
		}

		// StringConstant
		if isDoubleQuotes(c) {
			sv := tz.startsWithStringConstant()
			cur = newToken(
				cur, token.STRING_CONST, "", "", "", 0, sv,
			)
			continue
		}
	}
	return &head
}

func (tz *Tokenizer) startsWithKeyword() token.Keyword {
	for k, v := range token.Keywords {
		l := len(k)
		d, err := tz.re.Peek(l)
		if err == io.EOF {
			// TODO return err
		}
		if k == string(d) {
			tz.re.Discard(l)
			return v
		}
	}
	return "" // TODO: Should token.Keyword cotain an empty string??
}

func (tz *Tokenizer) startsWithIdentifier(r rune) string {
	id := string(r)
	for {
		c, _, err := tz.re.ReadRune()
		if err == io.EOF {
			break
		}
		if isAlpherUnder(c) || unicode.IsNumber(c) {
			id = id + string(c)
			continue
		}
		tz.re.UnreadRune()
		break
	}
	return id
}

func (tz *Tokenizer) startsWithIntegerConstant(r rune) int {
	sr := string(r)
	for {
		c, _, err := tz.re.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsNumber(c) {
			sr = sr + string(c)
			continue
		}
		tz.re.UnreadRune()
		break
	}
	iv, err := strconv.Atoi(sr)
	if err != nil {
		// TODO return err
	}
	return iv
}

func (tz *Tokenizer) startsWithStringConstant() string {
	var sv string
	for {
		c, _, err := tz.re.ReadRune()
		if err == io.EOF {
			break
		}
		if isDoubleQuotes(c) {
			break
		}
		sv = sv + string(c)
	}
	return sv
}

func newToken(
	cur *token.Token,
	tt token.TokenType,
	kw token.Keyword,
	sb string,
	id string,
	iv int,
	sv string,
) *token.Token {
	nt := token.Token{
		TokenType:  tt,
		Keyword:    kw,
		Symbol:     sb,
		Identifier: id,
		IntVal:     iv,
		StringVal:  sv,
	}
	cur.Next = &nt
	return &nt
}

func isAlpherUnder(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || (r == '_')
}

func isDoubleQuotes(r rune) bool {
	return r == '"'
}

func (tz *Tokenizer) isComment(r rune) (bool, string) {
	if r != '/' {
		return false, ""
	}
	c, _, err := tz.re.ReadRune()
	if err == io.EOF {
		return false, ""
	}
	if c == '/' {
		return true, token.COMMENT
	}
	if c == '*' {
		return true, token.COMMENT_AST
	}
	tz.re.UnreadRune()
	return false, ""
}

func (tz *Tokenizer) skipComment(ct string) {
L:
	for {
		c, _, err := tz.re.ReadRune()
		if err == io.EOF {
			break
		}
		switch ct {
		case token.COMMENT:
			if c == '\n' {
				break L
			}
		case token.COMMENT_AST:
			if c == '*' {
				c2, _, err := tz.re.ReadRune()
				if err == io.EOF {
					break L
				}
				if c2 == '/' {
					break L
				}
				tz.re.UnreadRune()
			}
		}
	}
}
