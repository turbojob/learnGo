package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8888")
	fmt.Println(conn)
}
