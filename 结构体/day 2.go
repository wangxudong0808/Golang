package main

import "fmt"

type Student_day2 struct { //结构体是值类型，不利用地址去转存
	sid int //Student的成员对象
	//t1, t2 int
	name   string
	age    int
	course []string
	tt     map[string]string
}

func main() {
	var s Student_day2 //实例化结构体的时候才会去开辟空间,s是Student_day2的实例对象
	fmt.Println(s)
	s.sid = 132342
	s.name = "Wangxudong" //字符串的底层是数组，占两个值，第一个是地址，第二个是长度，根据地址和长度打印整个字符
	fmt.Println(s)
	fmt.Println(s.name)
	//切片声明完之后切片没有开辟数组
	s.course = make([]string, 3)
	s.course = []string{"数学", "英语", "语文"}
	fmt.Println(s)
	fmt.Printf("s的地址%p\n", &s)
	fmt.Printf("s.sid的地址%p\n", &s.sid)
	fmt.Printf("s.name的地址%p\n", &s.name)
	fmt.Printf("s.age的地址%p\n", &s.age)
}
