package main

import (
	"fmt"
	"net"
	"time"
)

// Globals
const UDPSendPort string = ":20024"

func listener(port string) {
	pc, err := net.ListenPacket("udp", port)
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		// print IP addres
		fmt.Println(addr.String())
			// print msg. on port
			fmt.Println(string(buf[:n]))
            
	}

}
func udpSend() {
	conn, err := net.Dial("udp", UDPSendPort)
	if err != nil {
		panic(err)
	}

    defer conn.Close()

        _, err = conn.Write([]byte("test"))
        if err != nil {
            panic(err)
        }
    }

func main() {
	go listener(UDPSendPort)
	udpSend()

    time.Sleep(1000000)
}
