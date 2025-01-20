package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readByte_file(file *os.File) {
	date := make([]byte, 9)
	file.Read(date)
	fmt.Println(date, string(date))
}

func readByLine_file(file_line *os.File) {
	reader := bufio.NewReader(file_line)
	lineContent, _ := reader.ReadString('\n')
	fmt.Print(lineContent) //读到换行符就停止

	//一行一行读
	for true {
		lineContent_second, _, err := reader.ReadLine()
		//fmt.Println(lineContent_second)
		fmt.Println(string(lineContent_second)) //读一行
		if err != nil {
			break
		}
	}
	fmt.Println("--------------------------------")
	//一个字符一个字符读
	//for true {
	//	lineContent, err := reader.ReadString('\n')
	//	fmt.Print(string(lineContent))
	//	if err == io.EOF {
	//		break
	//	}
	//}
}

func readAll_file(file_name string) {
	//读取整个文件：只适合小文件
	contest, _ := ioutil.ReadFile(file_name)
	fmt.Print(string(contest))
}

func main() {
	//文件读操作
	// os.Open  文件句柄 file
	file, _ := os.Open("名称") // open函数返回 *File error
	// go的工作目录必须在同级目录下
	//if err != nil {
	//	fmt.Println("打开失败")
	//}
	//fmt.Println(file)

	//读文件，按字节读
	readByte_file(file) // 光标定位在了第三个汉字
	fmt.Println("----------")
	//读文件，按照行读
	readByLine_file(file) //从第三个汉字开始读
	//读取整个文件
	readAll_file("名称")
}
