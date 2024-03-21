package main

import "fmt"

func main() {
	// 都是整数，除法后去掉小数部分保留整数部分
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	fmt.Println(10 / 4.0)

	// a % b = a - a / b * b

	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	num := 97
	fmt.Println(num/7, num%7)

	// 键盘输入
	var name string
	var age byte
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Println(name, age)
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println(name, age)
}
