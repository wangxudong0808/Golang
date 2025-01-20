package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("网线")
	// int8 -256 --- +256
	// unint8 0 --- +256
	// '' 单引号标识为字符，""双引号标识为字符串
	var y byte // byte=uint8 一个字节
	y = 'a'
	fmt.Println(y, reflect.TypeOf(y)) // 97  uint8
	fmt.Printf("%c\n", y)             //char ---- a
	fmt.Printf("%d\n", y)             //97
	fmt.Printf("%b\n", y)             //97的二进制
	fmt.Printf("%x\n", y)             //97的十六进制

	//字符串的存储原理
	//字符串的空间是地址+长度，在创建一个字符串的时候，会在内存里有一个[5]byte数组。
	s := "hello"
	t := s[2:3] //长度为1
	fmt.Println(s)
	fmt.Println(t) //  l 和s引用的一样的数组的地址空间

	// rune数据类型
	var z rune // rune=int32 4个字节
	z = '网'
	fmt.Println(z) //编码之后：32593
	fmt.Printf("%c\n", z)
	fmt.Println(string(z)) //转换为符号打印

}
