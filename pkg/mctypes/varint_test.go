package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestVarInt(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []int32{0, 1, -1, 24421, 5525323, -4354366, 2147483647, -2147483648}

	for _, val := range vals {
		if err := mctypes.WriteVarInt(conn, val); err != nil {
			t.Errorf("Error: %v\n", err)
		}

		n, err := mctypes.ReadVarInt(conn)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		if n != val {
			t.Errorf("Fuck, got %d, wanted %d\n", n, val)
		}
	}
}
