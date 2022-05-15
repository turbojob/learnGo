package main

import (
	"fmt"
)

func pipe() {
	var channel1 chan int = make(chan int, 10)

	var writeOnlyChannel chan<- int
	writeOnlyChannel = make(chan int, 3)

	var readOnlyChannel <-chan int = make(chan int, 10)
	fmt.Println(writeOnlyChannel, readOnlyChannel)
	channel1 <- 10
	channel1 <- 100
	fmt.Println(len(channel1))

	num := <-channel1
	fmt.Println(num)

	for v := range channel1 {
		fmt.Println(v)
	}
	//关闭后只能读取数据
	close(channel1)
}
