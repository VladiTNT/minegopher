package mcproto

import (
	"bytes"
	"encoding/json"
)

const StatusPacketID int32 = 0

type ServerStatus struct {
	Version            StatusVersion     `json:"version"`
	Players            StatusPlayers     `json:"players"`
	Description        StatusDescription `json:"description"`
	Favicon            string            `json:"favicon"`
	EnforcesSecureChat bool              `json:"enforcesSecureChat"`
}

type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type StatusPlayers struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []StatusPlayer `json:"sample"`
}

type StatusPlayer struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type StatusDescription struct {
	Text string `json:"text"`
}

func PrepareStatusPacketPayload(ss *ServerStatus) ([]byte, error) {
	jsonData, err := json.Marshal(ss)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	if err := WriteString(&buf, string(jsonData)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
