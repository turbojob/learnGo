package main

func convert(s []int) []int {
	// write code here
	length := len(s)
	for i := 0; i < length/2; i++ {
		//s[i], s[length-i-1] := s[length-i-1], s[i]
		tmp := s[i]
		s[i] = s[length-i-1]
		s[length-i-1] = tmp
	}
	return s
}
