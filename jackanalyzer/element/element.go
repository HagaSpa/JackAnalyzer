package element

import (
	"encoding/xml"
	"strconv"
)

/*
Terminal Symbol
*/

type terminalSymbol interface {
}

// Same as *token.Keyword*
//  'class', 'method', 'function', 'constructor', 'int', 'boolean', 'char', 'void', 'var', 'static', 'field', 'let', 'do', 'if', 'else', 'while', 'return', 'true', 'false', 'null', 'this'
type keyword string

// Alphabet, number, underscore string.
//
// However, character strings starting with numbers are excluded
type identifier string

// Same as *token.symbols*
//  '{', '}', '(', ')', '[', ']', '.', ',', ';', '+', '-', '*', '/', '&', '|', '<', '>', '=', '~'
type symbol string

// 0 ~ 32767
type integerConstant int

// Unicode string without double quotes and newlines
type stringConstant string

/*
Program
*/

// Class represent to class.
//
//  'class' className '{' classVarDec* subroutineDec* '}'
type Class struct {
	Modi   keyword          // 'class'
	Cn     identifier       // identifier
	LBrace symbol           // '{'
	Cvds   []*ClassVarDec   // classVarDec*
	Sds    []*SubroutineDec // subroutineDec*
	RBrace symbol           // '}'
}

// ClassVarDec represent to classVarDec.
//
//  ( 'static' | 'field' ) type varName (',' varName)* ';'
type ClassVarDec struct {
	Modi keyword    // 'static' | 'field'
	Vt   keyword    // type
	Vn   identifier // varName
	Vns  []*NextVns // (',' varName)*
	Sc   symbol     // ';'
}

// NextVns is Next varNames.
//
//  (',' varName)*
type NextVns struct {
	Comma symbol
	Vn    identifier
}

// SubroutineDec represent to subroutineDec.
//
//  ( 'constructor' | 'function' | 'method' )
//  ( 'void' | Types ) subroutineName '(' parameterList ')'
//  subroutineBody
type SubroutineDec struct {
	Modi   keyword        // 'constructor' | 'function' | 'method'
	St     keyword        // 'void' | type
	Sn     identifier     // subroutineName
	LParen symbol         // '('
	Pl     ParameterList  // parameterList
	RParen symbol         // ')'
	Sb     SubroutineBody // subroutineBody
}

// ParameterList represent to parameterList.
//
//  (type varName (',' type varName)* )?
type ParameterList struct {
	Type Types
	Vn   identifier
	Next []*NextParam
}

// NextParam is the second and subsequent elements of ParameterList.
type NextParam struct {
	Comma symbol
	Type  Types
	Vn    identifier
}

// SubroutineBody represent to subroutineBody.
//
//  '{' varDec* statements '}'
type SubroutineBody struct {
	LBrace symbol       // '{'
	Vd     *VarDec      // varDec*
	Stmts  []*Statement // statements
	RBrace symbol       // '}'
}

// VarDec represent to varDec.
//
//  'var' type varName (',' varName)* ';'
type VarDec struct {
	Modi keyword    // 'var'
	Vt   Types      // type
	Vn   identifier // varName
	Vns  []*NextVns // (',' varName)*
	Sc   symbol     // ';'
}

// Types represent to type.
//
//  keyword ('int' | 'char' | 'boolean') | identifier (className)
type Types interface {
	types()
}

func (k keyword) types()    {}
func (i identifier) types() {}

/*
Statement
*/

// Statement is statements
//  statement*
type Statement interface {
	statement()
}

