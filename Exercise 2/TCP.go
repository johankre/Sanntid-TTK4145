package TCP

import (
	"fmt"
	"net"
	"time"
)

const(
    TCPport1 = ":34933"
    TCPport2 = ":33546"
)

func TCPListen1(){
    l, err := net.Listen("tcp", TCPport1)

    if err != nil {
        panic(err)
    }
    
    defer l.Close()

    buf := make([]byte, 1024)
	for {
        conn, err := l.Accept()
        if err != nil {
            panic(err)
        }
        n, err := conn.Read(buf)
        if err != nil {
            panic(err)
        }
        if n > 0 {
            fmt.Println(buf[:n])
        }
        
	}

}

func TCPcon() {
    conn, err := net.Dial("tcp", TCPport1)
        if err != nil {
            panic(err)
        }
    defer conn.Close()

    _, err = conn.Write([]byte("test"))
        if err != nil {
            panic(err)
        }
}

func main(){
    go TCPListen1()
    TCPcon()
    time.Sleep(1000000)
}
