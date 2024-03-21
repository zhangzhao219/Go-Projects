package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)

	fmt.Println(rValue.Int() + 2)
}

type student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rValue := reflect.ValueOf(b)
	iV := rValue.Interface()
	fmt.Printf("%T", iV)
	stu, ok := iV.(*student)
	if ok {
		fmt.Println(stu.Name, stu.Age)
	}
	fmt.Println(rType.Kind(), rValue.Kind())
}

func reflect01(b interface{}) {
	rValue := reflect.ValueOf(b)
	rValue.Elem().SetInt(2)
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := &student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)
	num2 := 10
	reflect01(&num2)
	fmt.Println(num2)

}
