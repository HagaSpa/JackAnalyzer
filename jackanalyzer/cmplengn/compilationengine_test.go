package cmplengn

import (
	"bytes"
	"encoding/xml"
	"jackanalyzer/token"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		t token.Token
		s string
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
				s: "abcd",
			},
			&CompilationEngine{
				w: bytes.NewBufferString("abcd"),
				t: token.Token{
					TokenType:  token.IDENTIFIER,
					Identifier: "hoge",
				},
			},
			"abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := bytes.NewBufferString(tt.args.s)
			if got := New(w, tt.args.t); !reflect.DeepEqual(got.t, tt.want.t) {
				t.Errorf("New() = %#v, want %#v", got.t, tt.want.t)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("New() = %v, want %v", gotW, tt.wantW)
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
