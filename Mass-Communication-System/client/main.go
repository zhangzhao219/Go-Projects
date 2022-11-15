package main

import "fmt"

// 定义两个变量，一个表示用户ID，一个表示用户密码
var userId int
var userPwd string

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
		fmt.Println("请选择（1-3）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	if key == 1 {
		fmt.Println("请输入用户ID")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)
		// 先把登录的函数写在另外一个文件
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
