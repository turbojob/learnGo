package main

func compare(x int, y int, z int) int {
	// write code here
	var max1 int
	var min1 int
	max1 = max(x, max(y, z))
	min1 = min(x, min(y, z))
	return max1 - min1
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
