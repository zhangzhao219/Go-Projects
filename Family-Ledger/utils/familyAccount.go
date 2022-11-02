package utils

import "fmt"

// 定义结构体
type FamilyAccount struct {
	// 用户名和密码
	username string
	password string
	// 声明一个变量保存用户的输入
	key string
	// 声明一个变量，控制是否退出for循环
	loop bool
	// 定义账户的初始值
	balance float64
	// 定义每次收支的金额和说明
	money float64
	note  string
	// 收支的详情使用字符串来记录
	// 当有记录时对这个字符串进行拼接
	details string
	// 判断当前是否有输入或者输出的记录
	flag bool
}

// 编写一个工厂模式的构造方法，返回结构体的指针
func NewFamilyAcount() *FamilyAccount {
	return &FamilyAccount{
		username: "admin",
		password: "password",
		key:      "",
		loop:     true,
		balance:  10000.0,
		money:    0.0,
		note:     "",
		details:  "收  支\t收支前账户余额\t收支金额\t收支后账户余额\t说  明\n",
		flag:     false,
	}
}

// 给结构体绑定相应的方法

// 将显示明细写成一个方法
func (fa *FamilyAccount) showDetails() {
	if !fa.flag {
		fmt.Println("当前没有任何收支记录！")
	} else {
		fmt.Println("---------------------当前收支明细记录---------------------")
		fmt.Println(fa.details)
	}
}

// 将登记收入写成一个方法
func (fa *FamilyAccount) income() {
	fmt.Println("-------------------------登记收入-------------------------")
	fmt.Print("本次收入金额：")
	fmt.Scanln(&fa.money)
	// 收入金额不能是负数
	if fa.money <= 0 {
		fmt.Println("收入金额应为正数！")
		return
	}
	fmt.Print("本次收入说明：")
	fmt.Scanln(&fa.note)
	fa.balance += fa.money
	fa.details += fmt.Sprintf("收  入\t%v\t%v\t%v\t%v\n", fa.balance-fa.money, fa.money, fa.balance, fa.note)
	fmt.Println("收入登记成功！")
	fa.flag = true
}

// 将登记支出写成一个方法
func (fa *FamilyAccount) pay() {
	fmt.Println("-------------------------登记支出-------------------------")
	fmt.Print("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("余额的金额不足！")
		return
	} else if fa.money <= 0 {
		fmt.Println("支出金额应为正数！")
	}
	fmt.Print("本次支出说明：")
	fmt.Scanln(&fa.note)
	fa.balance -= fa.money
	fa.details += fmt.Sprintf("收  入\t%v\t%v\t%v\t%v\n", fa.balance+fa.money, fa.money, fa.balance, fa.note)
	fmt.Println("支出登记成功！")
	fa.flag = true
}

// 将退出系统写成一个方法
func (fa *FamilyAccount) exit() {
	var choice byte
	for {
		fmt.Print("确定退出？(y/n)：")
		fmt.Scanf("%c\n", &choice)
		if choice == 'y' {
			fa.loop = false
			break
		} else if choice == 'n' {
			break
		} else {
			fmt.Println("输入有误！！请重新输入")
		}
	}
}

// 用户登录的功能
func (fa *FamilyAccount) login(username string, password string) bool {
	if (username == fa.username) && (password == fa.password) {
		return true
	}
	return false
}

// 显示主菜单
func (fa *FamilyAccount) MainMenu() {
	for {
		var username, password string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Print("请输入密码：")
		fmt.Scanln(&password)
		if fa.login(username, password) {
			break
		} else {
			fmt.Println("用户名或密码错误！")
		}
	}
	// 显示主菜单
	for fa.loop {
		fmt.Println("\n---------------------家庭收支记账软件---------------------")
		fmt.Println("                       1 收支明细")
		fmt.Println("                       2 登记收入")
		fmt.Println("                       3 登记输出")
		fmt.Println("                       4 退出软件")
		fmt.Print("请选择(1-4)：")
		// 接收用户的输入
		fmt.Scanln(&fa.key)
		// 对用户的输入进行判断
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		default:
			fmt.Println("请输入正确的选项------")
		}
	}
}
