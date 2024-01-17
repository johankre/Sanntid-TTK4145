package main

import (
	"fmt"
	"net"
)

func sendTCPMessage(message string, serverAddr string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Convert the message to bytes
	messageBytes := []byte(message)

	// Send the message to the server
	_, err = conn.Write(messageBytes)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	// Read the response from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Server response: %s\n", buffer[:n])
}

func main() {
	serverAddr := "10.100.23.129:33546" // Set the server address and port

	// Send a message to the server
	message := "Stay humble, stack sats"
	sendTCPMessage(message, serverAddr)
}
