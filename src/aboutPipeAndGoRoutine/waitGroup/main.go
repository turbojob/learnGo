package main

import (
	"fmt"
	"sync"
)

func main() {
	var count sync.WaitGroup = sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		count.Add(1)
		go func(n int) {
			defer count.Done()
			fmt.Println(n)
		}(i)
	}
	count.Wait()

}
