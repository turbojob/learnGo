package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d://a.txt")
	if err != nil {
		fmt.Println("File open Error")
	}
	fmt.Printf("文件=%v", file)
	file.Close()
}
