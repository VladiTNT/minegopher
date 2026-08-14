package lanmulticast

import "fmt"

// NewMotdPacket makes an motd string that can be sent over udp to Minecraft clients to be seen in the LAN tab.
func NewMotdPacket(motd string, port uint16) []byte {
	return fmt.Appendf([]byte{}, "[MOTD]%s[/MOTD][AD]%d[/AD]", motd, port)
}
