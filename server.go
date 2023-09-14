package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var connectionMap map[string]*net.TCPConn
var count = 0

func main() {
	var tcpAddr *net.TCPAddr
	connectionMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.112:10000")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			continue
		}
		count += 1
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		fmt.Println("Total number of connections:", count)
		
		// go tcpPipe(tcpConn)
		for _, conn := range connectionMap {
			// conn.Write([]byte(tcpConn.RemoteAddr().String()))
			ipAddr := strings.Split(tcpConn.RemoteAddr().String(), ".")
			// fmt.Println(ipAddr)
			conn.Write([]byte(ipAddr[0] + "." + ipAddr[1] + "." + "56" + ipAddr[3] + ":" + strconv.Itoa(10000)))
		}
		connectionMap[tcpConn.RemoteAddr().String()] = tcpConn
		if (count == 2) {
			break
		}
	}
	fmt.Println("Start? Type yes/no")
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "yes" {
			go start()
		}
		if msg == "no" {
			go stop()
		}
	}
	// for {
	// 	var msg string
	// 	fmt.Scanln(&msg)
	// 	if msg == "quit" {
	// 		break
	// 	}
	// 	b := []byte(msg + "\n")
	// 	conn.Write(b)
	// }

}

func start() {
	for _, conn := range connectionMap {
		conn.Write([]byte("start"))
	}

}

func stop() {
	for _, conn := range connectionMap {
		conn.Write([]byte("stop"))
	}
}

func tcpPipe(conn *net.TCPConn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		broadcast(conn.RemoteAddr().String() + ":" + string(message))
	}

}

func broadcast(message string) {
	for _, conn := range connectionMap {
		conn.Write([]byte(message))
	}
}

// func process(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		var buf [128]byte
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Printf("read from conn failed, err:%v", err)
// 		}
// 		fmt.Printf("recv from client, content:%v\n", string(buf[:n]))
// 	}
// }
