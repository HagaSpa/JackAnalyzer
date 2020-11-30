package element

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func Test_genContent(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		want string
	}{
		{
			"string",
			"test",
			" test ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genContent(tt.s); got != tt.want {
				t.Errorf("genContent() = %v, want %v", got, tt.want)
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
				Cvds: []ClassVarDec{
					{
						Modifier:  "field",
						VarType:   "int",
						VarNames:  []VarName{"x", "y"},
						SemiColon: ";",
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
				Cvds: []ClassVarDec{
					{
						Modifier:  "field",
						VarType:   "int",
						VarNames:  []VarName{"x", "y"},
						SemiColon: ";",
					},
					{
						Modifier:  "field",
						VarType:   "int",
						VarNames:  []VarName{"size"},
						SemiColon: ";",
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
				Modifier:  "field",
				VarType:   "int",
				VarNames:  []VarName{"x", "y"},
				SemiColon: ";",
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
