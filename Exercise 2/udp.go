package main

import (
	"fmt"
	"net"
)

func startUDPServer(port string, quit chan int) {
	// Resolve the address to bind to
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	fmt.Println("Resolved UDP Address:", addr)

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
		if n != 0 {
			quit <- 0
		}
	}

}

func broadcastUDPMessage(message string, port string) {
	// Resolve the broadcast address
	broadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:"+port)
	if err != nil {
		fmt.Println("Error resolving broadcast address:", err)
		return
	}

	// Create a UDP connection to broadcast
	conn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close()

	// Convert the message to bytes
	messageBytes := []byte(message)

	// Send the message to all devices on the network
	_, err = conn.Write(messageBytes)
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}

	fmt.Println("Message broadcasted successfully:", message)

}

func main() {
	port := "20024" // Set the desired UDP port
	quit := make(chan int, 0)

	// Start the UDP server in a goroutine
	go startUDPServer(port, quit)

	// Broadcast a message
	message := "Stay humble, stack sats"
	broadcastUDPMessage(message, port)

	// Keep the program running
	<-quit
}
