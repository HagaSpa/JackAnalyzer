package element

import (
	"encoding/xml"
)

type Class struct {
	modifier      string          // 'class'
	className     className       // identifier
	lBrace        string          // '{'
	classVarDec   []classVarDec   // classVarDec*
	subtoutineDec []subroutineDec // subroutineDec*
	rBrace        string          // '}'
}

type classVarDec struct {
	modifier  string    // 'static' | 'field'
	varType   types     // 'int' | 'char' | 'boolean' | className
	varNames  []varName // varName (, varName)*
	semiColon string    // ';'
}

type subroutineDec struct {
	modifier      string         // 'constructor' | 'function' | 'method'
	subType       string         // 'void' | type
	subName       subroutineName // identifier
	lParan        string         // '('
	parameterList []parameter    // (type varName (, type, varName)*)?
	rParen        string         // ')'
	subBody       subroutineBody // subroutineBody
}

type subroutineBody struct {
	lBrace     string // '{'
	varDec     varDec // varDec*
	statements []statement
}

type varDec struct {
	modifier  string    // 'var'
	varType   types     // type
	varNames  []varName // varName (, varName)*
	semiColon string    // ';'
}

type statement struct{}

type parameter struct {
	paramType types
	parmaName varName
	comma     string
}
type types string          // 'int' | 'char' | 'boolean' | className
type className string      // identifier
type subroutineName string // identifier
type varName string        // identifier

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
	case className:
		str, _ := s.(className)
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
