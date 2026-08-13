package mcproto

import (
	"bytes"
	"errors"
	"io"
)

var (
	ErrInvalidPacketLength = errors.New("mcproto: packet length and payload size are not the same")
)

// GetPacket parses one Minecraft protocol packet, returning the packet id, payload and any errors.
func GetPacket(r ReaderByteReader) (int32, []byte, error) {
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

	return packetId, data, nil
}

func WritePacket(w io.Writer, packetId int32, payload []byte) error {
	var buf bytes.Buffer

	if err := WriteVarInt(&buf, packetId); err != nil {
		return err
	}

	if _, err := buf.Write(payload); err != nil {
		return err
	}

	if err := WriteVarInt(w, int32(buf.Len())); err != nil {
		return err
	}

	_, err := buf.WriteTo(w)
	return err
}
