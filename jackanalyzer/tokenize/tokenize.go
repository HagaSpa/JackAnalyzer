package tokenize

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type TokenType int
type Keyword string

type Token struct {
	next       *Token
	tokenType  TokenType
	keyword    Keyword
	symbol     string
	identifier string
	intVal     int
	stringVal  string
}

type JackTokenizer struct {
	re *bufio.Reader
}

const (
	_ TokenType = iota
	KEYWORD
	SYMBOL
	IDENTIFIER
	INT_CONST
	STRING_CONST
)

// keywords
const (
	CLASS       = "class"
	METHOD      = "method"
	FUNCTION    = "function"
	CONSTRUCTOR = "constructor"
	INT         = "int"
	BOOLEAN     = "boolean"
	CHAR        = "char"
	VOID        = "void"
	VAR         = "var"
	STATIC      = "static"
	FIELD       = "field"
	LET         = "let"
	DO          = "do"
	IF          = "if"
	ELSE        = "else"
	WHILE       = "while"
	RETURN      = "return"
	TRUE        = "true"
	FALSE       = "false"
	NULL        = "null"
	THIS        = "this"
)

var keywords = map[string]Keyword{
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

func (jt *JackTokenizer) startsWithKeyword() Keyword {
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
	return "" // TODO: Should I return an empty string?
}
