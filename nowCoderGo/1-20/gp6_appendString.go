package main

import "fmt"

func main() {
	var arr = []string{"hello", "-", "world"}

	fmt.Println(join(arr))
}
func join(s []string) string {
	// write code here
	var mySlice string
	for _, cur := range s {
		mySlice = mySlice + cur
	}
	return mySlice

}
