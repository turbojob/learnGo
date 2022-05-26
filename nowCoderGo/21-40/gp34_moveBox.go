package main

//注意：若字符串中包含中文
//那么传统的遍历字符串的方式就是错误，因为传统的对字符串的遍历是按照字节来遍历
//在utf-8中，一个汉字是3字节，一个字母是1个字节。
//解决方法是将str转为[]rune切片：
func pushBox(forwards string) bool {
	// write code here
	var x int
	var y int
	length := len(forwards)
	for i := 0; i < length; i++ {
		cur := forwards[i]
		switch cur {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
