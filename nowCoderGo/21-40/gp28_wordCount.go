package main

import "fmt"

func main() {
	fmt.Println(character("aaac"))
}
func character(s string) byte {
	// write code here
	arr := ([]byte)(s)
	count := 0
	var ans byte
	charaMap := make(map[byte]int)
	for _, j := range arr {

		charaMap[j]++
		if charaMap[j] > count {
			count = charaMap[j]
			ans = j
		}
	}

	return ans
}
