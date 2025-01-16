package main

import (
	"fmt"
	"reflect"
)

// 定义一个整型的指针变量
func function_Test1(x *int) {
	fmt.Println(x)
	*x++ //根据地址修改这个地址中x变量的值
}

func slcie_point(slice_p []int) { //函数接收到的值是数组的索引地址、长度、容量
	fmt.Println(slice_p)
	//slice_p[0] = 1323 //直接修改引用的数组值
	slice_p = append(slice_p, 5)
	slice_p[0] = 9898
	fmt.Printf("slice_point函数中slice_p的值为:%v\n,长度为%d\n,容量为%d\n", slice_p, len(slice_p), cap(slice_p))
}

func foo() {
	var bar_1 = func(int, int) {}
	var bar = func() {}
	fmt.Println(reflect.TypeOf(bar))   // func()类型
	fmt.Println(reflect.TypeOf(bar_1)) // func(int)
}

//匿名函数

func main() {
	// 传递的参数是实参，函数中申明的是形参
	//函数的值拷贝,地址拷贝
	x := 20
	fmt.Println(&x)
	function_Test1(&x) // 将x的地址传递到函数里
	fmt.Println(x)

	//--------
	var s = []int{1, 2, 3}
	slcie_point(s) //将一个切片的值传递到这个函数
	fmt.Println(s)

	//匿名函数,直接调用
	var t = func() {
		fmt.Println("匿名函数")
	}
	fmt.Println(t)
	t()
	t()
	t()
	//或者
	(func(number1, number2 int) {
		fmt.Println("直接调用，不需要另开辟空间存放在一个变量里")
		sum := number1 + number2
		fmt.Println(sum)
	})(10, 2) // 传递的是实参

	foo() //输出foo函数中bar变量的值，值是一个匿名函数

}
