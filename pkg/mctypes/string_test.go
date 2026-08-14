package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestString(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []string{"", "Vlad", "Nerd0x1Gaming", "🥀🥀🥀🥀🥀🥀🥀🥀🥀"}

	for _, val := range vals {
		if err := mctypes.WriteString(conn, []byte(val)); err != nil {
			t.Errorf("Error: %v\n", err)
		}

		p, err := mctypes.ReadString(conn)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}

		if string(p) != val {
			t.Errorf("Fuck, got %s, wanted %s\n", p, val)
		}
	}
}
