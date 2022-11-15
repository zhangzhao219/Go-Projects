package service

import (
	"Go-Projects/Customer-Management-System/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 完成对Customer的操作，包括增删改查
type CustomerService struct {
	// 存储当前的客户
	Customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	CustomerNum int
}

// 初始化CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{} // 初始化
	customerService.CustomerNum = 0
	return customerService
}

func (cs *CustomerService) ShowCustomerSlice() []model.Customer {
	return cs.Customers
}

func (cs *CustomerService) AddCustomer(customer model.Customer) bool {
	cs.Customers = append(cs.Customers, customer)
	cs.CustomerNum += 1
	return true
}

func (cs *CustomerService) DeleteCustomer(id int) bool {
	for index, cus := range cs.Customers {
		if cus.Id == id {
			cs.Customers = append(cs.Customers[:index], cs.Customers[index+1:]...)
			cs.CustomerNum -= 1
			return true
		}
	}
	return false
}

func (cs *CustomerService) ChangeCustomer(id int) bool {

	reader := bufio.NewReader(os.Stdin) // 标准输入输出

	for index, cus := range cs.Customers {
		if cus.Id == id {

			fmt.Printf("请输入修改的姓名(%v)：", cus.Name)
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if len(name) != 0 {
				cs.Customers[index].Name = name
			}

			fmt.Printf("请输入修改的性别(%v)：", cus.Gender)
			gender, _ := reader.ReadString('\n')
			gender = strings.TrimSpace(gender)
			if len(gender) != 0 {
				cs.Customers[index].Gender = gender
			}

			fmt.Printf("请输入修改的年龄(%v)：", cus.Age)
			age, _ := reader.ReadString('\n')
			age = strings.TrimSpace(age)
			if len(age) != 0 {
				t, _ := strconv.ParseInt(age, 10, 64)
				cs.Customers[index].Age = int(t)
			}

			fmt.Printf("请输入修改的电话(%v)：", cus.Phone)
			phone, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phone)
			if len(phone) != 0 {
				cs.Customers[index].Phone = phone
			}

			fmt.Printf("请输入修改的电子邮件(%v)：", cus.Email)
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			if len(email) != 0 {
				cs.Customers[index].Email = email
			}

			return true
		}
	}
	return false
}
