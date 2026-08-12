package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/VladiTNT/minegopher/pkg/mcproto"
)

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

	fmt.Printf("%+v\n", handShake)

	pid, pdata, err := mcproto.GetPacket(rd)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %q\n", pid, pdata)
}
