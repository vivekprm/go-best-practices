package gopher

import (
	"io"
)

type Gopher struct {
	Name     string
	AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	bw := &binWriter{
		w: w,
	}
	bw.Write(g.Name)
	bw.Write(g.AgeYears)
	return bw.size, bw.err
}
