package main

import (
	"fmt"
	util "sgg_go/chapter06/utils"
	"strconv"
	"strings"
)

var Fun1 = func(n1, n2 int) int {
	return n1 - n2
}

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
}

func getSum(a, b int) int {
	return a + b
}

func sum(args ...float64) float64 {
	result := 0.0
	for i := 0; i < len(args); i++ {
		result += args[i]
	}
	return result
}

func swap(n1, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func defersum(a, b int) int {
	defer fmt.Println(a)
	defer fmt.Println(b)
	defer fmt.Println(a + b)
	a++
	b++
	return a + b
}

func main() {
	fmt.Println(util.Cal(1, 2, 3))
	a := getSum
	fmt.Println(a(2, 3))
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 8))
	n1 := 0
	n2 := 1
	fmt.Println(n1, n2)
	swap(&n1, &n2)
	fmt.Println(n1, n2)

	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	b := func(n1, n2 int) int {
		return n1 + n2
	}
	fmt.Println(res1)
	fmt.Println(b(10, 20))
	fmt.Println(Fun1(10, 20))
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("fdsfds"))

	fmt.Println(defersum(8, 1002))

	// str函数
	str := "范德萨范德萨hello"
	fmt.Println(len(str))
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i])
	}
	fmt.Println()
	n, err := strconv.Atoi("123")
	fmt.Println(n, err)
}
