package main

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
)

func main() {
	conn, err := net.Dial("tcp", "mc.hypixel.net:25565")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf bytes.Buffer
	if err := mcproto.EncodeHandshake(&buf, 47, []byte("mc.hypixel.net"), 25565, 1); err != nil {
		panic(err)
	}

	fmt.Println("Sending handshake")
	if err := mcproto.WritePacket(conn, &buf); err != nil {
		panic(err)
	}

	buf.Reset()

	if err := mcproto.EncodeStatusRequest(&buf); err != nil {
		panic(err)
	}

	fmt.Println("Sending status request")
	if err := mcproto.WritePacket(conn, &buf); err != nil {
		panic(err)
	}

	buf.Reset()

	id, data, err := mcproto.ReadPacket(conn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %q\n", id, data)

	if id == mcproto.StatusID {
		ss, favicon, err := mcproto.DecodeStatusResponse(data)
		if err != nil {
			panic(err)
		}

		fmt.Println(ss.Version.Name, ss.Players.Online, ss.Description)
		if err := os.WriteFile("hypixel.net.png", favicon, os.ModeAppend); err != nil {
			panic(err)
		}
	}
}
