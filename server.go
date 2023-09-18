package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var connectionMap map[string]*net.TCPConn
var count int = 0
var allReady bool = false
var numOfNodesReady int32 = 0

func main() {
	var tcpAddr *net.TCPAddr
	connectionMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.112:18787")

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			continue
		}
		go listen(tcpConn)
		count += 1
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		fmt.Println("Total number of connections:", count)
		
		// go tcpPipe(tcpConn)
		for _, conn := range connectionMap {
			// conn.Write([]byte(tcpConn.RemoteAddr().String()))
			ipAddr := strings.Split(tcpConn.RemoteAddr().String(), ":")[0]
			fmt.Println(ipAddr)
			conn.Write([]byte(ipAddr + ":" + strconv.Itoa(5050)))
			// conn.Write([]byte("192.168.56.135:10000"))
		}
		connectionMap[tcpConn.RemoteAddr().String()] = tcpConn
		for {
			for _, conn := range connectionMap {
				conn.Write([]byte("check"))
			}
			if (numOfNodesReady == int32(count - 1)) {
				numOfNodesReady = 0
				break
			}
			time.Sleep(time.Second)
		}
		if (numOfNodesReady == 15) {
			break
		}
	}
	
	for {
		fmt.Println("check for check, yes for start, no for stop")
		var msg string
		fmt.Scanln(&msg)
		if msg == "check" {
			for _, conn := range connectionMap {
				conn.Write([]byte("check"))
			}
		}
		if msg == "yes" {
			start()
		}
		if msg == "no" {
			stop()
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

func listen(conn *net.TCPConn) {
	buf := make([]byte, 100)
	for {
		num, _ := conn.Read(buf)
		content := string(buf)[:num]
		fmt.Println(content)
		if (content == strconv.Itoa(count)) {
			atomic.AddInt32(&numOfNodesReady, 1)
		}
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
