package main

import (
	"fmt"
	"learnGo/aboutPackage/calculator"
	"learnGo/aboutPackage/util"
)

func main() {
	fmt.Println("我的第一个 main程序")
	fmt.Println(util.CalSum(10, 30))
	fmt.Println(calculator.Sum(10, 20))
}
