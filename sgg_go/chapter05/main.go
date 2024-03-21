package main

import (
	"fmt"
	"math"
)

func main() {
	if age := 20; age < 18 {
		fmt.Println("OK")
	} else {
		fmt.Println("not")
	}
	// 求 ax2+bx+c=0的根，判断有无根和几个根并输出
	var a float64
	var b float64
	var c float64
	fmt.Scanln(&a, &b, &c)

	if j := b*b - 4*a*c; j < 0 {
		fmt.Println("该方程没有根！")
	} else if j == 0 {
		fmt.Println("该方程只有一个根")
		fmt.Println((-b + math.Sqrt(j)) / (2.0 * a))
	} else {
		fmt.Println("该方程有两个根")
		fmt.Println((-b + math.Sqrt(j)) / (2.0 * a))
		fmt.Println((-b - math.Sqrt(j)) / (2.0 * a))
	}
	var day byte
	fmt.Scanf("%c", &day)
	switch day {
	case 'a':
		fmt.Println("1")
	case 'b':
		fmt.Println("2")
	case 65, 66, 67, 68:
		fmt.Println("0")
	default:
		fmt.Println("-1")
	}
	sentence2 := "hello北京"
	sentence := []rune(sentence2)
	for i := 0; i < len(sentence); i++ {
		fmt.Printf("%c ", sentence[i])
	}
	for index, str := range sentence {
		fmt.Printf("%d %c", index, str)
	}
	var ai int = 0
	for ai < 10 {
		fmt.Println("fdsfss")
		ai++
	}

	var num int = -1
	odd := 0
	even := 0
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		} else if num%2 == 1 {
			odd += 1
		} else {
			even += 1
		}
	}
	fmt.Println(odd, even)

}
