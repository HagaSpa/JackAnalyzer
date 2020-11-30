package element

import (
	"encoding/xml"
)

type Class struct {
	Modi   string          // 'class'
	Cn     ClassName       // identifier
	LBrace string          // '{'
	Cvds   []ClassVarDec   // classVarDec*
	Sds    []SubroutineDec // subroutineDec*
	RBrace string          // '}'
}

// ( 'static' | 'field' ) type varName ( ',' varName)* ';'
type ClassVarDec struct {
	Modi string    // 'static' | 'field'
	Vt   Types     // 'int' | 'char' | 'boolean' | className
	Vns  []VarName // varName (, varName)*
	Sc   string    // ';'
}

// ( 'constructor' | 'function' | 'method' )
// ( 'void' | Types ) subroutineName '(' parameterList ')'
// subroutineBody
type SubroutineDec struct {
	Modi   string         // 'constructor' | 'function' | 'method'
	St     string         // 'void' | type
	Sn     SubroutineName // identifier
	LParan string         // '('
	Pl     []Parameter    // (type varName (, type, varName)*)?
	RParen string         // ')'
	Sb     SubroutineBody // subroutineBody
}

// '{' varDec* statements '}'
type SubroutineBody struct {
	LBrace string      // '{'
	Vd     *VarDec     // varDec*
	Stmts  []Statement // statements
	RBrace string      // '}'
}

type VarDec struct {
	Modi string    // 'var'
	Vt   Types     // type
	Vns  []VarName // varName (, varName)*
	Sc   string    // ';'
}

type Parameter struct {
	Type  Types
	Vn    VarName
	Comma string
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
	e.EncodeElement(genContent(cl.Modi), genTagKeyword())
	e.EncodeElement(genContent(cl.Cn), genTagIdentifier())
	e.EncodeElement(genContent(cl.LBrace), genTagSymbol())

	// ClassVarDec
	if len(cl.Cvds) != 0 {
		for _, v := range cl.Cvds {
			v.genClassVarDec(e)
		}
	}

	e.EncodeToken(start.End())
	return nil
}

func (cd ClassVarDec) genClassVarDec(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "classVarDec"}}
	e.EncodeToken(start)
	e.EncodeElement(genContent(cd.Modi), genTagKeyword())
	e.EncodeElement(genContent(cd.Vt), genTagKeyword())

	for i, v := range cd.Vns {
		e.EncodeElement(genContent(v), genTagIdentifier())
		if i < len(cd.Vns)-1 {
			e.EncodeElement(genContent(","), genTagSymbol())
		}
	}

	e.EncodeElement(genContent(cd.Sc), genTagSymbol())
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
