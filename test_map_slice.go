package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name, gender, age, mail, select_number, index_number, floor, index_name string
	var sys_test = []map[string]string{} //{map[string]string{"客户姓名": name,"客户性别": gender,"客户年龄": age,"客户邮箱": mail}}
	for true {
		fmt.Println("--------------------------客户信息管理系统-------------------------")
		fmt.Println("1、添加客户", "2、查看客户", "3、修改客户信息", "4、删除客户", "5、退出", "6、打印所有用户信息")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Print("请输入你的选择[1-5]: ")
		fmt.Scan(&select_number)
		switch select_number {
		case "1":
			{
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
					var xinxi_bucket = map[string]string{"客户姓名:": name, "客户性别:": gender, "客户年龄:": age, "客户邮箱:": mail}
					sys_test = append(sys_test, xinxi_bucket)
					fmt.Println("是否返回上一层[y/n]")
					fmt.Scan(&floor)
					if floor == "y" {
						break
					} else if floor == "n" {
						continue
					}
				}
			}
		case "2":
			{
				for true {
					fmt.Println("-------------------查询客户详细信息----------------------")
					fmt.Printf("请输入你要查询用户的索引:")
					fmt.Scan(&index_number)
					for index_number_1, key_value := range sys_test {
						index_number_2 := strconv.Itoa(index_number_1)
						if index_number_2 == index_number {
							for key_uu, value_uu := range key_value {
								fmt.Println(key_uu, value_uu)
							}
						}
					}
					fmt.Println("是否返回上一层[y/n]")
					fmt.Scan(&floor)
					if floor == "y" {
						break
					} else if floor == "n" {
						continue
					}
				}
			}
		case "3":
			{
				for true {
					fmt.Println("-------------------更新客户信息----------------------")
					fmt.Printf("请输入你更新客户信息的名称:")
					fmt.Scan(&index_name)
					for _, key_value := range sys_test {
						if key_value["客户姓名:"] == index_name {
							fmt.Print("姓名:")
							fmt.Scan(&name)
							key_value["客户姓名:"] = name
							fmt.Print("性别:")
							fmt.Scan(&gender)
							key_value["客户性别:"] = gender
							fmt.Print("年龄:")
							fmt.Scan(&age)
							key_value["客户年龄:"] = age
							fmt.Print("邮箱:")
							fmt.Scan(&mail)
							key_value["客户邮箱:"] = mail
						}
					}
					fmt.Println("是否返回上一层[y/n]")
					fmt.Scan(&floor)
					if floor == "y" {
						break
					} else if floor == "n" {
						continue
					}
				}
			}
		case "4":
			{
				for true {
					fmt.Println("-------------------删除客户信息----------------------")
					fmt.Printf("请输入你要删除客户的名称:")
					fmt.Scan(&index_name)
					for index_value, key_value := range sys_test {
						if key_value["客户姓名:"] == index_name {
							sys_test = append(sys_test[:index_value], sys_test[index_value+1:]...)
						}
					}
					fmt.Println("是否返回上一层[y/n]")
					fmt.Scan(&floor)
					if floor == "y" {
						break
					} else if floor == "n" {
						continue
					}
				}
			}
		case "5":
			{
				break
			}
		case "6":
			fmt.Println("打印所有用户信息")
			for _, key_value_78 := range sys_test {
				for key_value_121, value_121 := range key_value_78 {
					fmt.Println(key_value_121, value_121)
				}
				fmt.Println("-------------------------------------------")
			}
		}
	}
}
