package main
import (
	"net"
	"fmt"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.48.135:50120")

	_, err := net.DialTCP("tcp", nil, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.48.133:50120")

	_, err = net.DialTCP("tcp", nil, tcpAddr)
	if (err != nil) {
		fmt.Println(err)
	}
}