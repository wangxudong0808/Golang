package main

import (
	"fmt"
)

// {"客户姓名": name,"客户性别": gender,"客户年龄": age,"客户邮箱": mail}}
var sys_test = []map[string]string{}
var select_number_bool string

func addCluster() {
	var name, gender, mail, age string
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
				fmt.Printf("本次添加的用户的序号为%d\n", index_value)
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
}

func judgeInput(number_input int) bool {
	if number_input >= 0 && number_input < len(sys_test) {
		return true
	} else {
		return false
	}
}

func watchCluster() {
	var index_user int
	for true {
		fmt.Println("请输入你要查询的用户序号")
		fmt.Scan(&index_user)
		if judgeInput(index_user) == true {
			fmt.Printf("序号为%d\n姓名:%s\n性别:%s\n年龄%s\n邮箱:%s\n",
				index_user,
				sys_test[index_user]["客户姓名"],
				sys_test[index_user]["客户性别"],
				sys_test[index_user]["客户年龄"],
				sys_test[index_user]["客户邮箱"])
		} else {
			fmt.Println("您输入的序号不存在", "请重新输入")
			continue
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
}

func main() {

	for true {
		fmt.Printf("%s\n", `
---------------客户信息管理系统---------------------
   1、添加客户
   2、查看客户
   3、更新客户
   4、删除客户
   5、退出
--------------------------------------------------
	`)
		var select_Number int
		fmt.Print("请输入你的选择:[1-5]")
		fmt.Scan(&select_Number)
		switch select_Number {
		case 1:
			addCluster()
		case 2:
			watchCluster()
		}

	}

}
