package main

import "fmt"

func initSlice() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	//fmt.Println(months)
	//fmt.Println("Length:", len(months))
	//fmt.Println("Capacity:", cap(months))
	quarter2 := months[3:6]
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
}
func append1() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}
func removeFromSlice() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)

		slice2 := make([]string, 3)
		fmt.Println(slice2)
		copy(slice2, letters[1:4])
	}
}
func main() {
	removeFromSlice()
}
