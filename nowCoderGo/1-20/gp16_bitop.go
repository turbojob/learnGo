package main

func bitOperate(a int, b int) []int {
	// write code here
	var slice []int
	slice = append(slice, a&b)
	slice = append(slice, a|b)
	slice = append(slice, a^b)
	return slice
}
