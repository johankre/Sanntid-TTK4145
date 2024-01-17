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

func main(){
    go TCPListen1()
    time.Sleep(1000000)
}
