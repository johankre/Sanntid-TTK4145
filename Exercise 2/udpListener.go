package main

import (
	"fmt"
	"net"
)

func listener(port string) {
	pc, err := net.ListenPacket("udp", port)
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)

	for {
		n, _, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		if n != 0 {
			fmt.Println(string(buf[:n]))
		}
	}

}

func main() {
	listener(":30000")
}
