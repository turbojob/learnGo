package main

import "sort"

func getNoRepeat(s []int) []int {
	// write code here
	var ans []int
	have := make(map[int]int)
	for _, i := range s {
		if _, ok := have[i]; ok == false {
			have[i] = 1
		} else {
			have[i]++
		}
	}
	for i, j := range have {
		if j == 1 {
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)
	return ans
}
