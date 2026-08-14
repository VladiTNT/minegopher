package mctypes

import "github.com/VladiTNT/minegopher/pkg/utils"

// WriteString writes a Minecraft compatible string to w.
func WriteString(w utils.DynamicWriter, data []byte) error {
	if err := WriteVarInt(w, int32(len(data))); err != nil {
		return err
	}

	_, err := w.Write(data)
	return err
}

// ReadString reads from r and returns the string encoded within according to Minecraft's protocol.
func ReadString(r utils.DynamicReader) ([]byte, error) {
	n, err := ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, n)

	_, err = r.Read(buf)
	return buf, err
}
