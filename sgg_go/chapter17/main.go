package main

import (
	"fmt"
	"time"
)

func outputHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello, World", i)
		time.Sleep(time.Second)
	}
}

var myMap = make(map[int]int, 10)

func testTan(n int, intChan1, intChan2 chan int) {
	for k := 1; k <= n; k++ {
		res := 1
		for i := 1; i <= n; i++ {
			res *= i
		}
		intChan1 <- k
		intChan2 <- res
	}
	close(intChan1)
	close(intChan2)
}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println(len(intChan))
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		fmt.Println(len(intChan), len(exitChan), v, ok)
		if !ok {
			break
		}
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	// go outputHello()
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Hello, Go", i)
	// 	time.Sleep(time.Second)
	// }
	// num := runtime.NumCPU()
	// fmt.Println(num)

	intChan1 := make(chan int, 20)
	intChan2 := make(chan int, 20)

	go testTan(20, intChan1, intChan2)

	for {
		v1, ok1 := <-intChan1
		v2, ok2 := <-intChan2
		if !ok1 && !ok2 {
			break
		}
		myMap[v1] = v2
	}

	for i, v := range myMap {
		fmt.Println(i, v)
	}

	// intChan := make(chan int, 50)
	// exitChan := make(chan bool, 1)

	// go readData(intChan, exitChan)
	// go writeData(intChan)

	// for {
	// 	_, ok := <-exitChan
	// 	fmt.Println(ok)
	// 	if !ok {
	// 		break
	// 	}
	// }

}
