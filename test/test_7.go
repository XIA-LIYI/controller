package main

import (
	"net"
	"fmt"
)

func main() {
	fmt.Println("Listening")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.56.135:50120")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
	}
}