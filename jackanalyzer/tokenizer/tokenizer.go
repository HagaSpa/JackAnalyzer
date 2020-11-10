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

		// IsSymbol?
		// TODO: if unicode.IsPunct() == true
		if token.IsSymbol(c) {
			cur = newToken(
				cur, token.SYMBOL, "", string(c), "", 0, "",
			)
			continue
		}

		// IsIdentifier?
		if unicode.IsLetter(c) {
			id := tz.startsWithIdentifier(c)
			cur = newToken(
				cur, token.IDENTIFIER, "", "", id, 0, "",
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
		// [A-Z] or [a-z] or _ or [0-9]
		if ('a' <= c && c <= 'z') ||
			('A' <= c && c <= 'Z') ||
			(c == '_') ||
			unicode.IsNumber(c) {
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
