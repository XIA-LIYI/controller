package main

import (
	// "bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	// "sync/atomic"
	// "time"
)

var connections = make([]*net.TCPConn, 25)
var ips [25]string
var count int = 0
var allReady bool = false
var numOfNodesReady int32 = 0
var canClose chan int = make(chan int)

func main() {
	go monitorInput()
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.112:18787")

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	} 

	for {
		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			continue
		}
		
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		fmt.Println("Total number of connections:", count)

		connections[count] = tcpConn
		ips[count] = tcpConn.RemoteAddr().String()
		count += 1
		
		// go tcpPipe(tcpConn)
		for i := 0; i< count; i++ {
			// conn.Write([]byte(tcpConn.RemoteAddr().String()))
			ipAddr := strings.Split(tcpConn.RemoteAddr().String(), ":")[0]
			connections[i].Write([]byte(ipAddr + ":" + strconv.Itoa(5050) + "\n"))
			// conn.Write([]byte("192.168.56.135:10000"))
		}
		
		
		check()
		if (count == 25) {
			tcpListener.Close()
			break
		}
	}
	
	fmt.Println("check for check, start for start, stop for stop")
	start()
	fmt.Println("starting")
	time.Sleep(time.Second * 120)
	getResult()
	<- canClose
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

func monitorInput() {
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "check" {
			for _, conn := range connections {
				go listen(conn)
				conn.Write([]byte("check\n"))
			}
		}
		if msg == "start" {
			start()
		}
		if msg == "stop" {
			getResult()
			canClose <- 1
		}
	}

}

func getResult() {
	for i := 0; i < count; i++ {
		fmt.Printf(ips[i] + ": ")
		for {
			connections[i].Write([]byte("stop\n"))
			buf := make([]byte, 150)
			num, err := connections[i].Read(buf)
			if (err != nil) {
				continue
			}
			content := string(buf)[:num]
			fmt.Printf(content)
			fmt.Printf("\n")
			break
		}
	}
}

func listen(conn *net.TCPConn) {
	for {
		buf := make([]byte, 100)
		num, _ := conn.Read(buf)
		content := string(buf)[:num]
		fmt.Println(content)
	}
}

func check() {
	for i := 0; i < count; i++ {
		fmt.Printf("Checking ip:" + ips[i] + " ")
		for {
			connections[i].Write([]byte("check\n"))
			buf := make([]byte, 100)
			num, _ := connections[i].Read(buf)
			content := string(buf)[:num]
			fmt.Println(content)
			if (content == strconv.Itoa(int(count - 1))) {
				break
			} else {
				continue
			}
		}
		time.Sleep(time.Second / 100)
	}

}

func start() {
	for _, conn := range connections {
		conn.Write([]byte("start\n"))
	}

}

// func tcpPipe(conn *net.TCPConn) {

// 	defer conn.Close()

// 	reader := bufio.NewReader(conn)

// 	for {
// 		message, err := reader.ReadString('\n')
// 		if err != nil {
// 			return
// 		}
// 		broadcast(conn.RemoteAddr().String() + ":" + string(message))
// 	}

// }

// func broadcast(message string) {
// 	for _, conn := range connectionMap {
// 		conn.Write([]byte(message))
// 	}
// }

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
