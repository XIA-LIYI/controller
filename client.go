package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
	"strconv"

)

var count int32 = 0
var totalByte uint64 = 0
var chans = []chan int {
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	// make(chan int),
	// make(chan int),
	// make(chan int),
}

var bytes [7]uint64

var sendingByte int = 1250000000 / 8

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.51.112:18787")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("connected!")

	go listen()

	startTime := time.Now()
	buf := make([]byte, 100)
	for {
		num, _ := conn.Read(buf)
		content := string(buf)[:num]
		fmt.Println(num, content)
		if (content == "check") {
			conn.Write([]byte(strconv.Itoa(int(count))))
			continue
		}
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
			time.Sleep(time.Second)
			if (err != nil) {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("connected!")
			}
			go onReceive(newConn, count)
			go onSend(newConn, chans[count])
			atomic.AddInt32(&count, 1)
			break
		}
	}
	elapsedTime := uint64(time.Since(startTime) / time.Millisecond / 1000)
	printSpeed(elapsedTime)

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

func printSpeed(elapsedTime uint64) {
	fmt.Println("Time consumed:", elapsedTime, "s")
	totalSpeed := totalByte / 1000 / elapsedTime * 8 / 1000
	fmt.Println("Speed is:", totalSpeed, "Mbps")
	for _, i := range bytes {
		speed := i / 1000 / elapsedTime * 8 / 1000
		fmt.Println("Single speed is:", speed, "Mbps")
	}
}

func listen() {
	fmt.Println("Listening")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:5050")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected:" + tcpConn.RemoteAddr().String())
		go onReceive(tcpConn, count)
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

func onReceive(conn *net.TCPConn, index int32) {
	// fmt.Println("start receiving")
	// conn.SetReadBuffer(128000)
	buf := make([]byte, 156250)
	for {
		num, _ := conn.Read(buf)
		atomic.AddUint64(&totalByte, uint64(num))
		bytes[index] += uint64(num)
	}

}

func onSend(conn *net.TCPConn, ch chan int) {
	// fmt.Println("start sending")
	<- ch
	ticker := time.NewTicker(time.Second / 1000)
	defer ticker.Stop()
	// conn.SetWriteBuffer(1000000)
	content := make([]byte, 156250)

	// fmt.Println("start sending")
	for {
		// <- ticker.C
		conn.Write(content)
	}


}

