package main

import (
	"fmt"
	"reflect"
)

func main() {
	// new函数，引用地址如果没有赋值，则不会开辟空间
	//var u *int
	//u = new(int)
	//fmt.Println(u)
	//*u = 80
	//fmt.Println(*u)

	//数组
	//数组必须限制长度，类型必须统一
	//列表的话，开辟的空间里，存在一个地址，和nex，nex是指定这个列表里下一个地址的指针
	var name [3]int //开辟了三个空间，name变量指向这个数组空间，默认值都是0
	//索引的概念
	fmt.Println(name)
	for i := 0; i <= 2; i++ {
		name[i] = i
		fmt.Println(name[i])
	}
	//申明并赋值
	var names = [3]string{"wan", "xu", "dong"}
	for i := 0; i <= 2; i++ {
		fmt.Print(names[i])
	}
	//省略长度赋值
	var name_s = [...]string{"w", "w", "e", "i"} //加...就是忽略长度
	for i := 0; i <= len(name_s)-1; i++ {
		fmt.Println(name_s[i])
	}
	//索引赋值
	var name_ss = [...]int{0: 1, 1: 2, 4: 5} //未赋值的位置为默认值
	for i := 0; i <= len(name_ss)-1; i++ {
		fmt.Print(name_ss[i])
		if name_ss[i] == 2 {
			name_ss[i] = 10
			fmt.Printf("遇到2啦修改为%d\n", name_ss[i])
		} else {
			fmt.Println("no")
		}
		fmt.Println(reflect.TypeOf(name_ss)) //[5]int 数组的数据类型里有长度显示
	}
	// 数组的长度不可变、有序性、一致性（数据类型）,切片的话长度可变
	// 切片操作,从数组里切片的话，会得到一个切片值
	var add = [...]string{"wang", "wang_second", "dosec"}
	s := add[1:3]                     // 3-1就是切出来的元素个数
	fmt.Println(s, reflect.TypeOf(s)) // [wang_second dosec] []string 切片的数据类型里没有长度

	//遍历数组
	for a, b := range add { //range函数会将数组中的索引赋值给a,索引对应的值赋值到b
		fmt.Println(a, b)

	}

	//切片：数据类型

}
