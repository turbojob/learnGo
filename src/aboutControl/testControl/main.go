package main

import "fmt"

func FizzBuzz() {
	//编写 FizzBuzz 程序
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}
	}
}
func findprimes(number int) bool {
	for i := 2; i <= number-1; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number != 1 {
		return true
	} else {
		return false
	}
}

func getPrime() {
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}
}
func main() {
	getPrime()
}
