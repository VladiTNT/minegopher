package mcproto

import (
	"bytes"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

// Ping-Pong packets have a packet if of 1.
const PingPongID int32 = 1

// Encodes a Ping-Pong packet into buf. If you are sending this pong back to a client it is ideal that the
// varlong is the same as the one received in the client's ping request.
func EncodePingPong(buf *bytes.Buffer, n int64) error {
	if err := mctypes.WriteVarInt(buf, PingPongID); err != nil {
		return err
	}

	return mctypes.WriteVarLong(buf, n)
}

// Decodes a Ping-Pong packet payload into the varlong encoded within.
func DecodePingPong(payload []byte) (int64, error) {
	return mctypes.ReadVarLong(bytes.NewReader(payload))
}
