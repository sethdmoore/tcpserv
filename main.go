package main

import (
	"fmt"
	"net"
	//"os"
	"strings"
	//"github.com/sethdmoore/tcpserv/envconfig"
)

func light(p string) {

}

func handleReq(conn net.Conn) {
	//a := map[string]func(string)
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error reading: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))
	err := conn.Write([]byte("recvd\n"))
	if err != nil {
		fmt.Printf("Error writing: %s\n", err)
	}
	err := conn.Close()
	if err != nil {
		fmt.Printf("Error closing: %s\n", err)
	}
}

func main() {
	host := "0.0.0.0"
	port := "8999"
	conn := strings.Join([]string{host, port}, ":")
	l, err := net.Listen("tcp", conn)
	if err != nil {
		fmt.Printf("Could not listen on ")
	}
	defer l.Close()
	fmt.Printf("Listening on %s\n", conn)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Couldn't accept conn: %s", err)
			continue
		}

		go handleReq(conn)
	}
}
