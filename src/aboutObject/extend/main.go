package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (a Animal) ao() {
	fmt.Println("大声喊叫")
}
func (a Animal) show() {
	fmt.Println("体重是", a.Weight)
	fmt.Println("年龄是", a.Age)
}

type Cat struct {
	Animal
}

func (c Cat) scratch() {
	fmt.Println("挠人")
}

func main() {
	cat := Cat{}
	cat.Animal.Age = 18
	cat.Animal.Weight = 1002.12
	cat.ao()
	cat.show()
	cat.scratch()

}
