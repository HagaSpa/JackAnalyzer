package cmplengn

import (
	"encoding/xml"
	"io"
	"jackanalyzer/token"
	"strconv"
)

type CompilationEngine struct {
	w io.Writer
	t token.Token
	e *xml.Encoder
}

func New(w io.Writer, t token.Token) *CompilationEngine {
	e := xml.NewEncoder(w)
	ce := &CompilationEngine{
		w: w,
		t: t,
		e: e,
	}
	return ce
}

// generate Element for *xml.EncodeElement.
func genElement(t token.Token) (string, xml.StartElement) {
	var c string // contents
	var l string // labels
	switch t.TokenType {
	case token.KEYWORD:
		c = string(t.Keyword)
		l = "keyword"
	case token.IDENTIFIER:
		c = t.Identifier
		l = "identifier"
	case token.SYMBOL:
		c = t.Symbol
		l = "symbol"
	case token.INT_CONST:
		c = strconv.Itoa(int(t.IntVal))
		l = "integerConstant"
	case token.STRING_CONST:
		c = t.StringVal
		l = "stringConstant"
	}
	return " " + c + " ", xml.StartElement{Name: xml.Name{Local: l}}
}
