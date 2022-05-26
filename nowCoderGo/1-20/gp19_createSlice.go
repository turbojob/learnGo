package main

func main() {

}
func makeslice(length int, capacity int) []int {
	// write code here
	var slice []int = make([]int, length, capacity)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	return slice
}
