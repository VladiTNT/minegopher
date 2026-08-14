package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestVarLong(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []int64{0, 1, -1, 523231, -32663262, 235626622346, 9223372036854775807, -9223372036854775808}

	for _, val := range vals {
		if err := mctypes.WriteVarLong(conn, val); err != nil {
			t.Errorf("Error: %v\n", err)
		}

		n, err := mctypes.ReadVarLong(conn)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		if n != val {
			t.Errorf("Fuck, got %d, wanted %d\n", n, val)
		}
	}
}
