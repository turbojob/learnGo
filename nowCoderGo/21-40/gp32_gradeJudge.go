package main

func judgeScore(score int) string {
	// write code here
	var ans string
	switch {
	case score < 60:
		ans = "不及格"
	case score >= 60 && score < 80:
		ans = "中等"
	case score >= 80 && score < 90:
		ans = "良好"
	case score >= 90:
		ans = "优秀"
	}
	return ans
}
