package element

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func initEncBytes() (*xml.Encoder, *bytes.Buffer) {
	var b bytes.Buffer
	e := xml.NewEncoder(&b)
	e.Indent("", "  ")
	return e, &b
}

func Test_genElement(t *testing.T) {
	tests := []struct {
		name  string
		s     interface{}
		want  string
		want1 xml.StartElement
	}{
		{
			"test keyword",
			keyword("class"),
			" class ",
			xml.StartElement{
				Name: xml.Name{
					Local: "keyword",
				},
			},
		},
		{
			"test identifier",
			identifier("hoge"),
			" hoge ",
			xml.StartElement{
				Name: xml.Name{
					Local: "identifier",
				},
			},
		},
		{
			"test symbol",
			symbol(","),
			" , ",
			xml.StartElement{
				Name: xml.Name{
					Local: "symbol",
				},
			},
		},
		{
			"test integerConstant",
			integerConstant(123),
			" 123 ",
			xml.StartElement{
				Name: xml.Name{
					Local: "integerConstant",
				},
			},
		},
		{
			"test stringConstant",
			stringConstant("THE AVERAGE IS:"),
			" THE AVERAGE IS: ",
			xml.StartElement{
				Name: xml.Name{
					Local: "stringConstant",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := genElement(tt.s)
			if got != tt.want {
				t.Errorf("genElement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("genElement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_class_MarshalXML(t *testing.T) {
	tests := []struct {
		name string
		cl   Class
		want string
	}{
		{
			"test",
			Class{
				Modi:   "class",
				Cn:     "Main",
				LBrace: "{",
			},
			`
<class>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
</class>
`,
		},
		{
			"test2",
			Class{
				Modi:   "class",
				Cn:     "Main",
				LBrace: "{",
				Cvds: []*ClassVarDec{
					{
						Modi: "field",
						Vt:   "int",
						Vn:   "x",
						Vns: []*NextVns{
							{
								Comma: ",",
								Vn:    "y",
							},
						},
						Sc: ";",
					},
				},
			},
			`
<class>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
  <classVarDec>
    <keyword> field </keyword>
    <keyword> int </keyword>
    <identifier> x </identifier>
    <symbol> , </symbol>
    <identifier> y </identifier>
    <symbol> ; </symbol>
  </classVarDec>
</class>
`,
		},
		{
			"classVarDed loop",
			Class{
				Modi:   "class",
				Cn:     "Main",
				LBrace: "{",
				Cvds: []*ClassVarDec{
					{
						Modi: "field",
						Vt:   "int",
						Vn:   "x",
						Vns: []*NextVns{
							{
								Comma: ",",
								Vn:    "y",
							},
						},
						Sc: ";",
					},
					{
						Modi: "field",
						Vt:   "int",
						Vn:   "size",
						Sc:   ";",
					},
				},
			},
			`
<class>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
  <classVarDec>
    <keyword> field </keyword>
    <keyword> int </keyword>
    <identifier> x </identifier>
    <symbol> , </symbol>
    <identifier> y </identifier>
    <symbol> ; </symbol>
  </classVarDec>
  <classVarDec>
    <keyword> field </keyword>
    <keyword> int </keyword>
    <identifier> size </identifier>
    <symbol> ; </symbol>
  </classVarDec>
</class>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, _ := xml.MarshalIndent(tt.cl, "", "  ")
			// trim \n
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(string(output), want) {
				t.Errorf("class.MarshalXML() = %v", string(output))
				t.Errorf("wantXml = %v", want)
			}
		})
	}
}

func TestClassVarDec_genClassVarDec(t *testing.T) {
	tests := []struct {
		name string
		cd   *ClassVarDec
		want string
	}{
		{
			"test",
			&ClassVarDec{
				Modi: "field",
				Vt:   "int",
				Vn:   "x",
				Vns: []*NextVns{
					{
						Comma: ",",
						Vn:    "y",
					},
				},
				Sc: ";",
			},
			`
<classVarDec>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> ; </symbol>
</classVarDec>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.cd.genClassVarDec(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("cd.genClassVarDec() = %v", got)
				t.Errorf("wantXml = %v", want)
			}
		})
	}
}

func TestParameterList_genParameterList(t *testing.T) {
	tests := []struct {
		name string
		pl   *ParameterList
		want string
	}{
		{
			"test keyword",
			&ParameterList{
				Type: keyword("int"),
				Vn:   "Ax",
			},
			`
<parameterList>
  <keyword> int </keyword>
  <identifier> Ax </identifier>
</parameterList>
`,
		},
		{
			"test identifier",
			&ParameterList{
				Type: identifier("Hoge"),
				Vn:   "Ax",
			},
			`
<parameterList>
  <identifier> Hoge </identifier>
  <identifier> Ax </identifier>
</parameterList>
`,
		},
		{
			"test no parameter",
			nil,
			`
<parameterList></parameterList>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.pl.genParameterList(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genParameterList() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestExpression_genExpression(t *testing.T) {
	tests := []struct {
		name string
		exp  *Expression
		want string
	}{
		{
			"test",
			&Expression{
				Term: &IntegerConstant{
					V: 1,
				},
				Next: nil,
			},
			`
<expression>
  <term>
    <integerConstant> 1 </integerConstant>
  </term>
</expression>
`,
		},
		{
			"test Next",
			&Expression{
				Term: &IntegerConstant{
					V: 1,
				},
				Next: []*BopTerm{
					{
						Bop: "+",
						Term: &IntegerConstant{
							V: 2,
						},
					},
				},
			},
			`
<expression>
  <term>
    <integerConstant> 1 </integerConstant>
  </term>
  <symbol> + </symbol>
  <term>
    <integerConstant> 2 </integerConstant>
  </term>
</expression>
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.exp.genExpression(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genExpression() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func Test_genTerm(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		want string
	}{
		{
			"test IntegerConstant",
			&IntegerConstant{
				V: 134,
			},
			`
<term>
  <integerConstant> 134 </integerConstant>
</term>
`,
		},
		{
			"test StringConstant",
			&StringConstant{
				V: "test",
			},
			`
<term>
  <stringConstant> test </stringConstant>
</term>
`,
		},
		{
			"test KeywordConstant",
			&KeywordConstant{
				V: "int",
			},
			`
<term>
  <keyword> int </keyword>
</term>
`,
		},
		{
			"test VarName",
			&VarName{
				V: "Hoge",
			},
			`
<term>
  <identifier> Hoge </identifier>
</term>
`,
		},
		{
			"test CallIndex",
			&CallIndex{
				Vn: "a",
				LB: "[",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RB: "]",
			},
			`
<term>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
  <symbol> ] </symbol>
</term>
`,
		},
		{
			"test SubroutineCall",
			&SubroutineCall{
				Name: "Main",
				Dot:  ".",
				Sn:   "main",
				LP:   "(",
				ExpL: []Expression{
					{
						Term: &VarName{
							V: "i",
						},
					},
				},
				RP: ")",
			},
			`
<term>
  <identifier> Main </identifier>
  <symbol> . </symbol>
  <identifier> main </identifier>
  <symbol> ( </symbol>
  <expressionList>
    <expression>
      <term>
        <identifier> i </identifier>
      </term>
    </expression>
  </expressionList>
  <symbol> ) </symbol>
</term>
`,
		},
		{
			"test Args",
			&Args{
				LP: "(",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RP: ")",
			},
			`
<term>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
</term>
`,
		},
		{
			"test UopTerm",
			&UopTerm{
				Uop: "-",
				Term: &VarName{
					V: "i",
				},
			},
			`
<term>
  <symbol> - </symbol>
  <term>
    <identifier> i </identifier>
  </term>
</term>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			genTerm(tt.s, e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genTerm() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestIntegerConstant_genIntegerConstant(t *testing.T) {
	tests := []struct {
		name string
		ic   *IntegerConstant
		want string
	}{
		{
			"test",
			&IntegerConstant{
				V: 123,
			},
			`
<integerConstant> 123 </integerConstant>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.ic.genIntegerConstant(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genIntegerConstant() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestStringConstant_genStringConstant(t *testing.T) {
	tests := []struct {
		name string
		sc   StringConstant
		want string
	}{
		{
			"test",
			StringConstant{
				V: "THE AVERAGE IS:",
			},
			`
<stringConstant> THE AVERAGE IS: </stringConstant>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.sc.genStringConstant(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genStringConstant() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestKeywordConstant_genKeywordConstant(t *testing.T) {
	tests := []struct {
		name string
		kc   *KeywordConstant
		want string
	}{
		{
			"test",
			&KeywordConstant{
				V: "int",
			},
			`
<keyword> int </keyword>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.kc.genKeywordConstant(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genKeywordConstant() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestVarName_genVarName(t *testing.T) {
	tests := []struct {
		name string
		vn   *VarName
		want string
	}{
		{
			"test",
			&VarName{
				V: "Hoge",
			},
			`
<identifier> Hoge </identifier>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.vn.genVarName(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genVarName() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestCallIndex_genCallIndex(t *testing.T) {
	tests := []struct {
		name string
		ci   *CallIndex
		want string
	}{
		{
			"test",
			&CallIndex{
				Vn: "a",
				LB: "[",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RB: "]",
			},
			`
<identifier> a </identifier>
<symbol> [ </symbol>
<expression>
  <term>
    <identifier> i </identifier>
  </term>
</expression>
<symbol> ] </symbol>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.ci.genCallIndex(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genCallIndex() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestSubroutineCall_genSubroutineCall(t *testing.T) {
	tests := []struct {
		name string
		sbc  *SubroutineCall
		want string
	}{
		{
			"test",
			&SubroutineCall{
				Name: "Main",
				Dot:  ".",
				Sn:   "main",
				LP:   "(",
				ExpL: []Expression{
					{
						Term: &VarName{
							V: "i",
						},
					},
				},
				RP: ")",
			},
			`
<identifier> Main </identifier>
<symbol> . </symbol>
<identifier> main </identifier>
<symbol> ( </symbol>
<expressionList>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
</expressionList>
<symbol> ) </symbol>
`,
		},
		{
			"test Name and Dot is not exist.",
			&SubroutineCall{
				Sn: "main",
				LP: "(",
				ExpL: []Expression{
					{
						Term: &VarName{
							V: "i",
						},
					},
				},
				RP: ")",
			},
			`
<identifier> main </identifier>
<symbol> ( </symbol>
<expressionList>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
</expressionList>
<symbol> ) </symbol>
`,
		},
		{
			"testing when ExpL has multiple Bop.",
			&SubroutineCall{
				Sn: "main",
				LP: "(",
				ExpL: []Expression{
					{
						Term: &VarName{
							V: "i",
						},
						Next: []*BopTerm{
							{
								Bop: "+",
								Term: &VarName{
									V: "j",
								},
							},
						},
					},
				},
				RP: ")",
			},
			`
<identifier> main </identifier>
<symbol> ( </symbol>
<expressionList>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
    <symbol> + </symbol>
    <term>
      <identifier> j </identifier>
    </term>
  </expression>
</expressionList>
<symbol> ) </symbol>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.sbc.genSubroutineCall(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genSubroutineCall() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestArgs_genArgs(t *testing.T) {
	tests := []struct {
		name string
		args *Args
		want string
	}{
		{
			"test",
			&Args{
				LP: "(",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RP: ")",
			},
			`
<symbol> ( </symbol>
<expression>
  <term>
    <identifier> i </identifier>
  </term>
</expression>
<symbol> ) </symbol>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.args.genArgs(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genArgs() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestUopTerm_genUopTerm(t *testing.T) {
	tests := []struct {
		name string
		ut   *UopTerm
		want string
	}{
		{
			"test",
			&UopTerm{
				Uop: "-",
				Term: &VarName{
					V: "i",
				},
			},
			`
<symbol> - </symbol>
<term>
  <identifier> i </identifier>
</term>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.ut.genUopTerm(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genUopTerm() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func Test_genStatement(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		want string
	}{
		{
			"test LetStatement",
			&LetStatement{
				Modi: "let",
				Vn:   "a",
				Eq:   "=",
				Rexp: Expression{
					Term: &SubroutineCall{
						Name: "Array",
						Dot:  ".",
						Sn:   "new",
						LP:   "(",
						ExpL: []Expression{
							{
								Term: &VarName{
									V: "length",
								},
							},
						},
						RP: ")",
					},
				},
				Sc: ";",
			},
			`
<letStatement>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> = </symbol>
  <expression>
    <term>
      <identifier> Array </identifier>
      <symbol> . </symbol>
      <identifier> new </identifier>
      <symbol> ( </symbol>
      <expressionList>
        <expression>
          <term>
            <identifier> length </identifier>
          </term>
        </expression>
      </expressionList>
      <symbol> ) </symbol>
    </term>
  </expression>
  <symbol> ; </symbol>
</letStatement>
`,
		},
		{
			"test IfStatement",
			&IfStatement{
				Modi: "if",
				LP:   "(",
				LExp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RP: ")",
				LB: "{",
				Stmts: []Statement{
					&LetStatement{
						Modi: "let",
						Vn:   "s",
						Eq:   "=",
						Rexp: Expression{
							Term: &VarName{
								V: "i",
							},
						},
						Sc: ";",
					},
				},
				RB: "}",
			},
			`
<ifStatement>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <statements>
    <letStatement>
      <keyword> let </keyword>
      <identifier> s </identifier>
      <symbol> = </symbol>
      <expression>
        <term>
          <identifier> i </identifier>
        </term>
      </expression>
      <symbol> ; </symbol>
    </letStatement>
  </statements>
  <symbol> } </symbol>
</ifStatement>
`,
		},
		{
			"test WhileStatement",
			&WhileStatement{
				Modi: "while",
				LP:   "(",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
					Next: []*BopTerm{
						{
							Bop: "<",
							Term: &VarName{
								V: "length",
							},
						},
					},
				},
				RP: ")",
				LB: "{",
				Stmts: []Statement{
					&LetStatement{
						Modi: "let",
						Vn:   "i",
						Eq:   "=",
						Rexp: Expression{
							Term: &VarName{
								V: "i",
							},
							Next: []*BopTerm{
								{
									Bop: "+",
									Term: &IntegerConstant{
										V: 1,
									},
								},
							},
						},
						Sc: ";",
					},
				},
				RB: "}",
			},
			`
<whileStatement>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
    <symbol> &lt; </symbol>
    <term>
      <identifier> length </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <statements>
    <letStatement>
      <keyword> let </keyword>
      <identifier> i </identifier>
      <symbol> = </symbol>
      <expression>
        <term>
          <identifier> i </identifier>
        </term>
        <symbol> + </symbol>
        <term>
          <integerConstant> 1 </integerConstant>
        </term>
      </expression>
      <symbol> ; </symbol>
    </letStatement>
  </statements>
  <symbol> } </symbol>
</whileStatement>
`,
		},
		{
			"test DoStatement",
			&DoStatement{
				Modi: "do",
				Sub: &SubroutineCall{
					Name: "Output",
					Dot:  ".",
					Sn:   "printString",
					LP:   "(",
					ExpL: []Expression{
						{
							Term: &StringConstant{
								V: "THE AVERAGE IS: ",
							},
						},
					},
					RP: ")",
				},
				Sc: ";",
			},
			`
<doStatement>
  <keyword> do </keyword>
  <identifier> Output </identifier>
  <symbol> . </symbol>
  <identifier> printString </identifier>
  <symbol> ( </symbol>
  <expressionList>
    <expression>
      <term>
        <stringConstant> THE AVERAGE IS:  </stringConstant>
      </term>
    </expression>
  </expressionList>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
</doStatement>
`,
		},
		{
			"test ReturnStatement",
			&ReturnStatement{
				Modi: "return",
				Sc:   ";",
			},
			`
<returnStatement>
  <keyword> return </keyword>
  <symbol> ; </symbol>
</returnStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			genStatement(tt.s, e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestLetStatement_genLetStatement(t *testing.T) {
	tests := []struct {
		name string
		ls   *LetStatement
		want string
	}{
		{
			"test let a = Array.new(length);",
			&LetStatement{
				Modi: "let",
				Vn:   "a",
				Eq:   "=",
				Rexp: Expression{
					Term: &SubroutineCall{
						Name: "Array",
						Dot:  ".",
						Sn:   "new",
						LP:   "(",
						ExpL: []Expression{
							{
								Term: &VarName{
									V: "length",
								},
							},
						},
						RP: ")",
					},
				},
				Sc: ";",
			},
			`
<letStatement>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> = </symbol>
  <expression>
    <term>
      <identifier> Array </identifier>
      <symbol> . </symbol>
      <identifier> new </identifier>
      <symbol> ( </symbol>
      <expressionList>
        <expression>
          <term>
            <identifier> length </identifier>
          </term>
        </expression>
      </expressionList>
      <symbol> ) </symbol>
    </term>
  </expression>
  <symbol> ; </symbol>
</letStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.ls.genLetStatement(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genLetStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestIfStatement_genIfStatement(t *testing.T) {
	tests := []struct {
		name string
		is   *IfStatement
		want string
	}{
		{
			"test if (i) { let s = i; }",
			&IfStatement{
				Modi: "if",
				LP:   "(",
				LExp: Expression{
					Term: &VarName{
						V: "i",
					},
				},
				RP: ")",
				LB: "{",
				Stmts: []Statement{
					&LetStatement{
						Modi: "let",
						Vn:   "s",
						Eq:   "=",
						Rexp: Expression{
							Term: &VarName{
								V: "i",
							},
						},
						Sc: ";",
					},
				},
				RB: "}",
			},
			`
<ifStatement>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <statements>
    <letStatement>
      <keyword> let </keyword>
      <identifier> s </identifier>
      <symbol> = </symbol>
      <expression>
        <term>
          <identifier> i </identifier>
        </term>
      </expression>
      <symbol> ; </symbol>
    </letStatement>
  </statements>
  <symbol> } </symbol>
</ifStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.is.genIfStatement(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genIfStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestWhileStatement_genWhileStatement(t *testing.T) {
	tests := []struct {
		name string
		ws   *WhileStatement
		want string
	}{
		{
			"test while (i < length) {let i = i + 1; }",
			&WhileStatement{
				Modi: "while",
				LP:   "(",
				Exp: Expression{
					Term: &VarName{
						V: "i",
					},
					Next: []*BopTerm{
						{
							Bop: "<",
							Term: &VarName{
								V: "length",
							},
						},
					},
				},
				RP: ")",
				LB: "{",
				Stmts: []Statement{
					&LetStatement{
						Modi: "let",
						Vn:   "i",
						Eq:   "=",
						Rexp: Expression{
							Term: &VarName{
								V: "i",
							},
							Next: []*BopTerm{
								{
									Bop: "+",
									Term: &IntegerConstant{
										V: 1,
									},
								},
							},
						},
						Sc: ";",
					},
				},
				RB: "}",
			},
			`
<whileStatement>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <expression>
    <term>
      <identifier> i </identifier>
    </term>
    <symbol> &lt; </symbol>
    <term>
      <identifier> length </identifier>
    </term>
  </expression>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <statements>
    <letStatement>
      <keyword> let </keyword>
      <identifier> i </identifier>
      <symbol> = </symbol>
      <expression>
        <term>
          <identifier> i </identifier>
        </term>
        <symbol> + </symbol>
        <term>
          <integerConstant> 1 </integerConstant>
        </term>
      </expression>
      <symbol> ; </symbol>
    </letStatement>
  </statements>
  <symbol> } </symbol>
</whileStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.ws.genWhileStatement(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genWhileStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestDoStatement_genDoStatement(t *testing.T) {
	tests := []struct {
		name string
		do   *DoStatement
		want string
	}{
		{
			`test do Output.printString("THE AVERAGE IS: ");`,
			&DoStatement{
				Modi: "do",
				Sub: &SubroutineCall{
					Name: "Output",
					Dot:  ".",
					Sn:   "printString",
					LP:   "(",
					ExpL: []Expression{
						{
							Term: &StringConstant{
								V: "THE AVERAGE IS: ",
							},
						},
					},
					RP: ")",
				},
				Sc: ";",
			},
			`
<doStatement>
  <keyword> do </keyword>
  <identifier> Output </identifier>
  <symbol> . </symbol>
  <identifier> printString </identifier>
  <symbol> ( </symbol>
  <expressionList>
    <expression>
      <term>
        <stringConstant> THE AVERAGE IS:  </stringConstant>
      </term>
    </expression>
  </expressionList>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
</doStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.do.genDoStatement(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genDoStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestReturnStatement_genReturnStatement(t *testing.T) {
	tests := []struct {
		name string
		rs   *ReturnStatement
		want string
	}{
		{
			"test return;",
			&ReturnStatement{
				Modi: "return",
				Sc:   ";",
			},
			`
<returnStatement>
  <keyword> return </keyword>
  <symbol> ; </symbol>
</returnStatement>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.rs.genReturnStatement(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genReturnStatement() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestVarDec_genVarDec(t *testing.T) {
	tests := []struct {
		name string
		vd   *VarDec
		want string
	}{
		{
			"test var Array a;",
			&VarDec{
				Modi: "var",
				Vt:   identifier("Array"),
				Vn:   "a",
				Sc:   ";",
			},
			`
<varDec>
  <keyword> var </keyword>
  <identifier> Array </identifier>
  <identifier> a </identifier>
  <symbol> ; </symbol>
</varDec>
`,
		},
		{
			"test var int i, sum;",
			&VarDec{
				Modi: "var",
				Vt:   keyword("int"),
				Vn:   "i",
				Vns: []*NextVns{
					{
						Comma: ",",
						Vn:    "sum",
					},
				},
				Sc: ";",
			},
			`
<varDec>
  <keyword> var </keyword>
  <keyword> int </keyword>
  <identifier> i </identifier>
  <symbol> , </symbol>
  <identifier> sum </identifier>
  <symbol> ; </symbol>
</varDec>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.vd.genVarDec(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genVarDec() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestSubroutineBody_genSubroutineBody(t *testing.T) {
	tests := []struct {
		name string
		sb   *SubroutineBody
		want string
	}{
		{
			"test { var Array a; return; }",
			&SubroutineBody{
				LB: "{",
				Vd: []*VarDec{
					{
						Modi: "var",
						Vt:   identifier("Array"),
						Vn:   "a",
						Sc:   ";",
					},
				},
				Stmts: []Statement{
					&ReturnStatement{
						Modi: "return",
						Sc:   ";",
					},
				},
				RB: "}",
			},
			`
<subroutineBody>
  <symbol> { </symbol>
  <varDec>
    <keyword> var </keyword>
    <identifier> Array </identifier>
    <identifier> a </identifier>
    <symbol> ; </symbol>
  </varDec>
  <statements>
    <returnStatement>
      <keyword> return </keyword>
      <symbol> ; </symbol>
    </returnStatement>
  </statements>
  <symbol> } </symbol>
</subroutineBody>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.sb.genSubroutineBody(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genSubroutineBody() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}

func TestSubroutineDec_genSubroutineDec(t *testing.T) {
	tests := []struct {
		name string
		sd   *SubroutineDec
		want string
	}{
		{
			`test:
			function void main() {
				var SquareGame game;
				let game = game;
				do game.run();
				do game.dispose();
				return;
			}`,
			&SubroutineDec{
				Modi: "function",
				St:   "void",
				Sn:   "main",
				LP:   "(",
				RP:   ")",
				Sb: SubroutineBody{
					LB: "{",
					Vd: []*VarDec{
						{
							Modi: "var",
							Vt:   identifier("SquareGame"),
							Vn:   "game",
							Sc:   ";",
						},
					},
					Stmts: []Statement{
						&LetStatement{
							Modi: "let",
							Vn:   "game",
							Eq:   "=",
							Rexp: Expression{
								Term: &VarName{
									V: "game",
								},
							},
							Sc: ";",
						},
						&DoStatement{
							Modi: "do",
							Sub: &SubroutineCall{
								Name: "game",
								Dot:  ".",
								Sn:   "run",
								LP:   "(",
								RP:   ")",
							},
							Sc: ";",
						},
						&DoStatement{
							Modi: "do",
							Sub: &SubroutineCall{
								Name: "game",
								Dot:  ".",
								Sn:   "dispose",
								LP:   "(",
								RP:   ")",
							},
							Sc: ";",
						},
						&ReturnStatement{
							Modi: "return",
							Sc:   ";",
						},
					},
					RB: "}",
				},
			},
			`
<subroutineDec>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> main </identifier>
  <symbol> ( </symbol>
  <parameterList></parameterList>
  <symbol> ) </symbol>
  <subroutineBody>
    <symbol> { </symbol>
    <varDec>
      <keyword> var </keyword>
      <identifier> SquareGame </identifier>
      <identifier> game </identifier>
      <symbol> ; </symbol>
    </varDec>
    <statements>
      <letStatement>
        <keyword> let </keyword>
        <identifier> game </identifier>
        <symbol> = </symbol>
        <expression>
          <term>
            <identifier> game </identifier>
          </term>
        </expression>
        <symbol> ; </symbol>
      </letStatement>
      <doStatement>
        <keyword> do </keyword>
        <identifier> game </identifier>
        <symbol> . </symbol>
        <identifier> run </identifier>
        <symbol> ( </symbol>
        <expressionList></expressionList>
        <symbol> ) </symbol>
        <symbol> ; </symbol>
      </doStatement>
      <doStatement>
        <keyword> do </keyword>
        <identifier> game </identifier>
        <symbol> . </symbol>
        <identifier> dispose </identifier>
        <symbol> ( </symbol>
        <expressionList></expressionList>
        <symbol> ) </symbol>
        <symbol> ; </symbol>
      </doStatement>
      <returnStatement>
        <keyword> return </keyword>
        <symbol> ; </symbol>
      </returnStatement>
    </statements>
    <symbol> } </symbol>
  </subroutineBody>
</subroutineDec>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, b := initEncBytes()
			tt.sd.genSubroutineDec(e)
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(got, want) {
				t.Errorf("genSubroutineDec() = \n %v", got)
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}
