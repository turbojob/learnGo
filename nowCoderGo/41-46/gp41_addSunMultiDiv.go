package main

func operate(a int, b int) []int {
	// write code here
	var ans []int
	ans = append(ans, a+b)
	ans = append(ans, a-b)
	ans = append(ans, a*b)
	ans = append(ans, a/b)
	ans = append(ans, a%b)
	return ans
}
