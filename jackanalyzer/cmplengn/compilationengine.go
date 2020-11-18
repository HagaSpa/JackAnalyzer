package cmplengn

import (
	"io"
	"jackanalyzer/token"
)

type CompilationEngine struct {
	w io.Writer
	t token.Token
}

// expression

// TODO: いやわっかんねー
// type expression struct {
// 	term term `xml:"term"`
// 	op op
// }
// type expressionList struct {
// 	expressions []expression `xml:"expression"`
// }

// Parameter is element for parameterList
type parameter struct {
	Keyword    string
	Identifier string
	Symbol     string
}
type op string              // '+' | '-' | '*' | '&' | '|' | '<' | '>' | '='
type unaryOp string         // '+' | '-'
type keywordConstant string // 'true' | 'false' | 'null' | 'this'

func New(w io.Writer, t token.Token) *CompilationEngine {
	ce := &CompilationEngine{
		w: w,
		t: t,
	}
	return ce
}
