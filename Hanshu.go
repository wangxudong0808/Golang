package main

import (
	"fmt"
	"reflect"
)

// 函数传参
// n 为形式参数（）没有确定值的
func sumNumber(n int) {
	fmt.Println(n)
}

func sumDoubleNumber(x, y int) {
	var sum = 0
	for i := x; i <= y; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

// 不限制传的参数,不定长参数
// 是一个int参数的切片
func sumDonotLength(nums ...int) {
	fmt.Println(reflect.TypeOf(nums))
	var sum = 0
	for _, v := range nums {
		sum = sum + v
	}
	fmt.Println(sum)
}

// 函数的返回值,括号外的int是返回值的数据类型
// 如果没有返回值的话，在调用之后不能将值赋给单独变量
func returnValue(number ...int) int {
	sum := 0
	for _, value := range number {
		sum = sum + value
	}
	return sum
}

func loginUser(use string, pwd string) (bool, string) {
	if use == "root" && pwd == "wxd0808" {
		return true, use
	} else {
		return false, ""
	}

}

// 返回值命名
func add2(s ...int) (z int) {
	for _, v := range s {
		z = z + v
	}
	return // 不返回值的话，默认返回z,如果z在函数中无值，则z默认为0
}

func selectNumber(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 变量的作用域：先从函数里找变量，再从gloab里找，然后从内键里找，都找不到的话就报错
// 函数拥有开辟作用域的能力
// 全局变量只能使用var进行声明，局部的变量可以使用:=声明
// if 和 for 都会开辟作用域
func test() {
	var x = 100
	fmt.Println(x)
	fmt.Println(&x)
}

var x = 1

func main() {
	//sumNumber(12)
	//sumDoubleNumber(1, 100) // 实际参数
	//sumDonotLength(1, 2, 3, 4, 5, 6, 7)
	//fmt.Println(returnValue(1, 2, 3, 4, 5, 6, 7))

	var username, password string
	fmt.Println("请输入用户名:")
	fmt.Scanln(&username)
	fmt.Println("请输入密码")
	fmt.Scanln(&password)
	status, user := loginUser(username, password)
	if status {
		fmt.Printf("登录成功，用户为%s\n", user)
	} else {
		fmt.Println("登录失败，用户名或密码错误")
	}

	//返回值命名
	//fmt.Println(add2(1, 2, 3))
	//test()
	//fmt.Println(x)
	//fmt.Println(&x)
}
