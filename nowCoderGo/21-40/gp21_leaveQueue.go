package main

func deleteElement(s []int, index int) []int {
	// write code here
	s1 := s[0 : index+1]
	s2 := s[index+1:]

	s1 = append(s1, s2...)
	return s1

}
