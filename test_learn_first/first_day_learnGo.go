package main

func main() {
	//var age = "123"
	////var age_int, err = strconv.Atoi(age)
	//var age_int, _ = strconv.Atoi(age)
	////fmt.Println("err", err)
	//fmt.Println(age_int, reflect.TypeOf(age_int))
	//age_string := strconv.Itoa(age_int + 10)
	//fmt.Println(age_string, reflect.TypeOf(age_string))
	//
	//var age_old = "33岁"
	//var age_new, err = strconv.Atoi(age_old)
	//fmt.Println(age_new, err)
	//
	////浮点转字符 strconv Parse系列函数
	////将字符串转换为整型
	//ret, _ := strconv.ParseInt("28", 10, 8) // 8是bit位，10是进制。
	//fmt.Println(ret, reflect.TypeOf(ret))
	////将字符串转换为浮点型
	//ret2, _ := strconv.ParseFloat("3.1415926", 64) //不管限制为多少，最后返回的是64位
	//fmt.Println(ret2, reflect.TypeOf(ret2))
	//fmt.Println(ret2 + 1)
	//将字符串转换为布尔值
	//ret3, _ := strconv.ParseBool("0")
	//fmt.Println(ret3, reflect.TypeOf(ret3))
	//ret4, _ := strconv.ParseBool("1")
	//fmt.Println(ret4, reflect.TypeOf(ret3))
	//ret5, _ := strconv.ParseBool("true")
	//fmt.Println(ret5, reflect.TypeOf(ret3))
	//ret6, err := strconv.ParseBool("F")
	//fmt.Println(ret6, err, reflect.TypeOf(ret3))

	//var name_set, test_number int
	//fmt.Scanf("%d+%d", &name_set, &test_number)
	//fmt.Println(name_set, test_number)

	//var first_name string
	//fmt.Println("请输入你的英文名字")
	//fmt.Scan(&first_name)
	//first_name_new := strings.ToUpper(first_name) //将输入的字符全部强制转换为大写字母
	//fmt.Println(first_name_new)
	//fmt.Println(strings.HasPrefix(first_name_new, "a")) //判断是否是a开头
	//切片
	//fmt.Println("请输入你的出生日期")
	//var name string
	//fmt.Scan(&name)
	//birthday := strings.Split(name, "-")
	//fmt.Println(birthday)
	//fmt.Printf("年%s,月%s,日%s", birthday[0], birthday[1], birthday[2])

	//var a, b int
	//fmt.Println("请按指定格式输入")
	//fmt.Scanf("%d + %d", &a, &b)
	//fmt.Println(a + b)

	//流程控制语句
	//var name = "root"
	//if strings.HasPrefix(name, "r") || name == "root" {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}

	//var name, password string
	//var count int = 1
	//fmt.Println("请输入用户名和密码")
	//fmt.Scan(&name, &password)
	//if name == "root" && password == "12345" {
	//	fmt.Println("登录成功")
	//} else {
	//	fmt.Println("登录失败请重新输入")
	//	count++
	//}

	//var score int
	//fmt.Println("请输入你的成绩")
	//fmt.Scan(&score)
	//if score < 0 || score > 100 {
	//	fmt.Println("输入的数字不合法")
	//} else if score > 90 {
	//	fmt.Println("考的还行")
	//} else if score > 60 {
	//	fmt.Println("好好好")
	//} else {
	//	fmt.Println("傻逼")
	//}
	//var week int
	//fmt.Scan(&week)
	//switch week { // switch使用一张表来进行哈希定位，空间换速度
	//case 0:
	//	fmt.Println("星期日")
	//case 1:
	//	fmt.Println("星期1")
	//case 2:
	//	fmt.Println("星期2")
	//case 3:
	//	fmt.Println("星期3")
	//case 4:
	//	fmt.Println("星期5")
	//case 5:
	//	fmt.Println("星期5")
	//case 6:
	//	fmt.Println("星期6")
	//default:
	//	fmt.Println("傻逼")
	//}

	//var year string
	//fmt.Println("请输入你的出生日期，按格式：月-日")
	//fmt.Scan(&year)
	//list_year := strings.Split(year, "-")
	////fmt.Println(list_year)
	//month := list_year[0]
	//day := list_year[1]
	//int_month, _ := strconv.ParseFloat(month, 64)
	//int_day, _ := strconv.ParseFloat(day, 64)
	//switch int_month {
	//case 1:
	//	if int_day >= 1 && int_day <= 19 {
	//		fmt.Println("魔蝎")
	//	} else {
	//		fmt.Println("水平")
	//	}
	//case 2:
	//	if int_day >= 1 && int_day <= 18 {
	//		fmt.Println("水平")
	//	} else {
	//		fmt.Println("双鱼")
	//	}
	//case 3:
	//	if int_day >= 1 && int_day <= 20 {
	//		fmt.Println("双鱼")
	//	} else {
	//		fmt.Println("白羊")
	//	}
	//case 4:
	//	if int_day >= 1 && int_day <= 19 {
	//		fmt.Println("白羊")
	//	} else {
	//		fmt.Println("金牛")
	//	}
	//case 5:
	//	if int_day >= 1 && int_day <= 20 {
	//		fmt.Println("金牛")
	//	} else {
	//		fmt.Println("双子")
	//	}
	//case 6:
	//	if int_day >= 1 && int_day <= 22 {
	//		fmt.Println("双子")
	//	} else {
	//		fmt.Println("巨蟹")
	//	}
	//case 7:
	//	if int_day >= 1 && int_day <= 23 {
	//		fmt.Println("巨蟹")
	//	} else {
	//		fmt.Println("狮子")
	//	}
	//case 8:
	//	if int_day >= 1 && int_day <= 23 {
	//		fmt.Println("狮子")
	//	} else {
	//		fmt.Println("处女")
	//	}
	//case 9:
	//	if int_day >= 1 && int_day <= 23 {
	//		fmt.Println("处女")
	//	} else {
	//		fmt.Println("天秤")
	//	}
	//case 10:
	//	if int_day >= 1 && int_day <= 23 {
	//		fmt.Println("天秤")
	//	} else {
	//		fmt.Println("天蝎")
	//	}
	//case 11:
	//	if int_day >= 1 && int_day <= 22 {
	//		fmt.Println("天蝎")
	//	} else {
	//		fmt.Println("射手")
	//	}
	//case 12:
	//	if int_day >= 1 && int_day <= 21 {
	//		fmt.Println("射手")
	//	} else {
	//		fmt.Println("摩羯")
	//	}
	//
	//}
	//for true {
	//	fmt.Println()
	//}

	//打断点的方式
	//var count = 100
	//for count > 0 {
	//	fmt.Println(count)
	//	count--
	//}
	//for i := 0; i < 10; i++ {
	//	if i == 6 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//break 跳出整个循环
	//continue跳出当前循环

	//var sal = 2045.115342
	//fmt.Printf("工资:%.2f\n", sal)
	//fmt.Printf("类型:%T\n", sal) //%t 是浮点型
	//// 给任意类型占位：
	//fmt.Printf("任意类型打印:%#v\n", sal) //打印原始值
	//// %p指针类型
	//fmt.Printf("指针:%p", &sal) //&取值符，地址16进制表达
	for i := 1; i <= 10; i++ {

	}
}
