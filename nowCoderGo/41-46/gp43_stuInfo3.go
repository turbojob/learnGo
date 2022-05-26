package main

import (
	"fmt"
)

type address struct {
	city     string
	province string
}
type stu struct {
	name  string
	sex   bool
	age   int
	score int
	add   address
}

func main() {
	a := stu{
		name:  "小明",
		age:   23,
		sex:   true,
		score: 88,
		add: address{
			city:     "长沙市",
			province: "湖南省",
		},
	}
	fmt.Println(a.name)
	fmt.Println(a.sex)
	fmt.Println(a.age)
	fmt.Println(a.score)
	fmt.Println(a.add.province)
	fmt.Println(a.add.city)
}
