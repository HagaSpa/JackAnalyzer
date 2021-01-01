package token

type TokenType int
type Keyword string

type Token struct {
	Next       *Token
	TokenType  TokenType
	Keyword    Keyword
	Symbol     string
	Identifier string
	IntVal     int
	StringVal  string
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

// comment type
const (
	COMMENT     = "//"
	COMMENT_AST = "/*"
)

var Keywords = map[string]Keyword{
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

var symbols = []rune{
	'{',
	'}',
	'(',
	')',
	'[',
	']',
	'.',
	',',
	';',
	'+',
	'-',
	'*',
	'/',
	'&',
	'|',
	'<',
	'>',
	'=',
	'~',
}

var ops = []string{
	"+",
	"-",
	"*",
	"/",
	"&",
	"|",
	"<",
	">",
	"=",
}

func IsSymbol(r rune) bool {
	for _, v := range symbols {
		if r == v {
			return true
		}
	}
	return false
}

func (t *Token) IsOp() bool {
	if t.TokenType == SYMBOL {
		for _, v := range ops {
			if t.Symbol == v {
				return true
			}
		}
	}
	return false
}

func (t *Token) HasMoreTokens() bool {
	return t.Next != nil
}

func (t *Token) Advance() {
	nxt := t.Next
	t.Next = nxt.Next
	t.TokenType = nxt.TokenType
	t.Keyword = nxt.Keyword
	t.Symbol = nxt.Symbol
	t.Identifier = nxt.Identifier
	t.IntVal = nxt.IntVal
	t.StringVal = nxt.StringVal
}
