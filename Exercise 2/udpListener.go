package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func listener() {
	l, err := net.Listen("udp", ":30000")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		// Listen for connection
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
}

func main(){
    listener()
}
