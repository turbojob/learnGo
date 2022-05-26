package main

func equal(s1 []int, s2 []int) bool {
	// write code here
	length := len(s1)
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
