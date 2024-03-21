package main

import (
	"fmt"
	"sgg_go/chapter02"
	"strconv"
	"unsafe"
)

var n1, n2 = 100, 100

func main() {
	// 全局变量
	fmt.Println(n1, n2)

	// 单变量声明
	var i int
	fmt.Println((i))
	var num = 10.10
	fmt.Println((num))
	num2 := 1000
	fmt.Println((num2))

	// 多变量声明
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)
	var n4, n5, n6 = 100, "fdsfds", 56.26
	fmt.Println(n4, n5, n6)
	n7, n8, n9 := 100, "fdsfds", 56.26
	fmt.Println(n7, n8, n9)

	fmt.Println("dsfds"+"fdsdf", 4+5)

	var a int = 89
	fmt.Println(a)

	var c byte = 98
	fmt.Println(c)
	fmt.Printf("%T", c)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(c))

	var d float32 = -123.000901
	fmt.Println(d)

	var e rune = '北'
	var f byte = 'f'
	fmt.Println(e)
	fmt.Printf("%c,%c", e, f)
	fmt.Println()

	var g = 'a' + 10
	fmt.Println(g)

	var h = false
	fmt.Println(h, unsafe.Sizeof(h))

	var address string = "fsdfds fds fds\n fds f" // 不可以操作内部的字符，字符串是不可变的
	var vary string = `fdsfd\n\\n`
	fmt.Println(address)
	fmt.Println(vary)

	var a4 float64 = 2.13456789
	b := float32(a4)
	fmt.Println(b)

	var num1 int = 99
	var numm2 float64 = 23.456
	var num3 bool = false
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%v", num1)
	a1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(str, a1)
	str = fmt.Sprintf("%v", numm2)
	a2, _ := strconv.ParseFloat(str, 64)
	fmt.Println(str, a2)
	str = fmt.Sprintf("%t", num3)
	a3, _ := strconv.ParseBool(str)
	fmt.Println(str, a3)
	str = fmt.Sprintf("%c", myChar)
	fmt.Println(str)

	// str = strconv.FormatInt(int64(num1), 10)
	// fmt.Println(str)
	// str = strconv.FormatFloat(numm2, 'f', 10, 64)
	// fmt.Println(str)
	// str = strconv.FormatBool(num3)
	// fmt.Println(str)

	var j int = 10
	var ptr *int = &j
	fmt.Println(*ptr)

	var zznum int = 10
	fmt.Println(&zznum)
	ptr2 := &zznum
	*ptr2 = 20
	fmt.Println(*ptr2)
	fmt.Println(zznum)

	// fmt.Println(chapter02.i)
	fmt.Println(chapter02.I)
}
