package main

import (
	"fmt"
)

func main() {
	// 字符串和字节串之间的转换
	var msg = "hello word" // 字符串
	var name = "王xdnb"
	//字符串转换为字节串
	fmt.Println([]byte(msg))
	fmt.Println([]byte(name)) //转为字节串
	fmt.Println([]rune(name)) //29579(unicode) 120 100 110 98 常用于求字符串里的字符长度

	//字节串转换为字符串
	test1 := []byte(name)
	test2 := []byte{231, 142, 139, 120, 100, 110, 98}
	fmt.Println(test2) //
	fmt.Println(test1)
	fmt.Println(string(test2))
	fmt.Println(string(test1))
}
