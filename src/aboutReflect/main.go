package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 200

	val := reflect.ValueOf(&num)
	(val).Elem().SetInt(900)
	fmt.Println(num)
}
