package main

import (
	"encoding/json"
	"os"
)

func main() {
	// 序列化：结构化的数据转换为字符串的方式
	// 反序列化：将字符串转换为结构化的数据

	// 将结构化的数据转换为json
	s := map[string]string{"name": "王旭东", "age": "20"}
	date, _ := json.Marshal(s) // 返回一个[]byte,字符切片
	//fmt.Println(reflect.TypeOf(date))  //uint 8
	//fmt.Println(string(date))
	//ioutil.WriteFile("data_json", date, 0666)
	file_test, _ := os.OpenFile("data_json", os.O_WRONLY|os.O_APPEND, 0666)
	file_test.Write(date)
	file_test.Write([]byte("\n"))

	// 将json数据转换为结构化数据
	//data, _ := os.OpenFile("data_json",os.O_RDONLY,0666)
	//data, _ := ioutil.ReadFile("data_json")
	//var data_test_number1 = make(map[string]string)
	//json.Unmarshal(data, &data_test_number1) // 将data的输入传到data_test_number1的空间中
	//fmt.Println(reflect.TypeOf(data_test_number1), data_test_number1)
}
