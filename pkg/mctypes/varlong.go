package mctypes

import (
	"errors"
	"io"
)

var (
	ErrVarLongTooBig = errors.New("mctypes: varlong too big")
)

// WriteVarLong writes n byte by byte as a Minecraft compatible varlong to w.
func WriteVarLong(w io.ByteWriter, n int64) error {
	u := uint64(n)
	for {
		b := byte(u & 0x7F)
		u >>= 7
		if u != 0 {
			b |= 0x80
		}
		if err := w.WriteByte(b); err != nil {
			return err
		}
		if u == 0 {
			break
		}
	}
	return nil
}

// ReadVarLong reads from r byte by byte and returns the varlong encoded within.
func ReadVarLong(r io.ByteReader) (int64, error) {
	var n int64
	var pos int

	for pos = 0; pos < 64; pos += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		n |= int64(b&0x7F) << pos

		if pos >= 70 {
			return 0, ErrVarLongTooBig
		}

		if b&0x80 == 0 {
			break
		}
	}

	return n, nil
}
