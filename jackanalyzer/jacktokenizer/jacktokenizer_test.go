package jacktokenizer

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *JackTokenizer
	}{
		{
			"test",
			args{
				r: strings.NewReader("abcdefg"),
			},
			&JackTokenizer{
				re: bufio.NewReader(strings.NewReader("abcdefg")),
			},
		},
	}
	for _, tt := range tests {
		jt := New(tt.args.r)
		got, _ := jt.re.ReadString('\n')
		want, _ := tt.want.re.ReadString('\n')
		t.Run(tt.name, func(t *testing.T) {
			if got != want {
				t.Errorf("New() = %v, want %v", got, want)
			}
		})
	}
}

func TestJackTokenizer_Tokenize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want *Token
	}{
		{
			"test white space",
			"wh ile",
			&Token{
				next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := New(strings.NewReader(tt.s))
			if got := jt.Tokenize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JackTokenizer.Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJackTokenizer_startsWithKeyword(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Keyword
	}{
		{
			"test",
			"class",
			CLASS,
		},
		{
			"method",
			"method",
			METHOD,
		},
		{
			"function",
			"function",
			FUNCTION,
		},
		{
			"constructor",
			"constructor",
			CONSTRUCTOR,
		},
		{
			"int",
			"int",
			INT,
		},
		{
			"boolean",
			"boolean",
			BOOLEAN,
		},
		{
			"char",
			"char",
			CHAR,
		},
		{
			"void",
			"void",
			VOID,
		},
		{
			"var",
			"var",
			VAR,
		},
		{
			"static",
			"static",
			STATIC,
		},
		{
			"field",
			"field",
			FIELD,
		},
		{
			"let",
			"let",
			LET,
		},
		{
			"do",
			"do",
			DO,
		},
		{
			"if",
			"if",
			IF,
		},
		{
			"else",
			"else",
			ELSE,
		},
		{
			"while",
			"while",
			WHILE,
		},
		{
			"return",
			"return",
			RETURN,
		},
		{
			"true",
			"true",
			TRUE,
		},
		{
			"false",
			"false",
			FALSE,
		},
		{
			"null",
			"null",
			NULL,
		},
		{
			"this",
			"this",
			THIS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := New(strings.NewReader(tt.s))
			if got := jt.startsWithKeyword(); got != tt.want {
				t.Errorf("JackTokenizer.startsWithKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
