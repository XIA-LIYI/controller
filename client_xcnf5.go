package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

)

var count int32 = 0
var totalByte uint64 = 0
var chans = []chan int {
	make(chan int),
	// make(chan int),
}

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.51.112:10000")
	localAddr, err := net.ResolveTCPAddr("tcp", "192.168.56.133:10000")
	conn, err := net.DialTCP("tcp", localAddr, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("connected!")

	go listen()

	startTime := time.Now()
	for {
		fmt.Println("waiting")
		buf := make([]byte, 100)
		num, _ := conn.Read(buf)
		content := string(buf)[:num]
		fmt.Println(num, content)
		if (content == "start") {
			startTime = time.Now()
			fmt.Println("Current number of connections is:", count)
			for i := range chans {
				chans[i] <- 0
			}
			fmt.Println("All are released!")
			continue
		}
		if (content == "stop") {
			break
		}
		for {
			addr, _ := net.ResolveTCPAddr("tcp", content)
			newConn, err := net.DialTCP("tcp", nil, addr)
			if (err != nil) {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("connected!")
			}
			go onReceive(newConn)
			go onSend(newConn, chans[count])
			atomic.AddInt32(&count, 1)
			break
		}
	}
	elapsedTime := uint64(time.Since(startTime) / time.Millisecond / 1000)
	fmt.Println("Time consumed:", elapsedTime, "s")
	speed := totalByte / 1000 / elapsedTime * 8 / 1000
	fmt.Println("Speed is:", speed, "Mbps")

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
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10000")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		go onReceive(tcpConn)
		go onSend(tcpConn, chans[count])
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
	fmt.Println("start receiving")
	for {
		buf := make([]byte, 12500000)
		num, _ := conn.Read(buf)
		atomic.AddUint64(&totalByte, uint64(num))
	}

}

func onSend(conn *net.TCPConn, ch chan int) {
	fmt.Println("start sending")
	<- ch
	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()
	content := make([]byte, 12500000)
	for {
		// fmt.Println("start sending")
		for {
			<- ticker.C
			conn.Write(content)
		}
	}

}

