package mcproto

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/VladiTNT/minegopher/pkg/mctypes"
)

// Server status packets use a packet id of 0.
const StatusID int32 = 0

// This is the JSON structure that a Minecraft client expects from a server when making a status request.
// Some fields are optional.
type ServerStatus struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Description struct {
		Text string `json:"text"`
	} `json:"description"`
	Favicon            string `json:"favicon"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat"`
}

// Prepares a status request packet in buf, send to a minecraft server to get server status info.
func EncodeStatusRequest(buf *bytes.Buffer) error {
	return mctypes.WriteVarInt(buf, StatusID)
}

// Prepares a status response packet in buf.
func EncodeStatusResponse(buf *bytes.Buffer, ss *ServerStatus) error {
	if err := mctypes.WriteVarInt(buf, StatusID); err != nil {
		return err
	}

	return json.NewEncoder(buf).Encode(ss)
}

// Encodes the given png into a base64 encoded string with the header that Minecraft clients expect, meant to
// be used in the Favicon field of the ServerStatus struct.
func EncodeFavicon(pngData []byte) string {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.RawStdEncoding, &buf)
	defer enc.Close()
	if _, err := enc.Write(pngData); err != nil {
		return ""
	}
	return fmt.Sprintf("data:image/png;base64,%s", buf.Bytes())
}
