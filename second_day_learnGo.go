package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// new函数，引用地址如果没有赋值，则不会开辟空间 ---> 适用于指针
	var u *int
	u = new(int)
	fmt.Println(u)
	*u = 80
	fmt.Println(*u)
	// make函数 -- > 适用于切片
	//var s []int //申明一个切片，没有默认值
	// s[0] = 1 //会报错，因为没有可以引用的数组。
	// 初始化空间

	// 定义一个数组：5：长度，10：容量  还有5个空间可以使用。S再开辟一个空间指向这个数组
	var s = make([]int, 5, 10)
	fmt.Println(s, len(s), cap(s))
	s[0] = 100
	fmt.Printf("修改之后：%d\n", s)

	//扩容 append函数
	var new_slice []int
	new_slice_1 := append(new_slice, 1, 2, 3, 4)
	fmt.Println(new_slice_1)
	new_slice_2 := append(new_slice_1, 1000, 2000)
	fmt.Println(new_slice_2)

	var t = []int{44, 55}
	new_slice_2 = append(new_slice_2, t...) //将一个切片追加到另一个切片中
	fmt.Println(new_slice_2)

	var slice_first = make([]int, 3, 7)
	slice_first_1 := append(slice_first, 100, 200, 300)
	fmt.Println(slice_first_1)

	// append重点！
	a1 := []int{11, 22, 33}
	fmt.Println(len(a1), cap(a1))
	c := append(a1, 44) // 11 22 33 44
	// append扩容之后的容量是之前的2倍，然后将数据拷贝到现有的扩容空间内。返回一个新的切片，新的切片是对原有切片的引用。
	// 扩容之后和原有数组没有关系啦
	a1[0] = 100
	c[0] = 200
	fmt.Println(a1)                // 100 22 33
	fmt.Println(c, len(c), cap(c)) // 200 22 33 44    4 6

	// make 提前申请容量
	a := make([]int, 3, 10)
	b := append(a, 3, 5)
	fmt.Println(a, len(a), cap(a)) // 000 3 10
	fmt.Println(b, len(b), cap(b)) // 00035 5 10
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	//切片面试题
	a_db := [4]int{1, 2, 3, 4}
	ss := a_db[0:2]                               // 1 2   2  4
	ss2 := ss                                     // 1 2
	ss3 := append(append(append(ss, 55), 66), 77) // 1 2 55 66   1 2 55 66 77  0 0 0
	ss[0] = 100
	fmt.Println(a_db) // 100 2 55 66
	fmt.Println(ss)   // 100 2
	fmt.Println(ss2)  // 100 2
	fmt.Println(ss3)  // 1 2 55 66 77
	fmt.Println(&a_db[0])
	fmt.Println(unsafe.Sizeof(a_db[0]), reflect.TypeOf(a_db[0]))
	var a_a int32 = 1
	fmt.Println(unsafe.Sizeof(a_a), reflect.TypeOf(a_a))
	fmt.Println(&ss[1])
	fmt.Println(&ss2[1])
	fmt.Println(&ss3[1])
}
