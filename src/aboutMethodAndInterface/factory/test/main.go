package main

import (
	"fmt"
	"learnGo/aboutMethodAndInterface/factory/student"
)

func main() {
	s := student.NewStudent("zb", 200)
	fmt.Println(*s)
}
