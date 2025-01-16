package main

import "fmt"

func main() {
	//map的存储原理
	var info = make(map[string]string)
	info["name"] = "wxd"
	info["gender"] = "男"
	fmt.Println(info)
	info2 := info
	info["class"] = "数学"
	fmt.Println(info2)

	//map的增删改查
	value, exist := info["ggg"]
	fmt.Println(value, exist)
	if exist {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	//删除
	delete(info, "class")
	for key_value, value_number := range info {
		fmt.Println(key_value, value_number)
	}
	//map的嵌套
	var slice_test = []map[string]string{}
	y := map[string]string{"name": "wang"}
	s := append(slice_test, y)
	fmt.Println(s)
	//map嵌套map
	var a = map[string]map[string]string{}
	b := map[string]string{"name": "waawaw"}
	c := map[string]string{"age": "77"}
	a["student"] = b
	a["jokne"] = c
	fmt.Println(a["student"]["name"])

}
