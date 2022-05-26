package main

func canConstruct(des string, src string) bool {
	// write code here
	have := make(map[byte]int)
	s := []byte(src)
	for _, i := range s {
		have[i]++
	}

	p := []byte(des)
	for _, i := range p {
		if have[i] == 0 {
			return false
		}
		have[i]--
	}
	return true
}
