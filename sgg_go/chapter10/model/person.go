package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		(*p).age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return (*p).age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		(*p).sal = sal
	} else {
		fmt.Println("工资范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return (*p).sal
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Println(stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (p *Pupil) Testing() {
	fmt.Println("小学生正在考试")
}

func (g *Graduate) Testing() {
	fmt.Println("大学生正在考试")
}
