package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestVarInt(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []int32{0, -1, 1, 7389, 9936236, -12321455, 9892982, 2147483647, -2147483648}

	for _, val := range vals {
		if err := mctypes.WriteVarInt(conn, val); err != nil {
			t.Errorf("Error writting value to conn: %v\n", err)
		}

		n, err := mctypes.ReadVarInt(conn)
		if err != nil {
			t.Errorf("Error reading value from conn: %v\n", err)
		}

		if n != val {
			t.Errorf("Error, got %d, wanted %d\n", n, val)
		}
	}
}
