package main

import (
	"fmt"
	"net"
)

/**
 *
 */
func main() {

	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error listen", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("start listen 8081")

	for {
		accept, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept", err.Error())
			return
		}
		go func(conn net.Conn) {
			buffer := make([]byte, 1024)
			read, err2 := conn.Read(buffer)
			if err2 != nil {
				fmt.Println("error reading", err.Error())
				conn.Close()
				return
			}
			msg := string(buffer[:read])
			fmt.Println("received msg", msg)
			conn.Close()
		}(accept)
	}
}
