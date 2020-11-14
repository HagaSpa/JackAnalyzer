package token

import "testing"

func TestIsSymbol(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			"{",
			'{',
			true,
		},
		{
			"}",
			'}',
			true,
		},
		{
			"(",
			'(',
			true,
		},
		{
			")",
			')',
			true,
		},
		{
			"[",
			'[',
			true,
		},
		{
			"]",
			']',
			true,
		},
		{
			".",
			'.',
			true,
		},
		{
			",",
			',',
			true,
		},
		{
			";",
			';',
			true,
		},
		{
			"+",
			'+',
			true,
		},
		{
			"-",
			'-',
			true,
		},
		{
			"*",
			'*',
			true,
		},
		{
			"/",
			'/',
			true,
		},
		{
			"&",
			'&',
			true,
		},
		{
			"|",
			'|',
			true,
		},
		{
			"<",
			'<',
			true,
		},
		{
			">",
			'>',
			true,
		},
		{
			"=",
			'=',
			true,
		},
		{
			"~",
			'~',
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymbol(tt.r); got != tt.want {
				t.Errorf("IsSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken_HasMoreTokens(t *testing.T) {
	type fields struct {
		Next       *Token
		TokenType  TokenType
		Keyword    Keyword
		Symbol     string
		Identifier string
		IntVal     int
		StringVal  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"true",
			fields{
				Next:       &Token{},
				TokenType:  KEYWORD,
				Keyword:    CLASS,
				Symbol:     "",
				Identifier: "",
				IntVal:     0,
				StringVal:  "",
			},
			true,
		},
		{
			"false",
			fields{
				Next: nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := &Token{
				Next:       tt.fields.Next,
				TokenType:  tt.fields.TokenType,
				Keyword:    tt.fields.Keyword,
				Symbol:     tt.fields.Symbol,
				Identifier: tt.fields.Identifier,
				IntVal:     tt.fields.IntVal,
				StringVal:  tt.fields.StringVal,
			}
			if got := act.HasMoreTokens(); got != tt.want {
				t.Errorf("Token.HasMoreTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
