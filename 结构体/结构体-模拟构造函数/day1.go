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

func main() {
	//
	teacher := map[string]string{
		"name": "dashabi",
		"age":  "56",
	}
	s1 := NewStudent("wangxudong", 23, teacher, []string{"english"})
	fmt.Println(s1.teacher["name"])
	fmt.Println(s1.class)
}
