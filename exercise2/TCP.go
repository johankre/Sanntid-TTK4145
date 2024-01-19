package main

import (
	"fmt"
	"net"
	"time"
)


const(
    TCPport1 = "10.100.23.129:34933"
    TCPport2 = ":33546"
)

func TCPListen1(conn net.Conn){
    buf := make([]byte, 1024)
	for {
        n, err := conn.Read(buf)
        if err != nil {
            panic(err)
        }
            fmt.Println(buf[:n])
        
	}

}

func TCPcon() (conn net.Conn){
    conn, err := net.Dial("tcp", TCPport1)
        if err != nil {
            panic(err)
        }
        return
}

func TCPCWrite(conn net.Conn) {
    _, err := conn.Write([]byte("test"))
        if err != nil {
            panic(err)
        }

}

func main(){
    conn := TCPcon()
    defer conn.Close()
    go TCPListen1(conn)
    time.Sleep(1000 * time.Millisecond)
}
