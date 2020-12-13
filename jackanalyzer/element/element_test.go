package element

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func Test_genCon(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		want string
	}{
		{
			"test keyword",
			keyword("class"),
			" class ",
		},
		{
			"test identifier",
			identifier("hoge"),
			" hoge ",
		},
		{
			"test symbol",
			symbol(","),
			" , ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genCon(tt.s); got != tt.want {
				t.Errorf("genCon() = %v, want %v", got, tt.want)
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
		cd   ClassVarDec
		want string
	}{
		{
			"test",
			ClassVarDec{
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
			var b bytes.Buffer
			e := xml.NewEncoder(&b)
			e.Indent("", "  ")
			// execute
			tt.cd.genClassVarDec(e)
			e.Flush()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(b.String(), want) {
				t.Errorf("cd.genClassVarDec() = %v", b.String())
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			e := xml.NewEncoder(&b)
			e.Indent("", "  ")
			// execute
			tt.pl.genParameterList(e)
			e.Flush()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")
			if !reflect.DeepEqual(b.String(), want) {
				t.Errorf("genParameterList() = \n %v", b.String())
				t.Errorf("wantXml = \n %v", want)
			}
		})
	}
}
