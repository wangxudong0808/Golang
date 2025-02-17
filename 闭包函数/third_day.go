package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func gongNeng() {
	fmt.Println("功能")
}
func boyDay(function_test1 func()) func() { //大函数叫装饰函数
	var count = 0
	return func() {
		function_test1()
		count++
		funcName := runtime.FuncForPC(reflect.ValueOf(function_test1).Pointer()).Name()
		fmt.Printf("%s函数调用功能%d次\n", funcName, count)
	}
}

func main() {
	test_function := boyDay(gongNeng)
	test_function()
	test_function()
	test_function()

}
