package element

type class struct {
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
	subBody       subroutineBody
}

type subroutineBody interface{}
type parameter struct {
	paramType types
	parmaName varName
	comma     string
}
type types string          // 'int' | 'char' | 'boolean' | className
type className string      // identifier
type subroutineName string // identifier
type varName string        // identifier
