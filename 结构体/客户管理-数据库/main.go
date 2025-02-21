package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Mainfase struct to hold user info
type Mainfase struct {
	Name   string
	Gender string
	Age    int
	Email  string
}

func index_user() int {
	var userID int
	fmt.Println("请输入用户Id")
	fmt.Scan(&userID)
	return userID
}

// NewMainfase constructor to create a new Mainfase
func NewMainfase(Name, Gender string, Age int, Email string) *Mainfase {
	return &Mainfase{
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Email:  Email,
	}
}

// Database connection string
const connStr = "postgres://postgres:36UdcuHtpZSfRj@192.168.3.130:30432/vludb?sslmode=disable"

// ConnectDB function to establish a connection to the PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AddUser function to insert a new user into the database
func AddUser(db *sql.DB, m *Mainfase) error {
	query := `INSERT INTO users (name, gender, age, email) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, m.Name, m.Gender, m.Age, m.Email)
	if err != nil {
		return err
	}
	return nil
}

// SelectUser function to query a user from the database by their ID
func SelectUser(db *sql.DB, userID int) (*Mainfase, error) {
	var m Mainfase
	query := `SELECT name, gender, age, email FROM users WHERE id = $1`
	row := db.QueryRow(query, userID)
	err := row.Scan(&m.Name, &m.Gender, &m.Age, &m.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with ID %d", userID)
		}
		return nil, err
	}
	return &m, nil
}

func AlterUser(db *sql.DB, userID int, Leixing map[string]interface{}) (*Mainfase, error) {
	var m Mainfase
	//fmt.Println(Leixing)
	for index, value := range Leixing {
		query := fmt.Sprintf(`UPDATE users SET %s = $2 WHERE id = $1`, index)
		_, err := db.Exec(query, userID, value)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no user found with ID %d", userID)
			}
			return nil, err
		}
	}
	//err := row.Scan(&m.Name, &m.Gender, &m.Age, &m.Email)
	return &m, nil
}

// Monologue prompts user input for creating a new user
func Monologue() *Mainfase {
	var Name, Gender, Email string
	var Age int
	fmt.Print("姓名:")
	fmt.Scan(&Name)
	fmt.Print("性别:")
	fmt.Scan(&Gender)
	fmt.Print("年龄:")
	fmt.Scan(&Age)
	fmt.Print("邮箱:")
	fmt.Scan(&Email)
	return NewMainfase(Name, Gender, Age, Email)
}

// Adduser_test handles the process of adding users to the database
func Adduser_test() {
	var select_number_bool string
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	for {
		fmt.Println("---------------添加客户开始-----------------------")
		returnMonologue := Monologue()
		err := AddUser(db, returnMonologue)
		if err != nil {
			fmt.Println("添加用户失败:", err)
		} else {
			fmt.Println("用户添加成功！")
		}
		// Ask if the user wants to go back
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

// Select_user handles the user query by ID from the database
func Select_user() {
	index_user_data := index_user()
	var select_number_bool string
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()
	for {
		if index_user_data == 0 {
			fmt.Println("请输入你要查询的用户序号")
			fmt.Scan(&index_user_data)
		}
		user, err := SelectUser(db, index_user_data)
		if err != nil {
			fmt.Println(err)
			index_user_data = 0
			continue
		}
		fmt.Printf("序号为%d  姓名:%s  性别:%s  年龄%d  邮箱:%s\n",
			index_user_data,
			user.Name,
			user.Gender,
			user.Age,
			user.Email)
		// Ask if the user wants to go back
		fmt.Print("是否返回上一层[y/n]：")
		fmt.Scan(&select_number_bool)
		if select_number_bool == "y" {
			break
		} else if select_number_bool == "n" {
			index_user_data = 0
			continue
		} else {
			fmt.Println("请重新输入你的选择{y/n}")
		}
	}
	//return index_user()
}
func Alter_user() {
	//var userID int
	var value, index, select_number_bool string
	for {
		//var Leixing map[string]interface{}
		//fmt.Println("请输入你要修改的用户ID")
		//fmt.Scan(&userID)
		userID := index_user()
		fmt.Println("请输入你要修改的属性")
		fmt.Scan(&index)
		fmt.Println("请输入你要修改的值")
		fmt.Scan(&value)
		Leixing := map[string]interface{}{index: value}
		db, err := ConnectDB()
		if err != nil {
			log.Fatal("数据库连接失败:", err)
		}
		defer db.Close()
		_, err1 := AlterUser(db, userID, Leixing)
		if err1 != nil {
			fmt.Println(err)
			//userID = 0
		}
		fmt.Print("是否返回上一层[y/n]：")
		fmt.Scan(&select_number_bool)
		if select_number_bool == "y" {
			break
		} else if select_number_bool == "n" {
			//userID = 0
			continue
		} else {
			fmt.Println("请重新输入你的选择{y/n}")
		}
	}
}

func main() {
	for {
		// Display menu
		fmt.Printf("%s\n", `---------------客户信息管理系统---------------------
   1、添加客户
   2、查看客户
   3、更新客户
   4、删除客户
   5、退出并保存
   6、展示所有用户
--------------------------------------------------`)

		var select_use int
		fmt.Println("请输入你的选择")
		fmt.Scan(&select_use)
		switch select_use {
		case 1:
			Adduser_test()
		case 2:
			Select_user()
		case 3:
			Alter_user()
		}
	}
}
