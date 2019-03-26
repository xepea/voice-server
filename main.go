package main

func main() {
	server := Server{
		Users:     make(map[User]bool),
		Broadcast: make(chan Packet),
	}
	server.Start()
}
