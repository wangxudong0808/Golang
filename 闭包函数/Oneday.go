package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := func(x int) int {
		fmt.Println("test")
		x = 100
		return x
	}
	//
	var name_data int
	fmt.Print(b, reflect.TypeOf(b(name_data)))
	fmt.Print(b, reflect.TypeOf(b(name_data)), b(name_data))
}
