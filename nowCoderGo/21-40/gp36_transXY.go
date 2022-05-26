package main

func transform(array [][]int) []int {
	// write code here
	m := len(array)
	n := len(array[0])
	var arr []int = make([]int, m*n)
	i := 0
	for a := 0; a < m; a++ {
		for b := 0; b < n; b++ {
			arr[i] = array[a][b]
			i++
		}
	}
	return arr
}