// LetStatement represent to let.
//
//  'let' varName ( '[' expression ']' )? '=' expression ';'
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
//  'if' '(' expression ')' '{' statements '}'
//  ( 'else' '{' statements '}' )?
type IfStatement struct {
	Modi    keyword      // 'if'
	LParen  symbol       // '('
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
//  'while' '(' expression ')' '{' statements '}'
type WhileStatement struct {
	Modi   keyword     // 'while'
	LParen symbol      // '('
	Lexp   Expression  // expression
	RParen symbol      // ')'
	LBrace symbol      // '{'
	Stmts  []Statement // statements
	RBrace symbol      // '}'
}

func (ws *WhileStatement) statement() {}

// DoStatement represent to do.
//
//  'do' subroutineCall ';'
type DoStatement struct {
	Modi keyword        // 'do'
	Subr SubroutineCall // subroutineCall
	Sc   symbol         // ';'
}

func (do *DoStatement) statement() {}

// ReturnStatement represent to return.
//
//  'return' expression? ';'
type ReturnStatement struct {
	Modi keyword     // 'return'
	Exp  *Expression // expression?
	Sc   symbol      // ';'
}

func (rs *ReturnStatement) statement() {}

/*
Expession
*/

// Expression is expression
type Expression struct {
	Term Term
	Next []*BopTerm
}

// BopTerm is Binary Operator Term
type BopTerm struct {
	Bop  symbol // binary operator
	Term Term
}

// Term is term
type Term interface {
	term()
}

// IntegerConstant is Term.
type IntegerConstant struct {
	V integerConstant
}

// StringConstant is Term.
type StringConstant struct {
	V stringConstant
}

// KeywordConstant is Term.
//
//  'true' | 'false' | 'null' | 'this'
type KeywordConstant struct {
	V keyword
}

// VarName is Term.
//
//  varName
type VarName struct {
	V identifier
}

// CallIndex is Term.
//
//  varName '[' expression ']'
type CallIndex struct {
	Vn  identifier
	LB  symbol
	Exp Expression
	RB  symbol
}

// SubroutineCall is Term.
//
//  subroutineName '(' expressionList ')' |
//  (className | varName) '.' subroutineName '(' expressionList ')'
type SubroutineCall struct {
	Name identifier   // ClassName | VarName
	Dot  symbol       // .
	Sn   identifier   // string
	LP   symbol       // '('
	ExpL []Expression // (expression(, expression)*)?
	RP   symbol       // ')'
}

// Args is Term.
//
//  '(' expression ')'
type Args struct {
	LP  symbol     // '('
	Exp Expression // expression
	RP  symbol     // ')'
}

// UopTerm is Term.
//
//  unaryOp term
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

// MarshalXML implemented Marshaler.
func (cl Class) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// class
	start.Name.Local = "class"
	e.EncodeToken(start)
	e.EncodeElement(genCon(cl.Modi), genTag(cl.Modi))
	e.EncodeElement(genCon(cl.Cn), genTag(cl.Cn))
	e.EncodeElement(genCon(cl.LBrace), genTag(cl.LBrace))

	// ClassVarDec
	if len(cl.Cvds) != 0 {
		for _, v := range cl.Cvds {
			v.genClassVarDec(e)
		}
	}

	e.EncodeToken(start.End())
	return nil
}

func (cd *ClassVarDec) genClassVarDec(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "classVarDec"}}
	e.EncodeToken(start)
	e.EncodeElement(genCon(cd.Modi), genTag(cd.Modi))
	e.EncodeElement(genCon(cd.Vt), genTag(cd.Vt))
	e.EncodeElement(genCon(cd.Vn), genTag(cd.Vn))

	for _, v := range cd.Vns {
		e.EncodeElement(genCon(v.Comma), genTag(v.Comma))
		e.EncodeElement(genCon(v.Vn), genTag(v.Vn))
	}

	e.EncodeElement(genCon(cd.Sc), genTag(cd.Sc))
	e.EncodeToken(start.End())
}

func (pl *ParameterList) genParameterList(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "parameterList"}}
	e.EncodeToken(start)
	if pl == nil {
		// insert new line
		c := xml.CharData([]byte("\n"))
		e.EncodeToken(c)
	} else {
		e.EncodeElement(genCon(pl.Type), genTag(pl.Type))
		e.EncodeElement(genCon(pl.Vn), genTag(pl.Vn))
		for _, v := range pl.Next {
			e.EncodeElement(genCon(v.Comma), genTag(v.Comma))
			e.EncodeElement(genCon(v.Type), genTag(v.Type))
			e.EncodeElement(genCon(v.Vn), genTag(v.Vn))
		}
	}
	e.EncodeToken(start.End())
}

