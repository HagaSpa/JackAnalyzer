package jacktokenizer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Type int
type KeyWord int

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

const (
	_ Type = iota
	KEYWORD
	SYMBOL
	IDENTIFIER
	INT_CONST
	STRING_CONST
)

const (
	_ KeyWord = iota
	CLASS
	METHOD
	FUNCTION
	CONSTRUCTOR
	INT
	BOOLEAN
	CHAR
	VOID
	VAR
	STATIC
	FIELD
	LET
	DO
	IF
	ELSE
	WHILE
	RETURN
	TRUE
	FALSE
	NULL
	THIS
)

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

		// skip white space
		if unicode.IsSpace(c) {
			continue
		}
	}
	return &head
}
