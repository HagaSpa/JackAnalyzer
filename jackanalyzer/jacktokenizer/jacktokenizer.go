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
	s *bufio.Scanner
}

func New(r io.Reader) *JackTokenizer {
	s := bufio.NewScanner(r)
	jt := &JackTokenizer{
		s: s,
	}
	return jt
}

func (jt *JackTokenizer) Tokenize(r io.Reader) *Token {
	// todo
	return nil
}
