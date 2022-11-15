package model

import "fmt"

// 定义Customer结构体，表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式返回Customer的结构体，在CustomerService里面使用
// 感觉就是新建一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (cu *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", cu.Id, cu.Name, cu.Gender, cu.Age, cu.Phone, cu.Email)
}
