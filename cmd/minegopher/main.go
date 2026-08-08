package main

import (
	"context"
	"fmt"
	"os"

	"github.com/VladiTNT/minegopher/pkg/lancast"
)

func main() {
	if err := lancast.BroadcastLAN(context.Background(), "nerd", 25565); err != nil {
		fmt.Fprintf(os.Stderr, "Error with lan upd multicast: %v\n", err)
		os.Exit(1)
	}
}
