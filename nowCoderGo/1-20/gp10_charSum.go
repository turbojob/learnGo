package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum("123", "23"))
}
func sum(a string, b string) string {
	// write code here
	a1, _ := strconv.ParseInt(a, 0, 0)
	a2, _ := strconv.ParseInt(b, 0, 0)
	s1 := strconv.Itoa(int(a1 + a2))
	return s1

}
