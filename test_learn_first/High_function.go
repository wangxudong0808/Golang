package main

import (
	"fmt"
	"reflect"
	"time"
)

// 高阶函数：
// 以函数作为参数或者返回值是一个函数
func number() {
	fmt.Println("number")
	time.Sleep(time.Second * 1)
}

func function_a(number_1 func()) {
	fmt.Println(function_a, reflect.TypeOf(function_a))
	//function_a()
	number()
}

func time_unix_now(func_number1 func()) int64 {
	tim1 := time.Now().Unix()
	func_number1()
	tim2 := time.Now().Unix()
	return (tim2 - tim1)
}

func return_function() func() int {
	var test = func() int {
		//fmt.Println("new function")
		return 12
	}
	return test
}

func main() {
	//fmt.Println(number, reflect.TypeOf(number)) // 返回函数的地址、函数的类型
	//function_a(number)
	//
	//time_now := time.Now().Unix()
	//fmt.Printf("打印当前时间戳%d\n数据类型:%v\n ", time_now, reflect.TypeOf(time_now))
	////
	////time_1 := time.Now().Unix()
	////fmt.Println(time_1)
	////// time.Sleep(time.Second * 3) // 睡眠三秒
	////time_2 := time.Now().Unix()
	////fmt.Println(time_2)
	//
	//fmt.Printf("耗时：%d\n", time_unix_now(number))

	//函数作为返回值
	var f func() int
	f = return_function()
	fmt.Println(f())
}
