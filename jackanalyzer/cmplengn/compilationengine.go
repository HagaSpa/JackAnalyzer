package cmplengn

import (
	"encoding/xml"
	"io"
	"jackanalyzer/token"
)

type CompilationEngine struct {
	w io.Writer
	t token.Token
	e *xml.Encoder
}

func New(w io.Writer, t token.Token) *CompilationEngine {
	e := xml.NewEncoder(w)
	ce := &CompilationEngine{
		w: w,
		t: t,
		e: e,
	}
	return ce
}
