package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6}
	copy(slice1, slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)
	//[5 6 3 4]
	//[5 6]
	slice3 := []int{1, 2, 3, 4}
	slice4 := []int{5, 6}
	copy(slice4, slice3)
	fmt.Println(slice3)
	fmt.Println(slice4)
	//[1 2 3 4]
	//[1 2]
	//如果有空的呢？
	//全是
	//[]
	//[1 2]
	//[]
	//[1 2]
	slice6 := make([]int, 0, 0)
	slice7 := []int{1, 2}
	copy(slice6, slice7)
	fmt.Println(slice6)
	fmt.Println(slice7)

	slice8 := make([]int, 0, 0)
	slice9 := []int{1, 2}
	copy(slice9, slice8)
	fmt.Println(slice8)
	fmt.Println(slice9)

}
func sliceCopy(src []int, des []int) []int {
	// write code here

	return src
}
