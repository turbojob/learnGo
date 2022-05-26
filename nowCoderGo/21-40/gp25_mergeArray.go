package main

import "fmt"

func main() {
	s1 := []int{1, 2, 4, 0, 0, 0}
	s2 := []int{2, 5, 6}
	s3 := merge(s1, 3, s2, 3)
	fmt.Println(s3)
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	//类似 归并排序  从后面 开始判断  把大的 返回
	i := m + n - 1
	p1 := m - 1
	p2 := n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
			i--
		} else {
			nums1[i] = nums2[p2]
			p2--
			i--
		}
	}
	for p1 >= 0 {
		nums1[i] = nums1[p1]
		p1--
		i--
	}
	for p2 >= 0 {
		nums1[i] = nums2[p2]
		p2--
		i--
	}
	return nums1
}
