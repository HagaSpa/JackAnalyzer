package cmplengn

import (
	"bytes"
	"encoding/xml"
	"jackanalyzer/token"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		t token.Token
		e *xml.Encoder
	}
	tests := []struct {
		name  string
		args  args
		want  *CompilationEngine
		wantW string
	}{
		{
			"test",
			args{
				t: token.Token{
					TokenType:  token.IDENTIFIER,
					Identifier: "hoge",
				},
				e: xml.NewEncoder(bytes.NewBufferString("abcd")),
			},
			&CompilationEngine{
				t: token.Token{
					TokenType:  token.IDENTIFIER,
					Identifier: "hoge",
				},
				e: xml.NewEncoder(bytes.NewBufferString("abcd")),
			},
			"abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.t, tt.args.e); !reflect.DeepEqual(got.t, tt.want.t) {
				t.Errorf("New() = %#v, want %#v", got.t, tt.want.t)
			}
		})
	}
}

func Test_genElement(t *testing.T) {
	tests := []struct {
		name  string
		t     token.Token
		want  string
		want1 xml.StartElement
	}{
		{
			"test KEYWORD",
			token.Token{
				TokenType: token.KEYWORD,
				Keyword:   token.CLASS,
			},
			" class ",
			xml.StartElement{Name: xml.Name{Local: "keyword"}},
		},
		{
			"test IDENTIFIER",
			token.Token{
				TokenType:  token.IDENTIFIER,
				Identifier: "hoge",
			},
			" hoge ",
			xml.StartElement{Name: xml.Name{Local: "identifier"}},
		},
		{
			"test SYMBOL",
			token.Token{
				TokenType: token.SYMBOL,
				Symbol:    ",",
			},
			" , ",
			xml.StartElement{Name: xml.Name{Local: "symbol"}},
		},
		{
			"test INT_CONST",
			token.Token{
				TokenType: token.INT_CONST,
				IntVal:    123,
			},
			" 123 ",
			xml.StartElement{Name: xml.Name{Local: "integerConstant"}},
		},
		{
			"test STRING_CONST",
			token.Token{
				TokenType: token.STRING_CONST,
				StringVal: "THE AVERAGE IS:",
			},
			" THE AVERAGE IS: ",
			xml.StartElement{Name: xml.Name{Local: "stringConstant"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := genElement(tt.t)
			if got != tt.want {
				t.Errorf("genElement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("genElement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCompilationEngine_compileExpression(t *testing.T) {
	tests := []struct {
		name    string
		t       token.Token
		want    string
		wantErr error
	}{
		{
			"test (op term)* is not exist.",
			token.Token{
				TokenType: token.KEYWORD,
				Keyword:   "class",
			},
			`
<expression></expression>
`,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			e := xml.NewEncoder(&b)
			e.Indent("", "  ")
			ce := New(tt.t, e)
			wantErr := ce.compileExpression()
			e.Flush()
			got := b.String()
			want := strings.TrimRight(strings.TrimLeft(tt.want, "\n"), "\n")

			if !reflect.DeepEqual(got, want) {
				t.Errorf("ce.compileExpression() = %v", got)
				t.Errorf("wantXml = %v", want)
			}
			if wantErr != tt.wantErr {
				t.Errorf("ce.compileExpression() error got = %v, want = %v ", wantErr, tt.wantErr)
			}
		})
	}
}
