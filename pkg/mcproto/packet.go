package mcproto

import (
	"errors"
)

type PacketId int32

const (
	Status PacketId = iota
)

var (
	ErrInvalidPacketLength = errors.New("mcproto: packet length and payload size are not the same")
)

func GetPacket(r ReaderByteReader) (PacketId, []byte, error) {
	packetLength, _, err := ReadVarInt(r)
	if err != nil {
		return 0, nil, err
	}

	packetId, byteCount, err := ReadVarInt(r)
	if err != nil {
		return 0, nil, err
	}

	payloadLength := int(packetLength) - byteCount

	data := make([]byte, payloadLength)

	n, err := r.Read(data)
	if err != nil {
		return 0, nil, err
	}

	if n != payloadLength {
		return 0, nil, ErrInvalidPacketLength
	}

	return PacketId(packetId), data, nil
}
