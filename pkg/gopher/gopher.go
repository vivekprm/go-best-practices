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
	bw.Write(int32(len(g.Name)))
	bw.Write([]byte(g.Name))
	bw.Write(int64(g.AgeYears))
	return bw.size, bw.err
}
