package main

import "fmt"

type Student_day4 struct { //结构体是值类型，不利用地址去转存
	sid int //Student的成员对象
	//t1, t2 int
	name   string
	age    int
	course []string
	tt     map[string]string
}

func initSid(s *Student_day4) {
	fmt.Println(s)
	s.sid = 99 // == (*s).sid = 99
}

func initSid_2(p *Student_day4) {
	fmt.Println(*p)
	(*p).sid = 56
}

func main() {
	//声明并赋值
	var s = Student_day4{name: "wangxud", sid: 100, course: make([]string, 5)}
	//fmt.Println(s)

	var s1 = s //值拷贝
	s1.name = "niubi"
	//fmt.Printf("s1:=%s  \n", s1.name)
	//fmt.Printf("s:=%s  ", s1.name)
	initSid(&s)
	fmt.Printf("-----s:=%d\n,地址为%p\n", s.sid, &s.sid)
	initSid_2(&s)
	fmt.Printf("----initSid_2之后,s:=%d\n,地址为%p\n", s.sid, &s.sid)

	var s2 = &s // == 	var s2 = &Student_day4{name: "wangxud", sid: 100, course: make([]string, 5)}
	initSid(s2)
	fmt.Printf("s2的值为%d\n,地址为%p\n", s.sid, &s.sid)
}
