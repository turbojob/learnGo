package main

import "errors"

type error interface {
	Error() string
}

func Error() string {
	return "网络延迟"

}
func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		errors.New("网络延迟")
		return "网络延迟"
	}
	return " "
}
