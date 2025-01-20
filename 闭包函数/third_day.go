package main

import "fmt"

func gongNeng() {
	fmt.Println("功能")
}
func boyDay(function_test1 func()) func() { //大函数叫装饰函数
	var count = 0
	f := function_test1
	return func() {
		f()
		count++
		fmt.Printf("调用功能%d次\n", count)
	}
}

func main() {
	test_function := boyDay(gongNeng)
	test_function()
	test_function()
	test_function()

}
