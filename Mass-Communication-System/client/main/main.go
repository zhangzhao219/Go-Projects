package main

import (
	"Go-Projects/Mass-Communication-System/client/process"
	"fmt"
)

// 定义两个变量，一个表示用户ID，一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {

	// 接收用户的选择
	var key int
	// 判断是否还能继续显示菜单
	var loop = true
	// 循环展示菜单
	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统---------------")
		fmt.Println("---------------   1 登录聊天室")
		fmt.Println("---------------    2 注册用户")
		fmt.Println("---------------    3 退出系统")
		fmt.Println("请选择(1-3): ")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)
			// 完成登录
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入注册用户的ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入注册用户的密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入注册用户的名称")
			fmt.Scanln(&userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")

		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
