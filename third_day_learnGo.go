package main

import "fmt"

func main() {
	/*
		a := 20
		b := "王旭东"
		var a,b := 20,"王旭东"
	*/
	var a, b = 20, "王旭东"
	var name string = "如来神掌" // > var name = ""
	var (
		school = "农业大学"
		shiji  = "21世纪"
	)
	fmt.Println("sss")
	fmt.Printf("name=%d\nage=%s\nJineng=%s\nschool=%s\nshiji=%s\n", a, b, name, school, shiji)
}
