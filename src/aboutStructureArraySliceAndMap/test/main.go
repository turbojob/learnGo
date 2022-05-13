package main

import "fmt"

func main() {
	fmt.Println(fib(1))
}
func fib(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	ans := []int{1, 1}
	for i := 3; i <= n; i++ {
		ans = append(ans, ans[len(ans)-1]+ans[len(ans)-2])
	}
	return ans

}
