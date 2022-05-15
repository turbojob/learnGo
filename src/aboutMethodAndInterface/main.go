package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

func perimeter1(t triangle) int {
	return t.size * 3
}
func (t triangle) perimeter2() int {
	return t.size * 3
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func StoB() {
	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}

func initMethod() {
	//方法是一种特殊类型的函数，但存在一个简单的区别：
	//你必须在函数名称之前加入一个额外的参数。
	//此附加参数称为接收方。
	t := triangle{3}
	fmt.Println("Per", t.perimeter2())
}
func main() {
	initMethod()
}
