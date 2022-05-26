package main

func temperature(t float64) (ans string) {
	// write code here
	ans = ""
	defer func() {
		err := recover()
		if err != nil {
			ans = "体温异常"
		}
	}()
	if t > 37.5 {
		panic("体温异常  弹出警告！")
	}

	return ans

}
