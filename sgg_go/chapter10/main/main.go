package main

import (
	"fmt"
	"sgg_go/chapter10/model"
)

func main() {
	p1 := model.NewPerson("face")
	fmt.Println(p1.GetAge())
	fmt.Println(p1.GetSal())
	p1.SetAge(12)
	p1.SetSal(30000.00)
	fmt.Println(p1.GetAge())
	fmt.Println(p1.GetSal())

	pupil := &model.Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 8
	pupil.Testing()
	pupil.SetScore(80)
	pupil.ShowInfo()

	graduate := &model.Graduate{}
	graduate.Student.Name = "Mary"
	graduate.Student.Age = 8
	graduate.Testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
