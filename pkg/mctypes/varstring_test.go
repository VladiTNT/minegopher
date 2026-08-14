package mctypes_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

func TestVarString(t *testing.T) {
	conn := new(bytes.Buffer)
	vals := []string{"", "vlad", "yhgjui985 yhugj5468uhgj 453yhguji54", "🤪🤪🤪🤪"}

	for _, val := range vals {
		if err := mctypes.WriteVarString(conn, []byte(val)); err != nil {
			t.Errorf("Error writting value to conn: %v\n", err)
		}

		s, err := mctypes.ReadVarString(conn)
		if err != nil {
			t.Errorf("Error reading value from conn: %v\n", err)
		}

		if string(s) != val {
			t.Errorf("Error, got %s, wanted %s\n", s, val)
		}
	}
}
