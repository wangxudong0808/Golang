package main

import (
	"fmt"
	"reflect"
)

// 闭包函数
func getContent() func() {
	var i = 0
	//cntent := func() int { //content是闭包函数，谁引用了i谁就是闭包函数
	//	i++
	//	return i
	//}
	//return cntent()
	fmt.Println(&i)
	return func() {
		fmt.Println(&i, i)
		fmt.Println("----------")
		i++
		fmt.Println(i)
	}
}

// 闭包函数
func getContent_2(i int) func() {
	//cntent := func() int { //content是闭包函数，谁引用了i谁就是闭包函数
	//	i++
	//	return i
	//}
	//return cntent()
	return func() {
		fmt.Println("----------")
		i++
		fmt.Println(i)
	}
}

// 闭包的目的,做到数据隔离，不会对全局变量进行污染
// 引用了自由变量（外部非全局）的函数,这个外部自由变量会随着这个闭包函数一起去消亡。

func main() {
	//闭包，词法闭包，闭包函数：closure,引用了自由变量（外部非全局）的函数。
	sumNumber := getContent() // 对这个闭包函数进行了存放，保存了闭包函数和引用的变量
	fmt.Println(reflect.TypeOf(sumNumber))
	sumNumber()
	sumNumber()
	sumNumber()
	fmt.Println("-------------另一种方式--------------")

	sumNumber2 := getContent_2(1)
	sumNumber2()
	sumNumber2()
	sumNumber2()
	sumNumber3 := getContent_2(6)
	sumNumber3()
	sumNumber3()
	sumNumber3()

}
