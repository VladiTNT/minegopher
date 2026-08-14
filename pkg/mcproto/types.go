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

// WriteVarInt writes the given value to w as a Minecraft compatible VarInt.
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

// ReadVarInt parses the upcoming bytes and retruns a varint, it also returns the amount of bytes it consumed
// because Minecraft's protocol for whatever reason starts each packet with the packet size not the packet id.
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

		if b&0x80 == 0 {
			break
		}
	}

	return val, pos/7 + 1, nil
}

// WriteString writes a Minecraft compatible string to w.
func WriteString(w io.Writer, s string) error {
	data := []byte(s)

	if err := WriteVarInt(w, int32(len(data))); err != nil {
		return err
	}

	_, err := w.Write(data)

	return err
}

// ReadString reads a string encoded with Minecraft's protocol which uses a varint to encode the length and then
// the string data.
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

func ReadVarLong(r io.ByteReader) (int64, int, error) {
	var val int64
	var pos int

	for pos = 0; pos <= 64; pos += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, 0, err
		}

		val |= int64(b&0x7F) << pos

		if pos >= 70 {
			return 0, 0, ErrVarIntTooBig
		}

		if b&0x80 == 0 {
			break
		}
	}

	return val, pos/7 + 1, nil
}

func WriteVarLong(w io.Writer, val int64) error {
	uval := uint64(val)
	buf := make([]byte, 0, 10)

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
