package main

import "fmt"

func main() {
	// 声明一个变量保存用户的输入
	key := ""

	// 声明一个变量，控制是否退出for循环
	loop := true

	// 定义账户的初始值
	balance := 10000.0

	// 定义每次收支的金额和说明
	var money float64
	var note string

	// 收支的详情使用字符串来记录
	// 当有记录时对这个字符串进行拼接
	details := "收  支\t收支前账户余额\t收支金额\t收支后账户余额\t说  明\n"
	// 判断当前是否有输入或者输出的记录
	flag := false

	// 显示主菜单
	for loop {
		fmt.Println("\n---------------------家庭收支记账软件---------------------")
		fmt.Println("                       1 收支明细")
		fmt.Println("                       2 登记收入")
		fmt.Println("                       3 登记输出")
		fmt.Println("                       4 退出软件")
		fmt.Print("请选择(1-4)：")
		// 接收用户的输入
		fmt.Scanln(&key)
		// 对用户的输入进行判断
		switch key {
		case "1":
			if !flag {
				fmt.Println("当前没有任何收支记录！")
			} else {
				fmt.Println("---------------------当前收支明细记录---------------------")
				fmt.Println(details)
			}
		case "2":
			fmt.Println("-------------------------登记收入-------------------------")
			fmt.Print("本次收入金额：")
			fmt.Scanln(&money)
			// 收入金额不能是负数
			if money <= 0 {
				fmt.Println("收入金额应为正数！")
				break
			}
			fmt.Print("本次收入说明：")
			fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("收  入\t%v\t%v\t%v\t%v\n", balance-money, money, balance, note)
			fmt.Println("收入登记成功！")
			flag = true
		case "3":
			fmt.Println("-------------------------登记支出-------------------------")
			fmt.Print("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额的金额不足！")
				break
			} else if money <= 0 {
				fmt.Println("支出金额应为正数！")
			}
			fmt.Print("本次支出说明：")
			fmt.Scanln(&note)
			balance -= money
			details += fmt.Sprintf("收  入\t%v\t%v\t%v\t%v\n", balance+money, money, balance, note)
			fmt.Println("支出登记成功！")
			flag = true
		case "4":
			var choice byte
			for {
				fmt.Print("确定退出？(y/n)：")
				fmt.Scanf("%c\n", &choice)
				if choice == 'y' {
					loop = false
					break
				} else if choice == 'n' {
					break
				} else {
					fmt.Println("输入有误！！请重新输入")
				}
			}
		default:
			fmt.Println("请输入正确的选项------")
		}
	}
	fmt.Println("-------------------退出家庭收支记账软件-------------------")
}
