package main

import "net"

// User @ user struct
type User struct {
	Addr *net.UDPAddr
	Flow chan Packet
}
