package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(12321))
}
func isPalindrome(x int) bool {
	// write code here
	var str string = strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true

}
