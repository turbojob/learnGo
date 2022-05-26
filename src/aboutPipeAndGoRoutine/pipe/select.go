package main

import (
	"fmt"
	"time"
)

func select1() {
	//多个管道的选择  多路复用
	//多个 选择一个 去执行

}
func main() {
	var pipe1 chan int = make(chan int, 1)
	var pipe2 chan string = make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 5)
		pipe1 <- 10

	}()
	go func() {
		time.Sleep(time.Second * 12)
		pipe2 <- "go"
	}()

	select {
	case v := <-pipe1:
		fmt.Println(v)
	case v := <-pipe2:
		fmt.Println(v)
	default:
		fmt.Println("防止select 阻塞")
	}

}
