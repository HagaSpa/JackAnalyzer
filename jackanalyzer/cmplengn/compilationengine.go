package cmplengn

import (
	"encoding/xml"
	"io"
	"jackanalyzer/token"
	"strconv"

	"golang.org/x/xerrors"
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

func (ce *CompilationEngine) Compile() {}

func (ce *CompilationEngine) compileClass() {}

func (ce *CompilationEngine) compileClassVarDec() {}

func (ce *CompilationEngine) compileSubroutine() {}

func (ce *CompilationEngine) compileParameterList() {}

func (ce *CompilationEngine) compileVarDec() {}

func (ce *CompilationEngine) compileStatements() {}

func (ce *CompilationEngine) compileDo() {}

func (ce *CompilationEngine) compileLet() {}

func (ce *CompilationEngine) compileWhile() {}

func (ce *CompilationEngine) compileReturn() {}

func (ce *CompilationEngine) compileIf() {}

// term (op term)*
func (ce *CompilationEngine) compileExpression() error {
	start := xml.StartElement{Name: xml.Name{Local: "expression"}}
	ce.e.EncodeToken(start)
	ce.compileTerm()
	for ce.t.IsOp() {
		ce.e.EncodeElement(genElement(ce.t))
		if !ce.t.HasMoreTokens() {
			return xerrors.New("invalid syntax. compileExpression")
		}
		ce.t.Advance()
		ce.compileTerm()
	}
	ce.e.EncodeToken(start.End())
	return nil
}

func (ce *CompilationEngine) compileTerm() {}

func (ce *CompilationEngine) compileExpressionList() {}

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
