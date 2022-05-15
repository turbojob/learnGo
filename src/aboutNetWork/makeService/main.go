package main

import (
	"fmt"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", "8888")
	fmt.Println("server run")
	fmt.Println(lis)
	//等待客户端的链接
	for {
		conn, _ := lis.Accept()
		fmt.Println("来了 新的连接")
		fmt.Println(conn.RemoteAddr().String())
	}

}
