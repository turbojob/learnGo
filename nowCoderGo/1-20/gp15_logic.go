package main

func logicalOperation(a bool, b bool) []bool {
	// write code here
	var ans []bool
	ans = append(ans, a && b)
	ans = append(ans, a || b)
	ans = append(ans, !a)
	ans = append(ans, !b)
	return ans

}
