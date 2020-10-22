package jacktokenizer

import (
	"bufio"
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
	// todo
	head := Token{
		next: nil,
	}
	return &head
}
