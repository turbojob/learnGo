package main

import "fmt"

type animal interface {
	sleep()
	eat()
}
type tiger struct {
}

func (t tiger) sleep() {
	fmt.Println("sleep")
}
func (t tiger) eat() {
	fmt.Println("eat")
}
func main() {
	var t animal = tiger{}
	t.sleep()
	t.eat()

}
