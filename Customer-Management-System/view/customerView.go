package view

import (
	"Go-Projects/Customer-Management-System/model"
	"Go-Projects/Customer-Management-System/service"
	"fmt"
)

type CustomerView struct {
	key             string                   // 接收用户输入
	loop            bool                     // 表示是否循环的显示主菜单
	username        string                   // 用户的用户名
	password        string                   // 用户的密码
	customerService *service.CustomerService // 获取用户服务
}

func NewCustomerView() *CustomerView {
	return &CustomerView{
		key:             "",
		loop:            true,
		username:        "admin",
		password:        "password",
		customerService: service.NewCustomerService(),
	}
}

func (cv *CustomerView) login(username, password string) bool {
	if username == cv.username && password == cv.password {
		return true
	}
	return false
}

func (cv *CustomerView) addCustomer() {
	id := cv.customerService.CustomerNum + 1
	var name, gender, phone, email string
	var age int
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入性别：")
	fmt.Scanln(&gender)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入电话：")
	fmt.Scanln(&phone)
	fmt.Print("请输入电子邮件：")
	fmt.Scanln(&email)

	if cv.customerService.AddCustomer(*model.NewCustomer(id, name, gender, age, phone, email)) {
		fmt.Println("-------------------------添加成功-------------------------")
	} else {
		fmt.Println("-------------------------添加失败-------------------------")
	}
}

func (cv *CustomerView) changeCustomer() {
	var id int
	fmt.Print("请输入修改的ID号：")
	fmt.Scanln(&id)
	if cv.customerService.ChangeCustomer(id) {
		fmt.Println("-------------------------修改成功-------------------------")
	} else {
		fmt.Println("-------------------------添加失败-------------------------")
	}
}

func (cv *CustomerView) deleteCustomer() {
	var id int
	fmt.Print("请输入删除的ID号：")
	fmt.Scanln(&id)
	if cv.customerService.DeleteCustomer(id) {
		fmt.Println("-------------------------删除成功-------------------------")
	} else {
		fmt.Println("-------------------------删除失败-------------------------")
	}
}

func (cv *CustomerView) showCustomer() {
	if cv.customerService.CustomerNum == 0 {
		fmt.Println("没有客户！")
		return
	}
	fmt.Println("\n-------------------------客户列表-------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t电子邮件")
	for _, eachCustomer := range cv.customerService.ShowCustomerSlice() {
		fmt.Println(eachCustomer.GetInfo())
	}
}

func (cv *CustomerView) exit() {
	var choice byte
	for {
		fmt.Print("确定退出？(y/n)：")
		fmt.Scanf("%c\n", &choice)
		if choice == 'y' {
			cv.loop = false
			break
		} else if choice == 'n' {
			break
		} else {
			fmt.Println("输入有误！！请重新输入")
		}
	}
}

func (cv *CustomerView) MainMenu() {
	for {
		var username, password string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Print("请输入密码：")
		fmt.Scanln(&password)
		if cv.login(username, password) {
			break
		} else {
			fmt.Println("用户名或密码错误！")
		}
	}
	// 显示主菜单
	for cv.loop {
		fmt.Println("\n---------------------客户信息管理软件---------------------")
		fmt.Println("                         1 添加客户")
		fmt.Println("                         2 修改客户")
		fmt.Println("                         3 删除客户")
		fmt.Println("                         4 客户列表")
		fmt.Println("                         5 退    出")
		fmt.Print("请选择(1-5)：")
		// 接收用户的输入
		fmt.Scanln(&cv.key)
		// 对用户的输入进行判断
		switch cv.key {
		case "1":
			cv.addCustomer()
		case "2":
			cv.changeCustomer()
		case "3":
			cv.deleteCustomer()
		case "4":
			cv.showCustomer()
		case "5":
			cv.exit()
		default:
			fmt.Println("请输入正确的选项------")
		}
	}
}
