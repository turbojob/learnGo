package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("test" + strconv.Itoa(i))
	}
}
func main() {
	go test() //start
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("主协程" + strconv.Itoa(i))
	}
}
