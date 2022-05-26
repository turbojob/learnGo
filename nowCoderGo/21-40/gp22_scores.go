package main

func minAndMax(s []int) []int {
	// write code here
	var ans []int
	var a int = s[0]
	var b int = s[0]
	for _, j := range s {
		a = max(a, j)
		b = min(b, j)
	}

	ans = append(ans, b)
	ans = append(ans, a)
	return ans
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
