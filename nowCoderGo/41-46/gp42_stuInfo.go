package main

import "fmt"

type stu1 struct {
	name  string
	sex   bool
	age   int
	score int
}

func main() {
	a := stu1{
		name:  "小明",
		age:   23,
		sex:   true,
		score: 88,
	}
	fmt.Println(a.name)
	fmt.Println(a.sex)
	fmt.Println(a.age)
	fmt.Println(a.score)
}
