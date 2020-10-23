package jacktokenizer

import (
	"bufio"
	"fmt"
	"io"
)

type Type int
type KeyWord string

type Token struct {
	next       *Token
	tokenType  Type
	keyword    KeyWord
	symbol     string
	identifier string
	intVal     int
	stringVal  string
}

type JackTokenizer struct {
	re *bufio.Reader
}

func New(r io.Reader) *JackTokenizer {
	re := bufio.NewReader(r)
	jt := &JackTokenizer{
		re: re,
	}
	return jt
}

func (jt *JackTokenizer) Tokenize() *Token {
	head := Token{
		next: nil,
	}

	// tokenize until EOF comes out
	for {
		c, sz, err := jt.re.ReadRune()
		if err != nil {
			// TODO return err
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("%q [%d]\n", string(c), sz)
		// TODO call tokenize method
	}
	return &head
}
