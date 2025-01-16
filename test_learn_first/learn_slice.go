package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		C是根本性的指针，对指针的操作及其灵活，会造成内存泄漏
		go的指针：取值和取址
	*/
	var x = 10 //在内存里开辟一个空间，存放了10，任何空间都对应一个地址
	// &变量，获取变量地址
	fmt.Printf("赋值之前%p\n", &x)
	x = 100
	fmt.Printf("赋值之后%p\n", &x)
	//u := &x
	var u *int //定义一个指针类型,*这个地址中存放的值的类型,p是一个整型指针
	u = &x     // u 是指针变量，对应的一个整型的指针类型，
	fmt.Println(u)
	fmt.Printf("赋值之后%p\n", &u)
	//取值操作
	fmt.Printf("*u的为:%d\n", *u, reflect.TypeOf(*u))
	//根据地址重新赋值
	*u = 456
	fmt.Println(x)

	//指针三部曲  获取地址(& 变量)、地址赋值(指针类型)、取值操作（*指针变量）
	//不同的变量引用同一块空间
	var a = 1
	var b *int
	b = &a
	var c *int
	c = b //拷贝的地址，引用拷贝
	*b = 100
	*c = 120
	fmt.Println(a, *b, *c)
}
