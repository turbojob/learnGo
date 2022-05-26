package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scanf("%d%d", &a, &b)
	var myslice []bool
	myslice = append(myslice, a == b)
	myslice = append(myslice, &a == &b)
	fmt.Println(myslice)

}
func equal(a int, b int) []bool {
	// write code here
	var myslice []bool

	myslice = append(myslice, &a == &b)
	myslice = append(myslice, a == b)
	return myslice
}
