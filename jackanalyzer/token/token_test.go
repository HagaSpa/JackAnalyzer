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
