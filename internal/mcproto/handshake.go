package mcproto

import (
	"bytes"
	"encoding/binary"
)

// A HandShake is the first thing a Minecraft 1.7+ client sends to a server when reloading the server menu.
type HandShake struct {
	ProtocolVersion int32
	ServerAddr      string
	ServerPort      uint16
	NextState       int32
}

// Get parses the given packet payload into the handshake.
func (h *HandShake) Get(data []byte) error {
	var err error

	br := bytes.NewReader(data)

	h.ProtocolVersion, _, err = ReadVarInt(br)
	if err != nil {
		return err
	}

	h.ServerAddr, err = ReadString(br)
	if err != nil {
		return err
	}

	shortBuf := make([]byte, 2)
	_, err = br.Read(shortBuf)
	if err != nil {
		return err
	}

	if _, err := binary.Decode(shortBuf, binary.BigEndian, &h.ServerPort); err != nil {
		return err
	}

	h.NextState, _, err = ReadVarInt(br)
	if err != nil {
		return err
	}

	return nil
}
