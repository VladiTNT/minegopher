package mctypes

import (
	"errors"
	"io"
)

var (
	ErrVarIntTooBig = errors.New("mctypes: varint too big")
)

// WriteVarInt writes n byte by byte as a Minecraft compatible varint to w.
func WriteVarInt(w io.ByteWriter, n int32) error {
	u := uint32(n)
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

// ReadVarInt reads from r byte by byte and returns the varint encoded within.
func ReadVarInt(r io.ByteReader) (int32, error) {
	var n int32
	var pos int

	for pos = 0; pos < 32; pos += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		n |= int32(b&0x7F) << pos

		if pos >= 35 {
			return 0, ErrVarIntTooBig
		}

		if b&0x80 == 0 {
			break
		}
	}

	return n, nil
}
