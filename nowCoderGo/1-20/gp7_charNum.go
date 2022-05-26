package main

import "fmt"

func main() {
	var str = string("小明的英文名叫jack")
	fmt.Println(count(str))
}
func count(s string) int {
	// write code here
	var arr = []rune(s)
	ans := len(arr)
	return ans
	//等价于
	//var arr []int32    32 位，对应4个字节

}
