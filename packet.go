package main

import "net"

// Packet @ packet struct
type Packet struct {
	Addr *net.UDPAddr
	Data *[]byte
}
