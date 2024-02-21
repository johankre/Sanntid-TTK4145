package main

import (
    "main/network"
	"net"
    "time"
)


func main() {
    go network.Listener(network.UDPSendPort)


    var conn net.Conn = network.UdpInitDail(network.UDPSendPort)
    defer conn.Close()

    var testQueue = [4][3]int{
        {1, 2, 3},
        {3, 4, 5},
        {1, 4, 5},
        {1, 2, 5},
    }

    packet := network.Packet{
        Version: 1,
        ElevatorNum: 2,
        Guid: 23,
        Queue: testQueue,
    }

    for { 
        network.BrodcastPacket(packet, conn)
        time.Sleep(1000 * time.Millisecond)
    }
}
