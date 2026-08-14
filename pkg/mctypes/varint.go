package mctypes

import (
	"errors"
	"io"
)

// Writes n to w as a Minecraft compatible varint.
func WriteVarInt(w io.Writer, n int32) error {
	u := uint32(n)
	buf := make([]byte, 0, 5)

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

// Reads from the stream and returns the varint encoded within.
func ReadVarInt(r io.Reader) (int32, error) {
	var n int32
	var pos int
	buf := make([]byte, 1)

	for pos = 0; pos <= 32; pos += 7 {
		_, err := r.Read(buf)
		if err != nil {
			return 0, err
		}

		n |= int32(buf[0]&0x7F) << pos

		if pos >= 35 {
			return 0, errors.New("mctypes: varint too big")
		}

		if buf[0]&0x80 == 0 {
			break
		}
	}

	return n, nil
}
