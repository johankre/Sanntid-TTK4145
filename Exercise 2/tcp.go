package main

import (
	"fmt"
	"net"
)

func sendMessageToServer(serverAddr, message string) error {
	// Resolve the server address
	serverTCPAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("Error resolving server address: %v", err)
	}

	// Establish a TCP connection to the server
	conn, err := net.DialTCP("tcp", nil, serverTCPAddr)
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send the message to the server
	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("Error sending message to server: %v", err)
	}

	fmt.Printf("Message sent to server %s: %s\n", serverAddr, message)

	return nil
}

func main() {
	serverAddress := "10.100.23.129:34933" // Replace with the actual IP address and port of your server
	messageToSend := "Hello, server!"

	err := sendMessageToServer(serverAddress, messageToSend)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
