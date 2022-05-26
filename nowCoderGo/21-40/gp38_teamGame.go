package main

func canPass(score []int, target int) bool {
	// write code here
	for _, j := range score {
		if j > target {
			return true
		}
	}
	return false
}
