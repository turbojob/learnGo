package main

func factorial(i int) int {
	// write code here
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}
