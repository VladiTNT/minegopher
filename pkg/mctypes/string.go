package mctypes

import (
	"errors"
	"io"
)

// Wrties p to w as a Minecraft compatible string
func WriteString(w io.Writer, p []byte) error {
	if err := WriteVarInt(w, int32(len(p))); err != nil {
		return err
	}

	n, err := w.Write(p)
	if err != nil {
		return err
	}

	if n != len(p) {
		return errors.New("mctypes: bytes written and string length don't match")
	}

	return nil
}

// Reads from the stream and returns the string encoded within.
func ReadString(r io.Reader) ([]byte, error) {
	strLen, err := ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, strLen)

	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	if n != int(strLen) {
		return nil, errors.New("mctypes: bytes read and string length don't match")
	}

	return buf, nil
}
