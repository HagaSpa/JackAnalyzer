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

/* Statements */
type Statement interface {
	statement()
}

// LetStatement represent to let.
//
// 'let' varName ( '[' expression ']' )?
// '=' expression ';'
type LetStatement struct {
	Modi   string     // 'let'
	Vn     VarName    // varName
	LBrack string     // '['
	Lexp   Expression // expression
	RBrack string     // ']'
	Eq     string     // '='
	Rexp   Expression // expression
	Sc     string     // ';'
}

func (ls *LetStatement) statement() {}

// IfStatement represent to if.
//
// 'if' '(' expression ')' '{' statements '}'
// ( 'else' '{' statements '}' )?
type IfStatement struct {
	Modi    string       // 'if'
	LParan  string       // '('
	Lexp    Expression   // expression
	RParen  string       // ')'
	LBrace  string       // '{'
	Stmts   []Statement  // statements
	RBrace  string       // '}'
	Else    *string      // 'else'
	Elbrace *string      // '{'
	Estmts  []*Statement // statements
	Erbrace *string      // '}'
}

func (is *IfStatement) statement() {}

// WhileStatement represent to while.
//
// 'while' '(' expression ')' '{' statements '}'
type WhileStatement struct {
	Modi   string      // 'while'
	LParan string      // '('
	Lexp   Expression  // expression
	RParen string      // ')'
	LBrace string      // '{'
	Stmts  []Statement // statements
	RBrace string      // '}'
}

func (ws *WhileStatement) statement() {}

// DoStatement represent to do.
//
// 'do' subroutineCall ';'
type DoStatement struct {
	Modi string         // 'do'
	Subr SubroutineCall // subroutineCall
	Sc   string         // ';'
}

func (do *DoStatement) statement() {}

// ReturnStatement represent to return.
//
// 'return' expression? ';'
type ReturnStatement struct {
	Modi string      // 'return'
	Exp  *Expression // expression?
	Sc   string      // ';'
}

func (rs *ReturnStatement) statement() {}

/* Expession */
type Expression struct {
	Term Term
	Next BopTerm
}

// Term is term
type Term interface {
	term()
}

type IntegerConstant string
type StringConstant string
type KeywordConstant string
type VarName string
type CallIndex struct {
	Vn       VarName
	LBracket string
	Exp      Expression
	RBracket string
}
type SubroutineCall struct {
	Name           string       // ClassName | VarName
	Dot            string       // .
	SubName        string       // string
	LParan         string       // '('
	ExpressionList []Expression // (expression(, expression)*)?
	RParen         string       // ')'
}
type Args struct {
	LParan string
	Exp    Expression
	RParen string
}
type UopTerm struct {
	Uop  string // unary operator
	Term Term
}
type BopTerm struct {
	Bop  Op // binary operator
	Term Term
}

func (ic *IntegerConstant) term() {}
func (sc *StringConstant) term()  {}
func (kc *KeywordConstant) term() {}
func (vn *VarName) term()         {}
func (ci *CallIndex) term()       {}
func (sbc *SubroutineCall) term() {}
func (args *Args) term()          {}
func (ut *UopTerm) term()         {}

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
