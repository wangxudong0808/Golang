package main

import "fmt"

type Student1 struct {
	name    string
	age     int
	teacher map[string]string
	class   []string
}

func NewStudent1(name string, age int, teacher map[string]string, class []string) Student1 {
	return Student1{
		name:    name,
		age:     age,
		teacher: teacher,
		class:   class,
	}
}

// 使用指针接收器
func (s *Student1) read1() {
	fmt.Printf("%s正在读书\n,这个的地址为%p\n", s.name, &s.name)
}

// 使用指针接收器
func (s *Student1) learn1(studentName *Student1) {
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
	s1 := NewStudent1("wangxudong", 23, teacher, []string{"english"})
	fmt.Println(s1.teacher["name"])
	fmt.Println(s1.class)
	s1.read1()     // 这里 s1 会自动转换为指针
	s1.learn1(&s1) // 这里 s1 会自动转换为指针
	fmt.Println(&s1.name)
}
