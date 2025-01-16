package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// map 映射类型  ，通过hash映射的方式，底层的存储方式是数组的方式
	var students map[string]string // [括号里是key的类型]value的类型
	fmt.Println(students, reflect.TypeOf(students), len(students))

	//声明并初始化
	var stus = map[string]string{
		"name": "wangxudong",
		"age":  "20",
		"from": "china",
	}
	fmt.Printf("姓名:%s\n年龄:%s\n长度:%d\nmap:%s\n", stus["name"], stus["age"], len(stus), stus)
	fmt.Printf("%p\n", &stus)
	//写入一个key-value
	stus["gender"] = "man"
	fmt.Println(stus)
	//删除一个key-value，内置函数delete
	delete(stus, "from")
	fmt.Println(stus)

	// 基于make函数声明初始化,引用类型的数据可以先使用make在内存中声明一块空间
	// var stus2 map[string]string // map引用类型，未赋值之前不会开辟空间
	//  stus2["map"] = "rain" error---> panic: assignment to entry in nil map
	var stus2 = make(map[string]string)      // 动态扩容机制，默认加1；切片是2倍扩容 new返回一个指针地址。
	var stus3 = make(map[string]interface{}) // 加上空接口，支持所有类型
	fmt.Println(stus2)
	stus3["name"] = "wxd"
	stus3["age"] = 21
	stus3["gender"] = "男"
	delete(stus3, "age")
	fmt.Println(stus3)

	var number_data = []int{1, 2, 3, 4, 5}
	for _, data_number := range number_data {
		fmt.Printf("数据是:%d\n", data_number)
	}
	//for i := 1; i < 10; i = i + 2 {
	//	fmt.Print(i)
	//}
	for key, _ := range stus3 {
		if key == "age" {
			stus3["age"] = 33
		}
		fmt.Println(stus3[key])
	}

	//map的扩展
	var m = make(map[string][]string)
	m["Beijing"] = []string{"朝阳", "昌平", "海淀"}
	m["Neimeng"] = []string{"呼市", "集宁", "包头"}
	// 查询北京的第二个地区
	fmt.Println(m["Beijing"][1], reflect.TypeOf(m["Beijing"]))
	// 遍历每一个省份对应的城市名
	for key_number1, value_number1 := range m {
		for _, value_slice := range value_number1 {
			fmt.Println(key_number1, value_slice)
		}
		fmt.Println("----------------")
	}

	//练习
	students_status := []map[string]string{{"name": "yuan", "age": "23"}, {"name": "wxd", "age": "27"}, {"name": "dd", "age": "11"}}
	second_students := students_status[1]
	//fmt.Println(reflect.TypeOf(second_students)) //map类型
	fmt.Printf("姓名：%s\n年龄：%s\n", second_students["name"], second_students["age"])
	for _, slice_number := range students_status {
		//fmt.Println(reflect.TypeOf(slice_number)) map类型
		fmt.Printf("姓名：%s\n年龄：%s\n", slice_number["name"], slice_number["age"])
		fmt.Println("--------------------")
		//for _, value_number2 := range slice_number {
		//	fmt.Println(value_number2)
		//	//fmt.Println(reflect.TypeOf(value_number2))
		//	//fmt.Printf("姓名：%s\n年龄：%s\n", key_number2, value_number2)
		//}
	}
	// 添加一个新学生的信息到students_status
	var name, age string
	fmt.Println("请输入姓名和年龄")
	fmt.Scan(&name, &age)
	new_student := map[string]string{"name": name, "age": age}
	students_status = append(students_status, new_student)
	fmt.Println(students_status)

	//删除name=dd的map
	for slice_number, slice_value := range students_status {
		if slice_value["name"] == "dd" {
			students_status = append(students_status[0:slice_number], students_status[slice_number+1:]...)
		}
	}
	//students_status = append(students_status[0:2], students_status[3:]...)
	//students_status = append(students_status[0:2], students_status[3:]...)
	fmt.Println(students_status)

	// 将姓名为wxd的学生的年龄自加一岁
	for _, slice_value := range students_status {
		if slice_value["name"] == "wxd" {
			slice_int, _ := strconv.Atoi(slice_value["age"])
			//fmt.Println(slice_int)
			slice_int++
			//fmt.Println(reflect.TypeOf(slice_int), slice_int)
			slice_value["age"] = strconv.Itoa(slice_int)
		}
	}
	fmt.Println(students_status)
	// 根据age的大小重新排序
	status_age := []map[string]int{} //{map[string]int{"age": 23}, map[string]int{"age": 21}, map[string]int{"age": 55}}
	fmt.Println(status_age)

	//数组和map相比较，数组查询只能遍历查询，map（hash查询）可以直接根据key进行查询，时间复杂度ON来说，map的查询速度更快，空间换时间
	//
}
