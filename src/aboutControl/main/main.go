package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func giveANumber() int {
	return -1
}
func ifTest() {
	if num := giveANumber(); num < 0 {
		fmt.Println(num, "is Neg")
	} else if num < 10 {
		fmt.Println(num, "<10")
	} else {
		fmt.Println(num, "other case")
	}

}
func timeAndSwitch() {
	sec := time.Now().Unix()
	day := time.Now().Weekday()
	fmt.Println(day)
	rand.Seed(sec)

	var i = 0
	switch i {
	case 0:
		fmt.Println("isZero")
	case 1:
		fmt.Println("isOne")
	case 2:
		fmt.Println("isTwo")
	default:
		fmt.Println("default")
	}
	fmt.Println("Run!")
}
func switchAndFallThrough() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
func forAndNoWhile() {
	i := 10
	for i != 5 {
		i = int(rand.Int31n(14))
		fmt.Println(i, "不是5")
	}

}
func deferTest() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
	for i := 5; i <= 8; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

}
func deferAndPanicAndRecover() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newfile.Sync()
}
func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
func testPanic() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")

}
func main() {
	testPanic()
}
