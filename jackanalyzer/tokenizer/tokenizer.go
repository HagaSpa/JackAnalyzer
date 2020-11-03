package tokenizer

import (
	"bufio"
	"fmt"
	"io"
	"jackanalyzer/token"
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
		c, sz, err := tz.re.ReadRune()
		if err != nil {
			// TODO return err
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("%q [%d]\n", string(c), sz)

		// skip white space
		if unicode.IsSpace(c) {
			continue
		}

		// call startsWithKeyword
		kw := tz.startsWithKeyword()
		if kw != "" {
			cur = newToken(
				cur, token.KEYWORD, kw, "", "", 0, "",
			)
			tz.re.Discard(len(kw))
			continue
		}

	}
	return &head
}

func (tz *Tokenizer) startsWithKeyword() token.Keyword {
	tz.re.UnreadRune()
	for k, v := range token.Keywords {
		l := len(k)
		d, err := tz.re.Peek(l)
		if err == io.EOF {
			// TODO return err
		}
		if k == string(d) {
			return v
		}
	}
	return "" // TODO: Should token.Keyword cotain an empty string??
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
