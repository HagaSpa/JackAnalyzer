package cmplengn

import (
	"encoding/xml"
	"jackanalyzer/token"
	"strconv"

	"golang.org/x/xerrors"
)

type CompilationEngine struct {
	t token.Token
	e *xml.Encoder
}

func New(t token.Token, e *xml.Encoder) *CompilationEngine {
	ce := &CompilationEngine{
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

// Compile Expression.
//
//  term (op term)*
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

// Compile Term.
//
//  integerConstant | stringConstant | keywordConstant | varName | varName '[' expression ']' | subroutineCall | '(' expression ')' | unaryOp term
//
//  subroutineCall: subroutineName '(' expressionList ')' | (className | varName) '.' subroutineName '(' expressionList ')'
//  unaryOp: '-' | '~'
func (ce *CompilationEngine) compileTerm() error {
	start := xml.StartElement{Name: xml.Name{Local: "term"}}
	ce.e.EncodeToken(start)

	switch ce.t.TokenType {
	case token.INT_CONST, token.STRING_CONST, token.KEYWORD:
		// integerConstant | stringConstant | keywordConstant
	}
	ce.e.EncodeToken(start.End())
	return nil
}

func (ce *CompilationEngine) compileExpressionList() {}

func (ce *CompilationEngine) writeToken() {
	if ce.t.HasMoreTokens() {
		ce.e.EncodeElement(genElement(ce.t))
	}
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
