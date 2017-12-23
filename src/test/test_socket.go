package main

import (
	"net"
	"fmt"
)

func main() {
	Conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if (err != nil) {
		fmt.Println("can not access")
		return
	}
	Conn.Write([]byte("asdjio"))
}
