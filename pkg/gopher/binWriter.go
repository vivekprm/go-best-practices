package gopher

import (
	"encoding/binary"
	"io"
)

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

// Write writes a value to the provided writer in little endian form
func (w *binWriter) Write(v interface{}) {
	// Checks if something before failed
	if w.err != nil {
		return
	}
	switch v.(type) {
	case string:
		s := v.(string)
		w.Write(int32(len(s)))
		w.Write([]byte(s))
	case int:
		i := v.(int)
		w.Write(int64(i))
	default:
		if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
			w.size += int64(binary.Size(v))
		}
	}
}
