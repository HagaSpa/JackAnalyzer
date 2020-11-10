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
			" ",
			&token.Token{
				Next: nil,
			},
		},
		{
			"test class",
			"class",
			&token.Token{
				Next: &token.Token{
					TokenType: token.KEYWORD,
					Keyword:   token.CLASS,
				},
			},
		},
		{
			"test for symbol. '{'",
			"{",
			&token.Token{
				Next: &token.Token{
					TokenType: token.SYMBOL,
					Symbol:    "{",
				},
			},
		},
		{
			"test for while{}",
			"while {}",
			&token.Token{
				Next: &token.Token{
					Next: &token.Token{
						Next: &token.Token{
							TokenType: token.SYMBOL,
							Symbol:    "}",
						},
						TokenType: token.SYMBOL,
						Symbol:    "{",
					},
					TokenType: token.KEYWORD,
					Keyword:   token.WHILE,
				},
			},
		},
		{
			"test identifier",
			"hoge",
			&token.Token{
				Next: &token.Token{
					TokenType:  token.IDENTIFIER,
					Identifier: "hoge",
				},
			},
		},
		{
			"test integerConstant",
			"1234",
			&token.Token{
				Next: &token.Token{
					TokenType: token.INT_CONST,
					IntVal:    1234,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.s))
			if got := tz.Tokenize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JackTokenizer.Tokenize() = %#v, want %#v", got, tt.want)
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

func TestTokenizer_newToken(t *testing.T) {
	type args struct {
		cur *token.Token
		tt  token.TokenType
		kw  token.Keyword
		sb  string
		id  string
		iv  int
		sv  string
	}
	tests := []struct {
		name string
		args args
		want *token.Token
	}{
		{
			"keyword (CLASS)",
			args{
				cur: &token.Token{},
				tt:  token.KEYWORD,
				kw:  token.CLASS,
				// Does not use for KEYWORD
				sb: "",
				id: "",
				iv: 0,
				sv: "",
			},
			&token.Token{
				TokenType: token.KEYWORD,
				Keyword:   token.CLASS,
			},
		},
		{
			"keyword (METHOD)",
			args{
				cur: &token.Token{},
				tt:  token.KEYWORD,
				kw:  token.METHOD,
				// Does not use for KEYWORD
				sb: "",
				id: "",
				iv: 0,
				sv: "",
			},
			&token.Token{
				TokenType: token.KEYWORD,
				Keyword:   token.METHOD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newToken(tt.args.cur, tt.args.tt, tt.args.kw, tt.args.sb, tt.args.id, tt.args.iv, tt.args.sv)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenizer.newToken() = %v, want %v", got, tt.want)
			}
			// got == cur.Next?
			if !reflect.DeepEqual(got, tt.args.cur.Next) {
				t.Errorf("Tokenizer.newToken() = %v, want %v", got, tt.args.cur.Next)
			}
		})
	}
}

func TestTokenizer_startsWithIdentifier(t *testing.T) {
	type args struct {
		r rune
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				r: 't',
				s: "est",
			},
			"test",
		},
		{
			"test alphanumeric",
			args{
				r: 'h',
				s: "oge1",
			},
			"hoge1",
		},
		{
			"test contains white space",
			args{
				r: 'h',
				s: "oge1 hoge2",
			},
			"hoge1",
		},
		{
			"exclude japanese",
			args{
				r: 'h',
				s: "Aあ",
			},
			"hA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.args.s))
			if got := tz.startsWithIdentifier(tt.args.r); got != tt.want {
				t.Errorf("Tokenizer.startsWithIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizer_startsWithIntegerConstant(t *testing.T) {
	type args struct {
		r rune
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				r: '1',
				s: "23",
			},
			123,
		},
		{
			"test 012",
			args{
				r: '0',
				s: "12",
			},
			12,
		},
		{
			"test 000",
			args{
				r: '0',
				s: "00",
			},
			0,
		},
		{
			"test 101",
			args{
				r: '1',
				s: "01",
			},
			101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.args.s))
			if got := tz.startsWithIntegerConstant(tt.args.r); got != tt.want {
				t.Errorf("Tokenizer.startsWithIntegerConstant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlUn(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			"test lower alpha",
			'a',
			true,
		},
		{
			"test upper alpha",
			'A',
			true,
		},
		{
			"test under score",
			'_',
			true,
		},
		{
			"test japanese",
			'あ',
			false,
		},
		{
			"test numeric",
			'0',
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlUn(tt.r); got != tt.want {
				t.Errorf("Case: %s isAlUn() = %t, want %t", tt.name, got, tt.want)
			}
		})
	}
}
