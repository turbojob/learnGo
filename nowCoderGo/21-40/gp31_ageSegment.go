package main

func getAge(age int) string {
	// write code here
	var ans string
	switch {
	case age < 1 && age >= 0:
		ans = "婴儿"
	case age >= 1 && age <= 4:
		ans = "幼儿"
	case age >= 5 && age <= 11:
		ans = "儿童"
	case age >= 12 && age <= 18:
		ans = "少年"
	case age >= 19 && age <= 35:
		ans = "青年"
	case age >= 36 && age <= 59:
		ans = "中年"
	case age >= 60:
		ans = "老年"
	}
	return ans
}
