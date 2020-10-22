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
		want *Token
	}{
		{
			"test",
			&Token{
				next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jt := New(strings.NewReader(""))
			if got := jt.Tokenize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JackTokenizer.Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
