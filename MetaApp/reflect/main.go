package main

import (
	"fmt"
	"reflect"
)

type user struct {
	a int
	B int
	C string `json:"ccc"`
}

func main() {
	d := user{
		a: 1,
		B: 2,
		C: "fdsdfs",
	}
	typeUser := reflect.TypeOf(d)
	fieldNum := typeUser.NumField()
	fmt.Println(fieldNum)
	for i := 0; i < fieldNum; i++ {
		field := reflect.TypeOf(d).Field(i)
		fmt.Printf("%d %s offset %d anonymous %t type %s exported %t json tag %s\n", i,
			field.Name,            // 变量名称
			field.Offset,          // 相对于结构体首地址的内存偏移量，string类型会占据16个字节
			field.Anonymous,       // 是否为匿名成员
			field.Type,            // 数据类型，reflect.Type类型
			field.IsExported(),    // 包外是否可见（即是否以大写字母开头）
			field.Tag.Get("json"), // 获取成员变量后面``里面定义的tag
		)
	}
	//可以通过FieldByName获取Field
	if nameField, ok := typeUser.FieldByName("C"); ok {
		fmt.Printf("Name is exported %t\n", nameField.IsExported())
	}
	// 也可以根据FieldByIndex获取Field
	indexField := typeUser.FieldByIndex([]int{1}) //参数是个slice，因为有struct嵌套的情况
	fmt.Printf("third field name %s\n", indexField.Name)
}
