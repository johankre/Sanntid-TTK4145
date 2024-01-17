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
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
        fmt.Println(addr.String())
		if n != 0 {
			fmt.Println(string(buf[:n]))
            break
		}
	}

}

func main() {
	listener(":30000")
}
