package main

import (
	"fmt"
	"net"
	"strconv"
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
		if n != 0 {
			// print msg. on port
			fmt.Println(string(buf[:n]))
			break
		}
	}

}
func udpSend() {
	conn, err := net.Dial("udp", UDPSendPort)
	if err != nil {
		panic(err)
	}
	n := 0
	for {
		t := strconv.Itoa(n)
		msg := "Hellow world!, " + t
		conn.Write([]byte(msg))
		n++
		time.Sleep(100)
	}
}

func main() {
	// listener(":30000")
	udpSend()
}
