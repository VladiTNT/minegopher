package lanmulticast

import (
	"context"
	"net"
	"time"
)

const (
	// This is the address where Minecraft client's expect to receive UDP packets for LAN detection.
	DefaultMulticastAddress string = "224.0.2.60:4445"
	// This is the broadcast interval that most Minecraft servers use.
	DefaultBroadcastInterval time.Duration = 1500 * time.Millisecond
)

// This is a LAN multicast service, it makes the server visible in the LAN tab. The ServerPort field has to be
// the same as the port that the actual minecraft server is running on, otherwise clients can't join.
type Service struct {
	MulticastAddress  string
	BroadcastInterval time.Duration
	MessageOfTheDay   string
	ServerPort        uint16
}

// This is a constructor for a Service, it uses the default values for the address and interval.
func NewService(motd string, serverPort uint16) *Service {
	return &Service{
		MulticastAddress:  DefaultMulticastAddress,
		BroadcastInterval: DefaultBroadcastInterval,
		MessageOfTheDay:   motd,
		ServerPort:        serverPort,
	}
}

// Run starts the multicast service.
func (s *Service) Run(ctx context.Context) error {
	udpPacket := NewMotdPacket(s.MessageOfTheDay, s.ServerPort)

	udpAddr, err := net.ResolveUDPAddr("udp", s.MulticastAddress)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	ticker := time.NewTicker(s.BroadcastInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			_, err := conn.Write(udpPacket)
			if err != nil {
				return err
			}
		}
	}
}
