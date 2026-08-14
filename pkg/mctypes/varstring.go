package mctypes

// WriteVarString writes a Minecraft compatible string to w.
func WriteVarString(w dynamicWriter, data []byte) error {
	if err := WriteVarInt(w, int32(len(data))); err != nil {
		return err
	}

	_, err := w.Write(data)
	return err
}

// ReadVarString reads from r and returns the string encoded within according to Minecraft's protocol.
func ReadVarString(r dynamicReader) ([]byte, error) {
	n, err := ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, n)

	_, err = r.Read(buf)
	return buf, err
}
