package lancast

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	DefaultAddr     = "224.0.2.60:4445"
	DefaultInterval = 1500 * time.Millisecond
)

func CustomBroadcastLAN(ctx context.Context, addr string, motd string, port int, interval time.Duration) error {
	udpPacket := fmt.Appendf([]byte{}, "[MOTD]%s[/MOTD][AD]%d[/AD]", motd, port)

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			_, err := conn.Write(udpPacket)
			if err != nil {
				return err
			}
		}
	}
}

func BroadcastLAN(ctx context.Context, motd string, port int) error {
	return CustomBroadcastLAN(ctx, DefaultAddr, motd, port, DefaultInterval)
}
