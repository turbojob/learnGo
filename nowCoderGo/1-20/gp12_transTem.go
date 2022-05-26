package main

func temperature(f float64) float64 {
	// write code here
	//1、摄氏温度（C）与华氏温度（F）的换算式是：

	//C = 5×（F- 32）/9，F = 9×C /5+32
	return 5 * (f - 32) / 9
}
