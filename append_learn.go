package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s[0])
	s = append(s, 4)
	fmt.Printf("%p\n", &s) //地址没变，但是切片的值变啦

	// 向开头插入值或者切片
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(append([]int{0}, a...)) //向切片追加切片
	//任意位置插入
	fmt.Println(append(a[:2], append([]int{100}, a[2:]...)...)) // 1 2 100 3 4 5
	//删除任意位置
	fmt.Println(append(a[:2], a[3:]...))
	// --------------------切片
	var aar = [5]int{1, 2, 3, 4, 5}
	s1 := aar[0:3] // 1 2 3
	s2 := aar[2:5] // 3 4 5
	s3 := s2[0:2]  // 3 4
	println(s1)
	println(s2)
	println(s3)
}
