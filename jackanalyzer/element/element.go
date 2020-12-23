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
	Modi keyword     // 'let'
	Vn   identifier  // varName
	LB   symbol      // '['
	Lexp *Expression // expression
	RB   symbol      // ']'
	Eq   symbol      // '='
	Rexp Expression  // expression
	Sc   symbol      // ';'
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

// generate Element for *xml.EncodeElement.
func genElement(s interface{}) (string, xml.StartElement) {
	var c string // contents
	var l string // label
	switch v := s.(type) {
	case keyword:
		c = string(v)
		l = "keyword"
	case identifier:
		c = string(v)
		l = "identifier"
	case symbol:
		c = string(v)
		l = "symbol"
	case integerConstant:
		c = strconv.Itoa(int(v))
		l = "integerConstant"
	case stringConstant:
		c = string(v)
		l = "stringConstant"
	}
	return " " + c + " ", xml.StartElement{Name: xml.Name{Local: l}}
}

// MarshalXML implemented Marshaler.
func (cl Class) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// class
	start.Name.Local = "class"
	e.EncodeToken(start)
	e.EncodeElement(genElement(cl.Modi))
	e.EncodeElement(genElement(cl.Cn))
	e.EncodeElement(genElement(cl.LBrace))

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
	e.EncodeElement(genElement(cd.Modi))
	e.EncodeElement(genElement(cd.Vt))
	e.EncodeElement(genElement(cd.Vn))

	for _, v := range cd.Vns {
		e.EncodeElement(genElement(v.Comma))
		e.EncodeElement(genElement(v.Vn))
	}

	e.EncodeElement(genElement(cd.Sc))
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
		e.EncodeElement(genElement(pl.Type))
		e.EncodeElement(genElement(pl.Vn))
		for _, v := range pl.Next {
			e.EncodeElement(genElement(v.Comma))
			e.EncodeElement(genElement(v.Type))
			e.EncodeElement(genElement(v.Vn))
		}
	}
	e.EncodeToken(start.End())
}

func genStatement(s interface{}, e *xml.Encoder) {
	switch s.(type) {
	case *LetStatement:
		// TODO: call genLetStatement
	case *IfStatement:
		// TODO: call genIfStatement
	case *WhileStatement:
		// TODO: call genWhileStatement
	case *DoStatement:
		// TODO: call genDoStatement
	case *ReturnStatement:
		// TODO: call genReturnStatement
	}
}

func (ls *LetStatement) genLetStatement(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "letStatement"}}
	e.EncodeToken(start)
	e.EncodeElement(genElement(ls.Modi))
	e.EncodeElement(genElement(ls.Vn))
	if ls.LB != "" && ls.Lexp != nil && ls.RB != "" {
		e.EncodeElement(genElement(ls.LB))
		ls.Lexp.genExpression(e)
		e.EncodeElement(genElement(ls.RB))
	}
	e.EncodeElement(genElement(ls.Eq))
	ls.Rexp.genExpression(e)
	e.EncodeElement(genElement(ls.Sc))
	e.EncodeToken(start.End())
}

func (exp *Expression) genExpression(e *xml.Encoder) {
	start := xml.StartElement{Name: xml.Name{Local: "expression"}}
	e.EncodeToken(start)
	genTerm(exp.Term, e)
	for _, v := range exp.Next {
		e.EncodeElement(genElement(v.Bop))
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
	case *Args:
		v.genArgs(e)
	case *UopTerm:
		v.genUopTerm(e)
	}
	e.EncodeToken(start.End())
}

func (ic *IntegerConstant) genIntegerConstant(e *xml.Encoder) {
	e.EncodeElement(genElement(ic.V))
}

func (sc *StringConstant) genStringConstant(e *xml.Encoder) {
	e.EncodeElement(genElement(sc.V))
}

func (kc *KeywordConstant) genKeywordConstant(e *xml.Encoder) {
	e.EncodeElement(genElement(kc.V))
}

func (vn *VarName) genVarName(e *xml.Encoder) {
	e.EncodeElement(genElement(vn.V))
}

func (ci *CallIndex) genCallIndex(e *xml.Encoder) {
	e.EncodeElement(genElement(ci.Vn))
	e.EncodeElement(genElement(ci.LB))
	ci.Exp.genExpression(e)
	e.EncodeElement(genElement(ci.RB))
}

func (sbc *SubroutineCall) genSubroutineCall(e *xml.Encoder) {
	if sbc.Name != "" && sbc.Dot != "" {
		e.EncodeElement(genElement(sbc.Name))
		e.EncodeElement(genElement(sbc.Dot))
	}
	e.EncodeElement(genElement(sbc.Sn))
	e.EncodeElement(genElement(sbc.LP))
	start := xml.StartElement{Name: xml.Name{Local: "expressionList"}}
	e.EncodeToken(start)
	for _, v := range sbc.ExpL {
		v.genExpression(e)
	}
	e.EncodeToken(start.End())
	e.EncodeElement(genElement(sbc.RP))
}

func (args *Args) genArgs(e *xml.Encoder) {
	e.EncodeElement(genElement(args.LP))
	args.Exp.genExpression(e)
	e.EncodeElement(genElement(args.RP))
}

func (ut *UopTerm) genUopTerm(e *xml.Encoder) {
	e.EncodeElement(genElement(ut.Uop))
	genTerm(ut.Term, e)
}
