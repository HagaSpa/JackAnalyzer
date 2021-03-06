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
		{
			"test stringConstant",
			`"hoge"`,
			&token.Token{
				Next: &token.Token{
					TokenType: token.STRING_CONST,
					StringVal: "hoge",
				},
			},
		},
		{
			"test comment",
			"// hoge",
			&token.Token{
				Next: nil,
			},
		},
		{
			"test comment asterisk",
			`/* hoge
aaaaaaaa bbbbbbb ccccccc
			*/`,
			&token.Token{
				Next: nil,
			},
		},
		{
			"test jack code",
			`
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/10/ArrayTest/Main.jack

// (identical to projects/09/Average/Main.jack)

/** Computes the average of a sequence of integers. */

class Main {
	function void main() {
		var Array a;
	}
}
			`,
			&token.Token{
				Next: &token.Token{
					Next: &token.Token{
						Next: &token.Token{
							Next: &token.Token{
								Next: &token.Token{
									Next: &token.Token{
										Next: &token.Token{
											Next: &token.Token{
												Next: &token.Token{
													Next: &token.Token{
														Next: &token.Token{
															Next: &token.Token{
																Next: &token.Token{
																	Next: &token.Token{
																		Next: &token.Token{
																			TokenType: token.SYMBOL,
																			Symbol:    "}",
																		},
																		TokenType: token.SYMBOL,
																		Symbol:    "}",
																	},
																	TokenType: token.SYMBOL,
																	Symbol:    ";",
																},
																TokenType:  token.IDENTIFIER,
																Identifier: "a",
															},
															TokenType:  token.IDENTIFIER,
															Identifier: "Array",
														},
														TokenType: token.KEYWORD,
														Keyword:   token.VAR,
													},
													TokenType: token.SYMBOL,
													Symbol:    "{",
												},
												TokenType: token.SYMBOL,
												Symbol:    ")",
											},
											TokenType: token.SYMBOL,
											Symbol:    "(",
										},
										TokenType:  token.IDENTIFIER,
										Identifier: "main",
									},
									TokenType: token.KEYWORD,
									Keyword:   token.VOID,
								},
								TokenType: token.KEYWORD,
								Keyword:   token.FUNCTION,
							},
							TokenType: token.SYMBOL,
							Symbol:    "{",
						},
						TokenType:  token.IDENTIFIER,
						Identifier: "Main",
					},
					TokenType: token.KEYWORD,
					Keyword:   token.CLASS,
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

func Test_isAlpherUnder(t *testing.T) {
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
			if got := isAlpherUnder(tt.r); got != tt.want {
				t.Errorf("Case: %s isAlUn() = %t, want %t", tt.name, got, tt.want)
			}
		})
	}
}

func TestTokenizer_startsWithStringConstant(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"test",
			`test"`, // Suppose you are getting the first double quate with Tokenize().
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.s))
			if got := tz.startsWithStringConstant(); got != tt.want {
				t.Errorf("Tokenizer.startsWithStringConstant() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_isDoubleQuotes(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			"double quotes",
			'"',
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDoubleQuotes(tt.r); got != tt.want {
				t.Errorf("isDoubleQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizer_isComment(t *testing.T) {
	type args struct {
		r rune
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
		s     string
	}{
		{
			"comment",
			args{
				r: '/',
				s: "/",
			},
			true,
			token.COMMENT,
			"",
		},
		{
			"comment_ast",
			args{
				r: '/',
				s: "*",
			},
			true,
			token.COMMENT_AST,
			"",
		},
		{
			"r is not /",
			args{
				r: 'a',
				s: "abcd",
			},
			false,
			"",
			"abcd",
		},
		{
			"sigle slash",
			args{
				r: '/',
				s: "abcd",
			},
			false,
			"",
			"abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.args.s))
			got, got1 := tz.isComment(tt.args.r)
			if got != tt.want {
				t.Errorf("Tokenizer.isComment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Tokenizer.isComment() got1 = %v, want %v", got1, tt.want1)
			}
			l, _, _ := tz.re.ReadLine()
			if string(l) != tt.s {
				t.Errorf("Tokenizer.isComment() tz.re.ReadLine = %s, want %s", l, tt.s)
			}
		})
	}
}

func TestTokenizer_skipComment(t *testing.T) {
	type args struct {
		ct string
		s  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"comment",
			args{
				ct: token.COMMENT,
				s: `// comment
abc`,
			},
			"abc",
		},
		{
			"comment asterisk",
			args{
				ct: token.COMMENT_AST,
				s:  `/* comment cocococo */ abc`,
			},
			" abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := New(strings.NewReader(tt.args.s))
			tz.skipComment(tt.args.ct)
			l, _, _ := tz.re.ReadLine()
			if string(l) != tt.want {
				t.Errorf("Tokenizer.skipComment() tz.re.ReadLine = %s, want %s", l, tt.want)
			}
		})
	}
}
