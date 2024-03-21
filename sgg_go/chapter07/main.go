package main

import (
	"fmt"
	"math/rand"
)

func fbn(n int) []int {
	slice := make([]int, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}

func main() {
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 = [3]int{1, 2, 3}
	var array3 = [...]int{8, 9, 10}
	var array4 = [...]int{0: 1, 2: 5}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	var randarray [5]int
	for i := 0; i < 5; i++ {
		randarray[i] = rand.Int()
	}
	for _, value := range randarray {
		fmt.Println(value)
	}
	for _, value := range randarray {
		defer fmt.Println(value)
	}

	s2 := fbn(5)
	fmt.Println("fddfs")
	for _, value := range s2 {
		fmt.Printf("%d ", value)
	}
}
