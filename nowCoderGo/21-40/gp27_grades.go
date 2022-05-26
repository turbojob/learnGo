package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["小明"] = 60
	m1["小王"] = 70
	m1["张三"] = 95
	m1["李四"] = 98
	m1["王五"] = 100
	fmt.Println(m1)

}
