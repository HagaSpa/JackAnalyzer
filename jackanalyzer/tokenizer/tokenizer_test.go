package tokenizer

import (
	"bufio"
	"io"
	"jackanalyzer/token"
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
		want *Tokenizer
	}{
		{
			"test",
			args{
				r: strings.NewReader("abcdefg"),
			},
			&Tokenizer{
				re: bufio.NewReader(strings.NewReader("abcdefg")),
			},
		},
	}
	for _, tt := range tests {
		tz := New(tt.args.r)
		got, _ := tz.re.ReadString('\n')
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
		want *token.Token
	}{
		{
			"test white space",
			"wh ile",
			&token.Token{
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.s))
			if got := tz.Tokenize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JackTokenizer.Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJackTokenizer_startsWithKeyword(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want token.Keyword
	}{
		{
			"test",
			"class",
			token.CLASS,
		},
		{
			"method",
			"method",
			token.METHOD,
		},
		{
			"function",
			"function",
			token.FUNCTION,
		},
		{
			"constructor",
			"constructor",
			token.CONSTRUCTOR,
		},
		{
			"int",
			"int",
			token.INT,
		},
		{
			"boolean",
			"boolean",
			token.BOOLEAN,
		},
		{
			"char",
			"char",
			token.CHAR,
		},
		{
			"void",
			"void",
			token.VOID,
		},
		{
			"var",
			"var",
			token.VAR,
		},
		{
			"static",
			"static",
			token.STATIC,
		},
		{
			"field",
			"field",
			token.FIELD,
		},
		{
			"let",
			"let",
			token.LET,
		},
		{
			"do",
			"do",
			token.DO,
		},
		{
			"if",
			"if",
			token.IF,
		},
		{
			"else",
			"else",
			token.ELSE,
		},
		{
			"while",
			"while",
			token.WHILE,
		},
		{
			"return",
			"return",
			token.RETURN,
		},
		{
			"true",
			"true",
			token.TRUE,
		},
		{
			"false",
			"false",
			token.FALSE,
		},
		{
			"null",
			"null",
			token.NULL,
		},
		{
			"this",
			"this",
			token.THIS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.s))
			if got := tz.startsWithKeyword(); got != tt.want {
				t.Errorf("JackTokenizer.startsWithKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
