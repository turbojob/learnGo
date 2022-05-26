package main

import "fmt"

func main() {
	var s1 = make([]string, 5, 5)
	s1[0] = "帽子"
	s1[1] = "围巾"
	s1[2] = "衣服"
	s1[3] = "裤子"
	s1[4] = "袜子"
	fmt.Println(s1)
}
