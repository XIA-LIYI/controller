package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

)

var status int32 = 0
var count int32 = 0
var totalByte uint64 = 0

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.112:50120")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go listen()

	startTime := time.Now()
	for {
		buf := make([]byte, 100)
		num, _ := conn.Read(buf)
		if (num == 0) {
			continue
		}
		fmt.Println(num)
		fmt.Println(string(buf)[:num])
		content := string(buf)[:num]
		if (content == "start") {
			startTime = time.Now()
			atomic.StoreInt32(&status, 1)
			fmt.Println(count)
			continue
		}
		if (content == "stop") {
			atomic.StoreInt32(&status, 0)
			break
		}
		tcpAddr, err := net.ResolveTCPAddr("tcp", content)
		if (err != nil) {
			fmt.Println(err)
		}
		newConn, err := net.DialTCP("tcp", nil, tcpAddr)
		if (err != nil) {
			fmt.Println(err)
		}
		atomic.AddInt32(&count, 1)
		go onReceive(newConn)
		go onSend(newConn)

	}
	elapsedTime := uint64(time.Since(startTime) / time.Millisecond / 1000)
	fmt.Println("Time consumed:", elapsedTime, "s")
	speed := totalByte / 1000 / elapsedTime / 8
	fmt.Println("Time consumed:", speed, "Kbps")

	// 控制台聊天功能加入
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

func listen() {
	fmt.Println("Listening")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.48.136:50120")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		go onReceive(tcpConn)
		go onSend(tcpConn)
		atomic.AddInt32(&count, 1)
	}
}

func onMessageRecived(conn *net.TCPConn) {
	// reader := bufio.NewReader(conn)
	// for {
	// 	msg, err := reader.ReadString('\n')
	// 	fmt.Println(msg)
	// 	if err != nil {
	// 		// quitSemaphore <- true
	// 		break
	// 	}
	// }

	
}

func onReceive(conn *net.TCPConn) {
	for atomic.LoadInt32(&status) == 0 {

	}
	for atomic.LoadInt32(&status) == 1 {
		buf := make([]byte, 12500000)
		num, _ := conn.Read(buf)
		atomic.AddUint64(&totalByte, uint64(num))
		fmt.Println(num)
	}

}

func onSend(conn *net.TCPConn) {
	for atomic.LoadInt32(&status) == 0 {

	}
	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()
	content := make([]byte, 12500000)
	for atomic.LoadInt32(&status) == 1 {
		// fmt.Println("start sending")
		for {
			<- ticker.C
			fmt.Println("tick.")
			conn.Write(content)
		}
	}

}

