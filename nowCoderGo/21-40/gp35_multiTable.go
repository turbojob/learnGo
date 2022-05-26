package main

import "fmt"

func main() {
	fmt.Println("1*1=1")
	for i := 2; i <= 9; i++ {
		for j := 1; j <= i-1; j++ {
			str := fmt.Sprintf("%d*%d=%-3d", j, i, j*i)

			fmt.Printf("%s", str)
		}
		str := fmt.Sprintf("%d*%d=%d", i, i, i*i)
		fmt.Printf("%s\n", str)
	}
}
