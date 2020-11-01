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
	None KeyWord = iota // not keyword
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

var keywords = map[string]KeyWord{
	"class":       CLASS,
	"method":      METHOD,
	"function":    FUNCTION,
	"constructor": CONSTRUCTOR,
	"int":         INT,
	"boolean":     BOOLEAN,
	"char":        CHAR,
	"void":        VOID,
	"var":         VAR,
	"static":      STATIC,
	"field":       FIELD,
	"let":         LET,
	"do":          DO,
	"if":          IF,
	"else":        ELSE,
	"while":       WHILE,
	"return":      RETURN,
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"this":        THIS,
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

		// skip white space
		if unicode.IsSpace(c) {
			continue
		}
	}
	return &head
}

func (jt *JackTokenizer) startsWithKeyword() KeyWord {
	for k, v := range keywords {
		l := len(k)
		d, err := jt.re.Peek(l)
		if err == io.EOF {
			// TODO return err
		}
		if k == string(d) {
			return v
		}
	}
	return None
}
