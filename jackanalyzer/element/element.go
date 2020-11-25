package element

import (
	"encoding/xml"
)

type Class struct {
	Modifier      string          // 'class'
	ClassName     ClassName       // identifier
	LBrace        string          // '{'
	ClassVarDec   []ClassVarDec   // classVarDec*
	SubtoutineDec []SubroutineDec // subroutineDec*
	RBrace        string          // '}'
}

type ClassVarDec struct {
	Modifier  string    // 'static' | 'field'
	VarType   Types     // 'int' | 'char' | 'boolean' | className
	VarNames  []VarName // varName (, varName)*
	SemiColon string    // ';'
}

type SubroutineDec struct {
	Modifier      string         // 'constructor' | 'function' | 'method'
	SubType       string         // 'void' | type
	SubName       SubroutineName // identifier
	LParan        string         // '('
	ParameterList []Parameter    // (type varName (, type, varName)*)?
	RParen        string         // ')'
	SubBody       SubroutineBody // subroutineBody
}

type SubroutineBody struct {
	LBrace     string // '{'
	VarDec     VarDec // varDec*
	Statements []Statement
}

type VarDec struct {
	Modifier  string    // 'var'
	VarType   Types     // type
	VarNames  []VarName // varName (, varName)*
	SemiColon string    // ';'
}

type Parameter struct {
	paramType Types
	parmaName VarName
	comma     string
}
type Types string          // 'int' | 'char' | 'boolean' | className
type ClassName string      // identifier
type SubroutineName string // identifier
type VarName string        // identifier

/* Statements */
type Statement interface {
	statement()
}

/* Expession */
type Expression struct {
	Term Term
	Next []Ops
}
type Ops struct {
	Bop  Op // binary operator
	Term Term
}
type Term string

/*
FIXME: have no idea...

subroutineName '(' expressionList ')' |
(className | varName) '.' subroutineName '(' expressionList ')'
*/
type SubroutineCall struct {
	SubName        string
	LParan         string       // '('
	ExpressionList []Expression // (expression(, expression)*)?
	RParen         string       // ')'
	//  ClassName and VarName should be common inteface??
	Name string // ClassName | VarName
}

type KeywordConstant string
type UnaryOp string
type Op string

func (cl Class) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// class
	start.Name.Local = "class"
	e.EncodeToken(start)
	e.EncodeElement(genContent(cl.Modifier), genTagKeyword())
	e.EncodeElement(genContent(cl.ClassName), genTagIdentifier())
	e.EncodeElement(genContent(cl.LBrace), genTagSymbol())

	// ClassVarDec
	if len(cl.ClassVarDec) != 0 {
		for _, v := range cl.ClassVarDec {
			v.genClassVarDec(e)
		}
	}

	e.EncodeToken(start.End())
	return nil
}

func (cd ClassVarDec) genClassVarDec(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "classVarDec"}}
	e.EncodeToken(start)
	e.EncodeElement(genContent(cd.Modifier), genTagKeyword())
	e.EncodeElement(genContent(cd.VarType), genTagKeyword())

	for i, v := range cd.VarNames {
		e.EncodeElement(genContent(v), genTagIdentifier())
		if i < len(cd.VarNames)-1 {
			e.EncodeElement(genContent(","), genTagSymbol())
		}
	}

	e.EncodeElement(genContent(cd.SemiColon), genTagSymbol())
	e.EncodeToken(start.End())
}

func genContent(s interface{}) string {
	switch s.(type) {
	case string:
		str, _ := s.(string)
		return " " + str + " "
	case ClassName:
		str, _ := s.(ClassName)
		return " " + string(str) + " "
	case Types:
		str, _ := s.(Types)
		return " " + string(str) + " "
	case VarName:
		str, _ := s.(VarName)
		return " " + string(str) + " "
	default:
		return "" // invalid types
	}
}

func genTagKeyword() xml.StartElement {
	return xml.StartElement{Name: xml.Name{Local: "keyword"}}
}

func genTagIdentifier() xml.StartElement {
	return xml.StartElement{Name: xml.Name{Local: "identifier"}}
}

func genTagSymbol() xml.StartElement {
	return xml.StartElement{Name: xml.Name{Local: "symbol"}}
}
