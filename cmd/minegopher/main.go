package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"net"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
)

//go:embed VladPixelFace-64.png
var EmbeddedPng []byte

func main() {
	ln, err := net.Listen("tcp", ":25565")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rd := bufio.NewReader(conn)

	_, handShakeData, err := mcproto.GetPacket(rd)
	if err != nil {
		panic(err)
	}

	var handShake mcproto.HandShake
	if err := handShake.Get(handShakeData); err != nil {
		panic(err)
	}

	fmt.Printf("Proto: %d\nAddr: %s\nPort: %d\nNext: %d\n",
		handShake.ProtocolVersion,
		handShake.ServerAddr,
		handShake.ServerPort,
		handShake.NextState,
	)

	packetId, _, err := mcproto.GetPacket(rd)
	if err != nil {
		panic(err)
	}

	if packetId == mcproto.StatusPacketID {
		fmt.Println("Received status ping.")

		ss := mcproto.ServerStatus{
			Version: mcproto.StatusVersion{
				Name:     "26.2",
				Protocol: 776,
			},
			Players: mcproto.StatusPlayers{
				Max:    999999999,
				Online: 80085,
				Sample: []mcproto.StatusPlayer{
					{Name: "nerd", Id: "4566e69f-c907-48ee-8d71-d7ba5aa00d20"},
				},
			},
			Description: mcproto.StatusDescription{
				Text: "Sex gay cu Simion",
			},
			Favicon:            mcproto.FaviconString(EmbeddedPng),
			EnforcesSecureChat: false,
		}

		ssPayload, err := mcproto.PrepareStatusPacketPayload(&ss)
		if err != nil {
			panic(err)
		}

		if err := mcproto.WritePacket(conn, mcproto.StatusPacketID, ssPayload); err != nil {
			panic(err)
		}

		fmt.Println("Sent status ping")

		pingId, pingPayload, err := mcproto.GetPacket(rd)
		if err != nil {
			panic(err)
		}

		if err := mcproto.WritePacket(conn, pingId, pingPayload); err != nil {
			panic(err)
		}
	}
}
