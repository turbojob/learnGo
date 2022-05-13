package main

import (
	"fmt"
)

func initmap() {
	studentsAge := make(map[string]int)
	studentsAge["zb"] = 1
	studentsAge["zb"] = studentsAge["zb"] - 1
	ans, exist := studentsAge["zzz"]
	switch {
	case exist:
		fmt.Println(ans)
	default:
		fmt.Println("不存在")
	}
	fmt.Println(studentsAge["zzz"])
}
func delete1() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for k, v := range studentsAge {
		fmt.Println(k, v)
	}

}
func main() {
	delete1()
}
