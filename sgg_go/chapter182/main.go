package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---Start---")
	fmt.Println(s)
	fmt.Println("---End---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("No struct")
		return
	}

	num := val.NumField()

	for i := 0; i < num; i++ {
		fmt.Println(i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println(i, tagVal)
		}
	}

	// numOfField := val.NumField()
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())

}

func main() {
	var a Monster = Monster{
		Name:  "asgh",
		Age:   400,
		Score: 38.3,
	}
	TestStruct(a)
}
