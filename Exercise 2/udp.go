package main

import (
	"fmt"
	"net"
)

func startUDPServer(port string) {
	// Resolve the address to bind to
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP listener
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server is listening on port", port)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		// Print received data
		fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, buffer[:n])
	}
}

func main() {
	port := "30000" // Set the desired UDP port
	startUDPServer(port)
}
