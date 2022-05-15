package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read() {
	//文件不大的时候 一次性读取全部到内存
	slice, err := ioutil.ReadFile("d://a.txt")
	if err != nil {
		fmt.Println("Error")
	}
	//
	fmt.Printf("%s", slice)
	fmt.Println("%v", string(slice))
}

func readBuffer() {
	file, err := os.Open("d://a.txt")
	if err != nil {
		return
	}
	defer file.Close()
	bufferReader := bufio.NewReader(file)

	for {
		str, e := bufferReader.ReadString('\n')

		fmt.Print(str)

		if e == io.EOF {
			break
		}
		//bufferReader.ReadLine()
	}

}

func write() {
	file, _ := os.OpenFile("d://a.txt", os.O_WRONLY, 0)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("zb12213")
	write.Flush()
}
func main() {
	write()
}
