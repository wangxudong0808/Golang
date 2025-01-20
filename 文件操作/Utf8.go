package main

import "fmt"

func main() {
	var name rune
	// rune:对应uint32,四个字节，存放unicode编码表    byte:对应uint8,一个字节，存放Asc码编码表
	name = '苑' // unicode完全涵盖了所有的字符，
	fmt.Printf("unicode的十进制%d\n", name)
	fmt.Printf("unicode的十六进制%x\n", name)
	fmt.Printf("unicode的二进制%b\n", name) //1000001011010001

	//转成utf8: 动态伸缩的unicode

	// 三个字节的填充规则： 1110xxxx 10xxxxxx 10xxxxxx
	// unicode二进制:1000001011010001  ----> utf8:111001000 10001011 10010001,占三个字节
	// 最终的utf8值：1110010001000101110010001

	//简单的使用
	var name_first = "王abc"
	fmt.Println(name_first, len(name_first)) // len是字节  王：3个字节，xd:占两个字节
	for i := 0; i < len(name_first); i++ {   // 按字节遍历
		fmt.Println(name_first[i]) // 231 142 139 97 98 99,都是按照utf8编码存放的。字母的话utf8编码和ASCII码表一样
	}

	for i_index, v_value := range name_first { // 按照字符遍历
		fmt.Println(i_index, v_value /*unicode值*/, string(v_value))
		/*
			0 29579 王
			3 97 a
			4 98 b
			5 99 c
		*/
	} // utf8是针对unicode的编码规则

}
