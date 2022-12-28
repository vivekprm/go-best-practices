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
	switch x := v.(type) {
	case string:
		w.Write(int32(len(x)))
		w.Write([]byte(x))
	case int:
		w.Write(int64(x))
	default:
		if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
			w.size += int64(binary.Size(v))
		}
	}
}
