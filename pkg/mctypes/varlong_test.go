package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestVarLong(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []int64{0, 1, -1, 1234, 886554, -6548676, 9223372036854775807, -9223372036854775808}

	for _, val := range vals {
		if err := mctypes.WriteVarLong(conn, val); err != nil {
			t.Errorf("Error writtin value to conn: %v\n", err)
		}

		n, err := mctypes.ReadVarLong(conn)
		if err != nil {
			t.Errorf("Error reading value from conn: %v\n", err)
		}

		if n != val {
			t.Errorf("Error, got %d, wanted %d\n", n, val)
		}
	}
}
