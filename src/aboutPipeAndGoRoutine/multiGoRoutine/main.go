package main

import (
	"fmt"
	"time"
)

func main() {
	//1个
	go func() {
		fmt.Println("start 1 go")
	}()
	for i := 1; i <= 5; i++ {

		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(10000000)
}
