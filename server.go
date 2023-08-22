package main

import (
	"fmt"
	"net"
)

var ConnectionMap map[string]*net.TCPConn
var count = 0

func main() {
	var TcpAddr *net.TCPAddr
	ConnectionMap = make(map[string]*net.TCPConn)
	TcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8000")

	TcpListener, _ := net.ListenTCP("tcp", TcpAddr)

	defer TcpListener.Close()

	for {
		TcpConn, err := TcpListener.AcceptTCP()

		if err != nil {
			continue
		}
		fmt.Println("A client connected:" + TcpConn.RemoteAddr().String())

	}

	
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
		}
		fmt.Printf("recv from client, content:%v\n", string(buf[:n]))
	}
}
