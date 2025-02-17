package main

import "fmt"

func f2() *int {
	i := 5
	defer func() {
		i++
		fmt.Printf("%p\n", &i)
	}()
	fmt.Println(&i)
	return &i
}

func f3() (returt int) {
	defer func() {
		returt++
	}()
	return 5
}

func f4() (returt int) {
	defer func() {
		fmt.Println(returt)
		returt++
	}()
	return returt
}

func f5() (r int) { // r=0
	t := 5
	defer func() {
		t++
	}()
	return t
}
func f6() (r int) {
	fmt.Println(&r) //r的地址
	defer func(r int) {
		r++             // r=1
		fmt.Println(&r) //另一个r的地址
	}(r)
	return 5 //r=5 r=1 ret r   也就是5
}

func f7() (r int) {
	defer func(x int) {
		r = x + 1
	}(r) //defer注册时r=0
	return 5 // r=5 r=1 ret r=1
}

func main() {
	//f2()
	//fmt.Println(*f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
	//fmt.Println(f5()) // r=5 t=6 最后返回ret r
	fmt.Println(f7())
}