func (exp *Expression) genExpression(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "expression"}}
	e.EncodeToken(start)
	genTerm(exp.Term, e)
	for _, v := range exp.Next {
		e.EncodeElement(genCon(v.Bop), genTag(v.Bop))
		genTerm(v.Term, e)
	}
	e.EncodeToken(start.End())
}

func genTerm(s interface{}, e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "term"}}
	e.EncodeToken(start)
	switch v := s.(type) {
	case *IntegerConstant:
		v.genIntegerConstant(e)
	case *StringConstant:
		v.genStringConstant(e)
	case *KeywordConstant:
		v.genKeywordConstant(e)
	case *VarName:
		v.genVarName(e)
	case *CallIndex:
		v.genCallIndex(e)
	case *SubroutineCall:
		v.genSubroutineCall(e)
	case Args:
		// call genArgs
	case UopTerm:
		// call genUopTerm
	}
	e.EncodeToken(start.End())
}

func (ic *IntegerConstant) genIntegerConstant(e *xml.Encoder) {
	e.EncodeElement(genCon(ic.V), genTag(ic.V))
}

func (sc *StringConstant) genStringConstant(e *xml.Encoder) {
	e.EncodeElement(genCon(sc.V), genTag(sc.V))
}

func (kc *KeywordConstant) genKeywordConstant(e *xml.Encoder) {
	e.EncodeElement(genCon(kc.V), genTag(kc.V))
}

func (vn *VarName) genVarName(e *xml.Encoder) {
	e.EncodeElement(genCon(vn.V), genTag(vn.V))
}

func (ci *CallIndex) genCallIndex(e *xml.Encoder) {
	e.EncodeElement(genCon(ci.Vn), genTag(ci.Vn))
	e.EncodeElement(genCon(ci.LB), genTag(ci.LB))
	ci.Exp.genExpression(e)
	e.EncodeElement(genCon(ci.RB), genTag(ci.RB))
}

func (sbc *SubroutineCall) genSubroutineCall(e *xml.Encoder) {
	if sbc.Name != "" && sbc.Dot != "" {
		e.EncodeElement(genCon(sbc.Name), genTag(sbc.Name))
		e.EncodeElement(genCon(sbc.Dot), genTag(sbc.Dot))
	}
	e.EncodeElement(genCon(sbc.Sn), genTag(sbc.Sn))
	e.EncodeElement(genCon(sbc.LP), genTag(sbc.RP))
	start := xml.StartElement{Name: xml.Name{Local: "expressionList"}}
	e.EncodeToken(start)
	for _, v := range sbc.ExpL {
		v.genExpression(e)
	}
	e.EncodeToken(start.End())
	e.EncodeElement(genCon(sbc.RP), genTag(sbc.RP))
}

func (args *Args) genArgs(e *xml.Encoder) {
	e.EncodeElement(genCon(args.LP), genTag(args.LP))
	args.Exp.genExpression(e)
	e.EncodeElement(genCon(args.RP), genTag(args.RP))
}

// generate Contents for terminal symbol.
func genCon(s interface{}) string {
	var str string
	switch v := s.(type) {
	case keyword:
		str = string(v)
	case identifier:
		str = string(v)
	case symbol:
		str = string(v)
	case integerConstant:
		str = strconv.Itoa(int(v))
	case stringConstant:
		str = string(v)
	}
	return " " + str + " "
}

// generate xml tag for terminal symbol.
func genTag(s interface{}) xml.StartElement {
	var l string
	switch s.(type) {
	case keyword:
		l = "keyword"
	case identifier:
		l = "identifier"
	case symbol:
		l = "symbol"
	case integerConstant:
		l = "integerConstant"
	case stringConstant:
		l = "stringConstant"
	}
	return xml.StartElement{Name: xml.Name{Local: l}}
}
