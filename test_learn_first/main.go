package main

import (
	"fmt"
	"strings"
)

func main() {
	//var x = 10
	//var y = x
	//x = 20
	//fmt.Println(y)
	//fmt.Println(x)
	//// var a,_ = 1,2     _是匿名变量  匿名变量的作用，当函数固定返回两个值时，只取其中一个的话，将另一个定位为匿名变量即可
	//// ;和换行符 作为语句分隔符 var a = 1;var y = 100
	//// strings（类型）.方法(实参)
	//var s = "my name is \"wangxudong\""
	//fmt.Println(s)
	//N_name := `
	//1.shabi
	//2.lenqiu
	//`
	//fmt.Println(N_name)
	//fmt.Printf(strings.ToUpper(N_name))
	//var w = "wang xu dong"
	//fmt.Println(strings.HasPrefix(w, "wang")) //判断是否以wang这个字符串开头
	//fmt.Println(strings.Contains(w, "u"))     // 只要字符串里有，就会返回true
	//username := "   wang xu dong niubi  "
	////fmt.Println(strings.Trim(uername, " ")) //去除两端指定字符
	//username = strings.TrimSpace(username)
	//fmt.Println(username)
	//// index: 索引
	//index_username := strings.Index(username, "i")
	//fmt.Println(index_username) //如果出现-1则是没有找到这个字符

	// 从命令中取数据库用户名和密码
	//var cmd = "mysql .... -u root -p 123   "
	//index_user := strings.Index(cmd, "-u")
	//index_password := strings.Index(cmd, "-p")
	//fmt.Printf("用户名所在位置:%d\n密码所在位置%d\n", index_user, index_password)
	//index_user_name := cmd[index_user+2 : index_password]
	//index_user_password := cmd[index_password+2:]
	////fmt.Println(user_name, user_password)
	//user_name := strings.TrimSpace(index_user_name)
	//user_password := strings.TrimSpace(index_user_password)
	//fmt.Printf("User:%s\nPassword:%s", user_name, user_password)

	//分割 拼接
	name := "xu wang dong"
	nameSlice := strings.Split(name, " ") //指定分隔符
	first_name := nameSlice[1]
	second_name := nameSlice[0]
	third_name := nameSlice[2]
	user_name := first_name + second_name + third_name
	fmt.Println(user_name)
	//指定拼接符号
	var new_name = strings.Join(nameSlice, "-") // 将分割出来的字符，通过-进行拼接
	fmt.Println(new_name)
}
