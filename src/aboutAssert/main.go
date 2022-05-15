package main

import (
	"fmt"
)

type people interface {
	say()
}
type chinese struct {
	name string
}

func (c chinese) say() {
	fmt.Println("Chinese can eat")
}
func (c chinese) niu() {
	fmt.Println("Chinese could run")
}

type american struct {
	name string
}

func (c american) talk() {
	fmt.Println("am could talk")
}
func (c american) say() {
	fmt.Println("american can drink")
}
func process(p people) {

	//p 是否可以转成 chinese

	ch, ok := p.(chinese)
	if ok {
		ch.say()
		ch.niu()
	}

	ch1, ok1 := p.(american)
	if ok1 {
		ch1.say()
		ch1.talk()
	}
}

func main() {
	//多态 接口的时候  断言
	//判断类型
	c := chinese{name: "zb"}
	d := american{}
	process(c)
	process(d)
}
