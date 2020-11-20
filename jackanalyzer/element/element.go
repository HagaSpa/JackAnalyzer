package element

import (
	"encoding/xml"
)

type Class struct {
	modifier      string          // 'class'
	className     ClassName       // identifier
	lBrace        string          // '{'
	classVarDec   []ClassVarDec   // classVarDec*
	subtoutineDec []SubroutineDec // subroutineDec*
	rBrace        string          // '}'
}

type ClassVarDec struct {
	modifier  string    // 'static' | 'field'
	varType   Types     // 'int' | 'char' | 'boolean' | className
	varNames  []VarName // varName (, varName)*
	semiColon string    // ';'
}

type SubroutineDec struct {
	modifier      string         // 'constructor' | 'function' | 'method'
	subType       string         // 'void' | type
	subName       SubroutineName // identifier
	lParan        string         // '('
	parameterList []Parameter    // (type varName (, type, varName)*)?
	rParen        string         // ')'
	subBody       SubroutineBody // subroutineBody
}

type SubroutineBody struct {
	lBrace     string // '{'
	varDec     VarDec // varDec*
	statements []Statement
}

type VarDec struct {
	modifier  string    // 'var'
	varType   Types     // type
	varNames  []VarName // varName (, varName)*
	semiColon string    // ';'
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
	e.EncodeElement(genContent(cl.modifier), genTagKeyword())
	e.EncodeElement(genContent(cl.className), genTagIdentifier())
	e.EncodeElement(genContent(cl.lBrace), genTagSymbol())
	e.EncodeToken(start.End())
	return nil
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
