package cmplengn

import (
	"io"
	"jackanalyzer/token"
)

type CompilationEngine struct {
	w io.Writer
	t token.Token
}

func New(w io.Writer, t token.Token) *CompilationEngine {
	ce := &CompilationEngine{
		w: w,
		t: t,
	}
	return ce
}
