package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// {"客户姓名": name,"客户性别": gender,"客户年龄": age,"客户邮箱": mail}}
var sys_test = []map[string]string{}
var select_number_bool, name, gender, mail, age string

func addCluster() { //添加用户
	for true {
		fmt.Println("---------------添加客户开始-----------------------")
		fmt.Print("姓名:")
		fmt.Scan(&name)
		fmt.Print("性别:")
		fmt.Scan(&gender)
		fmt.Print("年龄:")
		fmt.Scan(&age)
		fmt.Print("邮箱:")
		fmt.Scan(&mail)

		var xinxi_bucket = map[string]string{"客户姓名": name, "客户性别": gender, "客户年龄": age, "客户邮箱": mail}
		sys_test = append(sys_test, xinxi_bucket)
		for index_value, slice_value := range sys_test {
			if slice_value["客户姓名"] == name {
				fmt.Printf("本次添加的用户的序号为%d\n", index_value+1)
				break
			}
		}
		fmt.Print("是否返回上一层[y/n]：")
		fmt.Scan(&select_number_bool)
		if select_number_bool == "y" {
			break
		} else if select_number_bool == "n" {
			continue
		} else {
			fmt.Println("请重新输入你的选择{y/n}")
		}
	}
} // 添加用户

func judgeInput(number_input int) bool {
	if number_input >= 0 && number_input < len(sys_test) {
		return true
	} else {
		return false
	}
} //判断用户输入的index是否合规

func select_user(index_user_data int) {
	for true {
		if index_user_data == 0 {
			fmt.Println("请输入你要查询的用户序号")
			fmt.Scan(&index_user_data)
		}
		if judgeInput(index_user_data-1) == true {
			fmt.Printf("序号为%d  姓名:%s  性别:%s  年龄%s  邮箱:%s\n",
				index_user_data,
				sys_test[index_user_data-1]["客户姓名"],
				sys_test[index_user_data-1]["客户性别"],
				sys_test[index_user_data-1]["客户年龄"],
				sys_test[index_user_data-1]["客户邮箱"])
		} else {
			fmt.Println("您输入的序号不存在", "请重新输入")
			index_user_data = 0
			continue
		}
		fmt.Print("是否返回上一层[y/n]：")
		fmt.Scan(&select_number_bool)
		if select_number_bool == "y" {
			fmt.Println("sdgfhj")
			break
		} else if select_number_bool == "n" {
			index_user_data = 0
			continue
		} else {
			fmt.Println("请重新输入你的选择{y/n}")
		}
	}
} //根据用户的序号查询并返回用户信息

func watchCluster() {
	var index_user int
	select_user(index_user)

} //菜单2查看用户信息

func getMapinsys_test(index int) map[string]string {
	var map_value map[string]string
	for index_value, value_slice := range sys_test {
		if index_value == index-1 {
			map_value = value_slice
		}
	}
	return map_value
} //根据index获取map数据

func update_user() {
	var update_index int
	fmt.Println("请输入要修改的用户序号")
	fmt.Scan(&update_index)
	//fmt.Println("用户现有信息如下:")
	//select_user(update_index)
	fmt.Println("请输入修改信息")
	fmt.Print("姓名:")
	fmt.Scan(&name)
	fmt.Print("性别:")
	fmt.Scan(&gender)
	fmt.Print("年龄:")
	fmt.Scan(&age)
	fmt.Print("邮箱:")
	fmt.Scan(&mail)
	update_user_status := getMapinsys_test(update_index)
	update_user_status["客户姓名"] = name
	update_user_status["客户性别"] = gender
	update_user_status["客户年龄"] = age
	update_user_status["客户邮箱"] = mail
	fmt.Println("修改后信息如下:")
	select_user(update_index)
} //更新用户的信息

func show_user() {
	for index_slice, value_number := range sys_test {
		fmt.Printf("序号为%d  姓名:%s  性别:%s  年龄%s  邮箱:%s\n",
			index_slice+1,
			value_number["客户姓名"],
			value_number["客户性别"],
			value_number["客户年龄"],
			value_number["客户邮箱"])
	}
} //展示所有用户的数据

func delete_user() {
	fmt.Println("所有用户信息如下")
	var delete_index int
	show_user()
	fmt.Println("请输入要删除的用户序号")
	fmt.Scan(&delete_index)
	if delete_index == 1 {
		sys_test = []map[string]string{}
	} else {
		sys_test = append(sys_test[:delete_index-1], sys_test[delete_index:]...)
	}
	fmt.Println("删除之后的用户信息")
	show_user()
} //删除指定序号的用户

// {"客户姓名": name,"客户性别": gender,"客户年龄": age,"客户邮箱": mail}}
// []map[string]string{}
func saveUser_file() {
	for _, value_number := range sys_test {
		fmt.Println(value_number)
		byte_data, _ := json.Marshal(value_number)
		//fmt.Println(byte_data, reflect.TypeOf(byte_data))
		//ioutil.WriteFile("User_file", byte_data, 0666)
		file_test, _ := os.OpenFile("User_file", os.O_CREATE|os.O_WRONLY, 0666)
		file_test.Write(byte_data)
		file_test.Write([]byte("\n"))
	}
}

func main() {

	for true {
		fmt.Printf("%s\n", `
---------------客户信息管理系统---------------------
   1、添加客户
   2、查看客户
   3、更新客户
   4、删除客户
   5、退出并保存
   6、展示所有用户
--------------------------------------------------
	`)
		stop_lit := false
		var select_Number int
		fmt.Print("请输入你的选择:[1-5]")
		fmt.Scan(&select_Number)
		switch select_Number {
		case 1:
			addCluster()
		case 2:
			watchCluster()
		case 3:
			update_user()
		case 4:
			delete_user()
		case 5:
			stop_lit = true
			saveUser_file()
		case 6:
			show_user()
		}
		if stop_lit == true {
			break
		}
	}

}
