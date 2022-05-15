package main

import (
	"fmt"
	"learnGo/aboutMethodAndInterface/demo2"
)

func main() {
	//创建结构体实例时指定字段值
	var s1 = demo2.Student{"zhaobo", 18}
	fmt.Println(s1)
	var s2 demo2.Student
	s2.Age = 100
	fmt.Println(s2)
}
func test1() {
	//跨包创建结构体
}
