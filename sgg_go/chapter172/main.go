package main

import "fmt"

func putNumber(intChan chan int) {
	for i := 1; i <= 20; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNumber(intChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= num; i++ {
			res += i
		}
		a := make(map[int]int)
		a[num] = res
		resChan <- a
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resChan := make(chan map[int]int, 2000)
	exitChan := make(chan bool, 8)

	go putNumber(intChan)
	for i := 0; i < 8; i++ {
		go primeNumber(intChan, resChan, exitChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		for i, v := range res {
			fmt.Println(i, v)
		}
	}

}
