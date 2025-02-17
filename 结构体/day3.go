package main

import "fmt"

type Student_day3 struct { //结构体是值类型，不利用地址去转存
	sid int //Student的成员对象
	//t1, t2 int
	name   string
	age    int
	course []string
	tt     map[string]string
}

func main() {
	//声明并赋值
	var s = Student_day3{name: "wangxud", sid: 100, course: make([]string, 5)}
	fmt.Println(s)
	//按顺序赋值的话，所有的成员对象必须都赋值
	var s1 = Student_day3{1022, "wangxudong", 10, []string{"wang", "xu", "dong"}, make(map[string]string)}
	fmt.Println("-----------------------------")
	fmt.Println(s1)
}
