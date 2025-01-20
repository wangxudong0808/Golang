package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func write_allfile() {
	string_file := "网上的看来封杀大家看法哈是否"
	ioutil.WriteFile(
		"测试1", []byte(string_file), 06666)
}
func main() {
	file, err := os.OpenFile("New_file", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	//O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	//// The remaining values may be or'ed in to control behavior.
	//O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	//O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	//O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	//O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	//O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	if err != nil {
		fmt.Println("打开文件失败")
	}
	//按字符串或者是按字节写操作
	file.WriteString("wangxudasssfd\n")
	// 创建一个字符串，将字符串变成字节串之后写入到文件中
	name_wirte := "wangxudongn王旭东"
	name_wirte1 := []byte(name_wirte)
	fmt.Println(name_wirte1, string(name_wirte1))
	file.Write(name_wirte1)
	//
	//// 缓冲写
	//write_file := bufio.NewWriter(file)
	//write_file.WriteString("网上的会发生尽快答复哈萨克积分")
	//write_file.Write([]byte{'w'}) // 将写的内容存放到缓存区
	//write_file.Flush()            //将缓存区的数据写道文件中

	//写整个文件
	//write_allfile()
}
