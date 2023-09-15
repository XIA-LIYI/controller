package main
import (
	"net"
	"fmt"
)

func main() {
	var tcpAddr *net.TCPAddr
	for {
		tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.83:50120")
	}
	

	_, err := net.DialTCP("tcp", nil, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.51.83:50120")

	_, err = net.DialTCP("tcp", nil, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}
}