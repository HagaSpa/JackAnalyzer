package element

import (
	"encoding/xml"
)

// 'class' className '{' classVarDec* subroutineDec* '}'
type Class struct {
	Modi   keyword          // 'class'
	Cn     identifier       // identifier
	LBrace symbol           // '{'
	Cvds   []*ClassVarDec   // classVarDec*
	Sds    []*SubroutineDec // subroutineDec*
	RBrace symbol           // '}'
}

// ( 'static' | 'field' ) type varName ( ',' varName)* ';'
type ClassVarDec struct {
	Modi keyword      // 'static' | 'field'
	Vt   keyword      // 'int' | 'char' | 'boolean' | className
	Vns  []identifier // varName (, varName)*
	Sc   symbol       // ';'
}

// ( 'constructor' | 'function' | 'method' )
// ( 'void' | Types ) subroutineName '(' parameterList ')'
// subroutineBody
type SubroutineDec struct {
	Modi   keyword        // 'constructor' | 'function' | 'method'
	St     keyword        // 'void' | type
	Sn     identifier     // subroutineName
	LParan symbol         // '('
	Pl     []Parameter    // (type varName (, type, varName)*)?
	RParen symbol         // ')'
	Sb     SubroutineBody // subroutineBody
}

// '{' varDec* statements '}'
type SubroutineBody struct {
	LBrace symbol      // '{'
	Vd     *VarDec     // varDec*
	Stmts  []Statement // statements
	RBrace symbol      // '}'
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
	e.EncodeElement(genContent(cl.Modi), genTrmSymTag(cl.Modi))
	e.EncodeElement(genContent(cl.Cn), genTrmSymTag(cl.Cn))
	e.EncodeElement(genContent(cl.LBrace), genTrmSymTag(cl.LBrace))

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
	e.EncodeElement(genContent(cd.Modi), genTrmSymTag(cd.Modi))
	e.EncodeElement(genContent(cd.Vt), genTrmSymTag(cd.Vt))

	for i, v := range cd.Vns {
		e.EncodeElement(genContent(v), genTrmSymTag(v))
		if i < len(cd.Vns)-1 {
			s := symbol(",")
			e.EncodeElement(genContent(s), genTrmSymTag(s))
		}
	}

	e.EncodeElement(genContent(cd.Sc), genTrmSymTag(cd.Sc))
	e.EncodeToken(start.End())
}

func genContent(s interface{}) string {
	var str string
	switch s.(type) {
	case string:
		str = s.(string)
	case keyword:
		str = string(s.(keyword))
	case identifier:
		str = string(s.(identifier))
	case symbol:
		str = string(s.(symbol))
	case ClassName:
		str = string(s.(ClassName))
	case Types:
		str = string(s.(Types))
	case VarName:
		str = string(s.(VarName))
	}
	return " " + str + " "
}

type keyword string
type identifier string
type symbol string

// generate xml tag for terminal symbol.
func genTrmSymTag(s interface{}) xml.StartElement {
	var l string
	switch s.(type) {
	case keyword:
		l = "keyword"
	case identifier:
		l = "identifier"
	case symbol:
		l = "symbol"
	}
	return xml.StartElement{Name: xml.Name{Local: l}}
}
