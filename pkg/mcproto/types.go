package mcproto

import (
	"errors"
	"io"
)

type ReaderByteReader interface {
	io.Reader
	io.ByteReader
}

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

// ReadVarInt reads byte by byte from br and returns the encoded varint and the amount of bytes consumed.
func ReadVarInt(br io.ByteReader) (int32, int, error) {
	var val int32
	var pos int

	for pos = 0; pos < 32; pos += 7 {
		b, err := br.ReadByte()
		if err != nil {
			return 0, 0, err
		}

		val |= int32(b&0x7F) << pos

		if pos >= 35 {
			return 0, 0, ErrVarIntTooBig
		}

		if (b & 0x80) == 0 {
			break
		}
	}

	return val, pos/7 + 1, nil
}

func WriteString(w io.Writer, s string) error {
	data := []byte(s)

	if err := WriteVarInt(w, int32(len(data))); err != nil {
		return err
	}

	_, err := w.Write(data)

	return err
}

func ReadString(r ReaderByteReader) (string, error) {
	strLen, _, err := ReadVarInt(r)
	if err != nil {
		return "", err
	}

	data := make([]byte, strLen)

	_, err = r.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
