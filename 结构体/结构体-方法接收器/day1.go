package main

import "fmt"

type Student struct {
	name    string
	age     int
	teacher map[string]string
	class   []string
}

func NewStudent(name string, age int, teacher map[string]string, class []string) Student {
	return Student{
		name:    name,
		age:     age,
		teacher: teacher,
		class:   class,
	}
}

// Student类型的方法接收器
func (s Student) read() { //唯一的属于了Student类型,s为变量实例对象
	//补充动作
	fmt.Printf("%s正在读书\n,这个的地址为%p\n", s.name, &s.name)
}

// 方法接收器的声明
func (s Student) learn(studentName *Student) {
	fmt.Println("正在学习")
	studentName.name = "WXD"
	fmt.Printf("学习的地址是%p\n", &studentName.name)
	fmt.Println(&s.name)
}

func main() {

	teacher := map[string]string{
		"name": "dashabi",
		"age":  "56",
	}
	s1 := NewStudent("wangxudong", 23, teacher, []string{"english"})
	fmt.Println(s1.teacher["name"])
	fmt.Println(s1.class)
	s1.read()
	s1.learn(&s1)
	//fmt.Println(s1)
	fmt.Println(&s1.name)

	//为什么两次s.name的地址都不一样？
	//每次调用方法时，Go 会创建一个新的副本（即一个新的 Student 实例），而不是直接操作原始的 Student 实例。
	//如果想一样的话，可以只用指针接收器的方式，如day2.go文件中：

}
