package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 切片的底层存储
	// 切片其实就是动态的数组，数组的长度是固定的，但是切片的长度是可伸缩的。
	var aar = [4]string{"name", "data", "niubi", "哈哈哈哈"}
	test_1 := aar[0:1] // 从数组中切出来的切片数据
	test_2 := aar[:]   // （数组中各个值的地址都是连续的）切片的空间里只存放了三个值：1、起始值的地址，2、长度，3、容量；
	//找到起始值的地址之后，根据长度（以test_1为例的话，长度为1-0=1），从现有的值取后几个值（因为数组的地址是有序的），容量则是从数组中切的值往后数(以test_1为例，则是4)
	fmt.Println(test_1, reflect.TypeOf(test_1))
	//切片的底层原理（数组中各个值的地址都是连续的）切片的空间里只存放了三个值：1、起始值的地址，2、长度，3、容量(数组的长度-切片的第一个索引地址)；
	test_1[0] = "wangxudong" //修改了aar数组的0索引的值
	fmt.Println(aar, reflect.TypeOf(aar))
	fmt.Println(test_2, reflect.TypeOf(test_2))
	//切片是对数组的引用
	new_test2 := test_2[1:] //从切片中取值
	fmt.Println(new_test2, reflect.TypeOf(new_test2))

	fmt.Println("------------------------------------------------------------------------------------------------")
	// 切片的空间里只存放了三个值：1、起始地址，2、长度，3、容量
	first_slice := aar[1:4]
	fmt.Printf("数组的值:%s\n", aar)
	fmt.Printf("_[1:4]_第一个切片的值为:%s\n", first_slice)
	second_slice := first_slice[1:3]
	fmt.Printf("从第一个切片_[1:3]_中切出来的值:%s\n", second_slice)
	fmt.Println("修改second_slick[1]的值为: Dosec ______________________________________________________")
	second_slice[1] = "Dosec"
	fmt.Printf("aar数组值:%s\nfirst_slice的值:%s\nsecond_slice的值:%s\n", aar, first_slice, second_slice)
	fmt.Printf("aar数组的容量：%d\nfirst_的容量：%d\nsecond_slick的容量：%d\n", cap(aar), cap(first_slice), cap(second_slice))

	fmt.Println("------------直接申明一个切片------------------------")
	// 直接申明一个切片
	var test_slice = []int{1, 2, 3, 4, 5} //底层存储的话，会先创建一个数组，然后test_slice是从 数组[:]引用数据
	fmt.Printf("切片test_slice的值为:%d\ntest_slice长度为:%d\n,test_slice容量为:%d\n", test_slice, len(test_slice), cap(test_slice))
	//3 6
	//5  6
	//4 5
	//5  5
	//6 6
	//1 4
	//var a = [...]int{1, 2, 3, 4, 5, 6}
	//a3 := a[1:5]
	//a6 := a3[1:2]
	//fmt.Println(len(a6), cap(a6))
	var a = []int{1, 2, 3, 4, 5, 6} // 定义一个切片
	b := a                          // 这一步的时候，b和a指向了同一个数组
	a[0] = 100                      //a将数组中0索引的值进行了修改，
	fmt.Println(b)                  //b拿到的是修改之后的值

}
