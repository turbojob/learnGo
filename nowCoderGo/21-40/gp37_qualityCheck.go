package main

func check(material []int, standard int) []int {
	// write code here
	var ans []int
	for _, j := range material {
		if j >= standard {
			ans = append(ans, j)
		}
	}
	return ans
}
