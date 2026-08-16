package main

import (
	"bytes"
	"fmt"
	"net"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
	"github.com/VladiTNT/minegopher/test/assets"
)

func main() {
	ln, err := net.Listen("tcp", ":25565")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Starting server")

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf bytes.Buffer

	id, data, err := mcproto.ReadPacket(conn)
	if err != nil {
		panic(err)
	}

	if id == mcproto.StatusID {
		h, err := mcproto.DecodeHandshake(data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %s %d %d\n", h.Protocol, h.Address, h.Port, h.Next)
	}

	id, data, err = mcproto.ReadPacket(conn)
	if err != nil {
		panic(err)
	}

	if id == mcproto.StatusID {
		ss := mcproto.ServerStatus{}
		ss.Version.Name = "26.2"
		ss.Version.Protocol = 776
		ss.Players.Max = 999999999
		ss.Players.Online = 80085
		ss.Description.Text = "Sex gay cu Simion"
		ss.Favicon = mcproto.EncodeFavicon(assets.VladFace)
		ss.EnforcesSecureChat = false

		mcproto.EncodeStatusResponse(&buf, &ss)
		mcproto.WritePacket(conn, &buf)
		buf.Reset()
	}

	id, data, err = mcproto.ReadPacket(conn)
	if err != nil {
		panic(err)
	}

	if id == mcproto.PingPongID {
		n, err := mcproto.DecodePingPong(data)
		if err != nil {
			panic(err)
		}

		if err := mcproto.EncodePingPong(&buf, n); err != nil {
			panic(err)
		}

		if err := mcproto.WritePacket(conn, &buf); err != nil {
			panic(err)
		}
	} else {
		panic(id)
	}
}
