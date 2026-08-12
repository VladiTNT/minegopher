package mcproto_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
)

func TestVarInt(t *testing.T) {
	vals := []int32{0, -1, 22, 64435, -2556123, 434367373, 2023486731}
	var buf bytes.Buffer

	for _, val := range vals {
		if err := mcproto.WriteVarInt(&buf, val); err != nil {
			t.Errorf("couldn't write varint to buf: %v\n", err)
		}
	}

	for _, val := range vals {
		num, _, err := mcproto.ReadVarInt(&buf)
		if err != nil {
			t.Errorf("couldn't read varint from buf: %v\n", err)
		}

		if num != val {
			t.Errorf("mismatch, got %d, wanted %d\n", num, val)
		}
	}
}

func TestString(t *testing.T) {
	vals := []string{
		"nerd",
		"4gf5hty8ubj7i96yu9ij8h0o7k66th4gj34tgfy58hr",
		"osdijhfoisdoisdoijhfsd09231098u2398h23498h23",
		"VladiTNT",
		"😊😊😊😊😊😊😊😊😊😊😊😊😊",
	}
	var buf bytes.Buffer

	for _, val := range vals {
		if err := mcproto.WriteString(&buf, val); err != nil {
			t.Errorf("couldn't write string to buf: %v\n", err)
		}
	}

	for _, val := range vals {
		s, err := mcproto.ReadString(&buf)
		if err != nil {
			t.Errorf("couldn't read string to buf: %v\n", err)
		}

		if s != val {
			t.Errorf("mismatch, got %s, wanted %s\n", s, val)
		}
	}
}
