package mcproto

import (
	"bytes"
	"errors"
	"io"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

// Reads a packet sent by a Minecraft client from the stream and returns the packet id and payload.
func ReadPacket(r io.Reader) (int32, []byte, error) {
	// Read packet length, each packet is prefixed with it's length
	packetLen, err := mctypes.ReadVarInt(r)
	if err != nil {
		return 0, nil, err
	}

	p := make([]byte, packetLen)

	// Read the packet body
	n, err := io.ReadFull(r, p)
	if err != nil {
		return 0, nil, err
	}

	if n != int(packetLen) {
		return 0, nil, errors.New("mcproto: bytes read and packet length don't match")
	}

	// Wrap it in bytes.Buffer
	buf := bytes.NewBuffer(p)

	// Read the first value of the packet, which is the packet id
	packetId, err := mctypes.ReadVarInt(buf)
	if err != nil {
		return 0, nil, err
	}

	return packetId, buf.Bytes(), nil
}

// Writes the contents of buf to w and prefixes it with the buffer length encoded as a varint, which is how
// Minecraft clients/servers expect to receive packets. The other functions in this package encode packets
// in bytes.Buffer objects, pair them with this function to flush into the network.
func WritePacket(w io.Writer, buf *bytes.Buffer) error {
	bufLen := int32(buf.Len())

	// Prefix packet with the packet length first
	if err := mctypes.WriteVarInt(w, bufLen); err != nil {
		return err
	}

	// Then write the packet contents
	n, err := buf.WriteTo(w)
	if err != nil {
		return err
	}

	if n != int64(bufLen) {
		return errors.New("mcproto: bytes written and buffer length don't match")
	}

	return nil
}
