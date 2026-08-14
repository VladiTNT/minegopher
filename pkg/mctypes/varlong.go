package mctypes

import (
	"errors"
	"io"
)

// Writes n to w as a Minecraft compatible varlong.
func WriteVarLong(w io.Writer, n int64) error {
	u := uint64(n)
	buf := make([]byte, 0, 10)

	for {
		b := byte(u & 0x7F)
		u >>= 7

		if u != 0 {
			b |= 0x80
		}

		buf = append(buf, b)

		if u == 0 {
			break
		}
	}

	_, err := w.Write(buf)
	return err
}

// Reads from the stream and returns the varlong encoded within.
func ReadVarLong(r io.Reader) (int64, error) {
	var n int64
	var pos int
	buf := make([]byte, 1)

	for pos = 0; pos <= 64; pos += 7 {
		_, err := r.Read(buf)
		if err != nil {
			return 0, err
		}

		n |= int64(buf[0]&0x7F) << pos

		if pos >= 70 {
			return 0, errors.New("mctypes: varlong too big")
		}

		if buf[0]&0x80 == 0 {
			break
		}
	}

	return n, nil
}
