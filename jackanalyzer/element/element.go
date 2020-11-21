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

type Statement struct{}

type Parameter struct {
	paramType Types
	parmaName VarName
	comma     string
}
type Types string          // 'int' | 'char' | 'boolean' | className
type ClassName string      // identifier
type SubroutineName string // identifier
type VarName string        // identifier

func (cl Class) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// class
	start.Name.Local = "class"
	e.EncodeToken(start)
	e.EncodeElement(genContent(cl.Modifier), genTagKeyword())
	e.EncodeElement(genContent(cl.ClassName), genTagIdentifier())
	e.EncodeElement(genContent(cl.LBrace), genTagSymbol())

	// TODO: if cl.ClassVarDec != nil: call (cvd ClassVarDec) MarshalXML

	e.EncodeToken(start.End())
	return nil
}

func (cl Class) genClassVarDec(e *xml.Encoder) {
	// TODO
}

func genContent(s interface{}) string {
	switch s.(type) {
	case string:
		str, _ := s.(string)
		return " " + str + " "
	case ClassName:
		str, _ := s.(ClassName)
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
