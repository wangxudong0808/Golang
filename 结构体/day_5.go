package main

import (
	"fmt"
	"reflect"
)

type Student_day5 struct { //结构体是值类型，不利用地址去转存
	sid int //Student的成员对象
	//t1, t2 int
	name   string
	age    int
	course []string
	tt     map[string]string
}

func initSidDay5(s *Student_day5) {
	s.sid = 123
}

func main() {
	//new开辟一块空间
	var s = new(Student_day5) // == &Student_day5{}
	fmt.Println(s)
	fmt.Println(s.name, &s.sid)
	initSidDay5(s)
	fmt.Println(s.sid, &s.sid)

	//new函数和make函数初始化的区别
	var s1 = make([]int, 3) //开辟空间，存放一个地址，一个长度，一个容量
	fmt.Println(reflect.TypeOf(s1), s1)

	var s2 = new([]int) //开辟一个空间，是返回的一个地址,不管new什么都是一个单一的地址
	fmt.Println(reflect.TypeOf(s2), s2)
}
