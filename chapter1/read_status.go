package chapter1

import (
	"bufio"
	"fmt"
	"net"
)

func TCPDemo() {
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
