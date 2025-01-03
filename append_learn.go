package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	fmt.Printf("%p\n", &s)
	s = append(s, 4)
	fmt.Printf("%p\n", &s) //地址没变，但是切片的值变啦

	// 向开头插入值或者切片
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(append([]int{0}, a...)) //向切片追加切片
	//任意位置插入
	fmt.Println(append(a[:2], append([]int{100}, a[2:]...)...)) // 1 2 100 3 4 5
	//删除任意位置
	fmt.Println(append(a[:2], a[3:]...))
}
