package mcproto_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
)

func TestVarInt(t *testing.T) {
	testValues := []int32{12, 32, 55445, 546334635, 7745, -1231, -12314354}
	var buf bytes.Buffer

	// Write test values
	for _, val := range testValues {
		if err := mcproto.WriteVarInt(&buf, val); err != nil {
			t.Errorf("TestVarInt: couldn't write varint to buf: %v\n", err)
		}
	}

	// Read test values
	for _, val := range testValues {
		n, err := mcproto.ReadVarInt(&buf)
		if err != nil {
			t.Errorf("TestVarInt: couldn't read varint from buf: %v\n", err)
		}

		if n != val {
			t.Errorf("TestVarInt: mismatch, got %d, needed %d\n", n, val)
		}
	}
}
