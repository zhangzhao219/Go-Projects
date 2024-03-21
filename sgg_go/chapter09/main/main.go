package main

import (
	"fmt"
	"sgg_go/chapter09/model"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	hobby string
}

type Person struct {
	name string
	age  int
	ptr  *int
	map1 map[string]string
}

func (p Person) speak() {
	fmt.Println(p.name, "是一个好人")
}

func (p Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) jisuan2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) getSum(a, b int) int {
	return a + b
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.hobby = "fdss"
	fmt.Println(cat1)
	// var cat2 Cat
	var p1 Person
	p1.map1 = make(map[string]string)
	p1.map1["er"] = "sfds"
	var a int = 3
	var b int = 5
	p1.ptr = &a
	p2 := &p1
	(*p2).ptr = &b
	fmt.Println(p1)
	fmt.Println(*p2)

	var cat2 *Cat = new(Cat)
	(*cat2).Name = "fdsfdsfds"
	fmt.Println(*cat2)

	var cat3 *Cat = &Cat{"fdsfdsfds", 12, "fds", "erwe"}
	fmt.Println(*cat3)
	pn := Person{"fdsfds", 45, nil, nil}
	fmt.Println(pn.getSum(1, 2))
	pn.jisuan()
	pn.jisuan2(5)
	pn.speak()

	var stu = model.NewStudent("sdfd", 45.32)
	fmt.Println((*stu).ReturnScore())

}
