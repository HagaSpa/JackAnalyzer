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
	Modi keyword      // 'var'
	Vt   Types        // identifier | keyword.
	Vns  []identifier // varName (, varName)*
	Sc   symbol       // ';'
}

type Parameter struct {
	Type  Types
	Vn    identifier
	Comma symbol
}

// 'int' | 'char' | 'boolean' | className
// keyword | identifier
type Types interface {
	types()
}

/* Statements */
type Statement interface {
	statement()
}

// LetStatement represent to let.
//
// 'let' varName ( '[' expression ']' )?
// '=' expression ';'
type LetStatement struct {
	Modi   keyword    // 'let'
	Vn     identifier // varName
	LBrack symbol     // '['
	Lexp   Expression // expression
	RBrack symbol     // ']'
	Eq     symbol     // '='
	Rexp   Expression // expression
	Sc     symbol     // ';'
}

func (ls *LetStatement) statement() {}

// IfStatement represent to if.
//
// 'if' '(' expression ')' '{' statements '}'
// ( 'else' '{' statements '}' )?
type IfStatement struct {
	Modi    keyword      // 'if'
	LParan  symbol       // '('
	Lexp    Expression   // expression
	RParen  symbol       // ')'
	LBrace  symbol       // '{'
	Stmts   []Statement  // statements
	RBrace  symbol       // '}'
	Else    *keyword     // 'else'
	Elbrace *symbol      // '{'
	Estmts  []*Statement // statements
	Erbrace *symbol      // '}'
}

func (is *IfStatement) statement() {}

// WhileStatement represent to while.
//
// 'while' '(' expression ')' '{' statements '}'
type WhileStatement struct {
	Modi   keyword     // 'while'
	LParan symbol      // '('
	Lexp   Expression  // expression
	RParen symbol      // ')'
	LBrace symbol      // '{'
	Stmts  []Statement // statements
	RBrace symbol      // '}'
}

func (ws *WhileStatement) statement() {}

// DoStatement represent to do.
//
// 'do' subroutineCall ';'
type DoStatement struct {
	Modi keyword        // 'do'
	Subr SubroutineCall // subroutineCall
	Sc   symbol         // ';'
}

func (do *DoStatement) statement() {}

// ReturnStatement represent to return.
//
// 'return' expression? ';'
type ReturnStatement struct {
	Modi keyword     // 'return'
	Exp  *Expression // expression?
	Sc   symbol      // ';'
}

func (rs *ReturnStatement) statement() {}

/* Expession */
type Expression struct {
	Term Term
	Next BopTerm
}

type BopTerm struct {
	Bop  symbol // binary operator
	Term Term
}

// Term is term
type Term interface {
	term()
}

type IntegerConstant string
type StringConstant string
type KeywordConstant string

// VarName is Term.
//
// varName
type VarName struct {
	V identifier
}

// CallIndex is Term.
//
// varName '[' expression ']'
type CallIndex struct {
	Vn       identifier
	LBracket symbol
	Exp      Expression
	RBracket symbol
}

// SubroutineCall is Term.
//
// subroutineName '(' expressionList ')' |
// (className | varName) '.' subroutineName '(' expressionList ')'
type SubroutineCall struct {
	Name           *identifier  // ClassName | VarName
	Dot            *symbol      // .
	SubName        identifier   // string
	LParan         symbol       // '('
	ExpressionList []Expression // (expression(, expression)*)?
	RParen         symbol       // ')'
}

// Args is Term.
//
// '(' expression ')'
type Args struct {
	LParan symbol
	Exp    Expression
	RParen symbol
}

// UopTerm is Term.
//
// unaryOp term
type UopTerm struct {
	Uop  symbol // unary operator
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
	}
	return " " + str + " "
}

type keyword string
type identifier string
type symbol string

func (k *keyword) types()    {}
func (i *identifier) types() {}

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
