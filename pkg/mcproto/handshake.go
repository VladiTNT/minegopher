package mcproto

import (
	"bytes"
	"encoding/binary"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

// Handshake is what a client (1.7+) sends as it's first message to a minecraft server
type Handshake struct {
	Protocol int32
	Address  []byte
	Port     uint16
	Next     int32
}

// Decodes a packet payload into the handshake data enccoded within.
func DecodeHandshakePacket(packetPayload []byte) (*Handshake, error) {
	r := bytes.NewReader(packetPayload)
	var err error
	var h Handshake

	h.Protocol, err = mctypes.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	h.Address, err = mctypes.ReadString(r)
	if err != nil {
		return nil, err
	}

	shortBuf := make([]byte, 2)
	_, err = r.Read(shortBuf)
	if err != nil {
		return nil, err
	}

	if _, err := binary.Decode(shortBuf, binary.BigEndian, &h.Port); err != nil {
		return nil, err
	}

	h.Next, err = mctypes.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
