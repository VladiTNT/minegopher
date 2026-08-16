package mcproto

import (
	"bytes"
	"encoding/binary"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

// Ping-Pong packets have a packet if of 1.
const PingPongID int32 = 1

// Encodes a Ping-Pong packet into buf. If you are sending this pong back to a client it is necessary that the
// long is the same as the one received in the client's ping request.
func EncodePingPong(buf *bytes.Buffer, n int64) error {
	if err := mctypes.WriteVarInt(buf, PingPongID); err != nil {
		return err
	}

	temp := make([]byte, 8)

	_, err := binary.Encode(temp, binary.BigEndian, n)
	if err != nil {
		return err
	}

	_, err = buf.Write(temp)
	return err
}

// Decodes a Ping-Pong packet payload into the long encoded within.
func DecodePingPong(payload []byte) (int64, error) {
	var n int64
	_, err := binary.Decode(payload, binary.BigEndian, &n)
	return n, err
}
