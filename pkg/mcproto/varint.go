package mcproto

import (
	"errors"
	"io"
)

var (
	ErrVarIntTooBig = errors.New("mcproto: varint too big")
)

// WriteVarInt writes the given value to w as a minecraft compatible VarInt.
func WriteVarInt(w io.Writer, val int32) error {
	uval := uint32(val)
	buf := make([]byte, 0, 5)

	for {
		b := byte(uval & 0x7F)
		uval >>= 7
		if uval != 0 {
			b |= 0x80
		}
		buf = append(buf, b)
		if uval == 0 {
			break
		}
	}

	_, err := w.Write(buf)
	return err
}

// ReadVarInt reads from br and returns a minecraft compatible VarInt. Ideal to use a bufferd reader because this
// function reads byte by byte and if each read requires a system call that could severely impact performance.
func ReadVarInt(br io.ByteReader) (int32, error) {
	var val int32
	for pos := 0; pos < 32; pos += 7 {
		b, err := br.ReadByte()
		if err != nil {
			return 0, err
		}

		val |= int32(b&0x7F) << pos

		if pos >= 35 {
			return 0, ErrVarIntTooBig
		}

		if (b & 0x80) == 0 {
			break
		}
	}
	return val, nil
}
