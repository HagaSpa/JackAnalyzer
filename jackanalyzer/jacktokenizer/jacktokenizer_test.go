package jacktokenizer

import (
	"bufio"
	"io"
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
				s: bufio.NewScanner(strings.NewReader("abcdefg")),
			},
		},
	}
	for _, tt := range tests {
		jt := New(tt.args.r)
		t.Run(tt.name, func(t *testing.T) {
			if jt.s.Text() != tt.want.s.Text() {
				t.Errorf("New() = %v, want %v", jt.s.Text(), tt.want.s.Text())
			}
		})
	}
}
