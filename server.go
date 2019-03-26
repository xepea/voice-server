package main

import (
	"fmt"
	"net"
)

// Server @ server struct
type Server struct {
	Users     map[User]bool
	Broadcast chan Packet
}

// Start @ starts voip server
func (server *Server) Start() {
	serverAddr := net.UDPAddr{
		Port: 2000,
		IP:   net.ParseIP("127.0.0.1"),
	}
	listener, err := net.ListenUDP("udp", &serverAddr)
	if err != nil {
		fmt.Println(err)
	}
	go server.Manager()
	server.Receive(listener)
}

// Manager @ manages packets
func (server *Server) Manager() {
	for {
		select {
		case packet := <-server.Broadcast:
			for user := range server.Users {
				user.Flow <- packet
			}
		}
	}
}

// Send @ sends udp packets
func (server *Server) Send(listener *net.UDPConn, user *User) {
	for {
		select {
		case packet, ok := <-user.Flow:
			if !ok {
				return
			}
			listener.WriteToUDP(*packet.Data, user.Addr)
		}
	}
}

// Receive @ receives udp packets
func (server *Server) Receive(listener *net.UDPConn) {
	for {
		buffer := make([]byte, 4096)
		_, addr, err := listener.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		user := User{Addr: addr, Flow: make(chan Packet)}
		exists := false
		for u := range server.Users {
			if u.Addr.String() == addr.String() {
				exists = true
			}
		}
		packet := Packet{Addr: addr, Data: &buffer}
		if !exists {
			server.Users[user] = true
			go server.Send(listener, &user)
			server.Broadcast <- packet
		} else {
			server.Broadcast <- packet
		}
	}
}
