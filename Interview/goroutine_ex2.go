package main

import "fmt"

func putNum(n int, putNumChan chan int) {
	for i := 1; i <= 2000; i++ {
		putNumChan <- i
	}
	close(putNumChan)
}

func dealNum(x int) int {
	res := 0
	for i := 1; i <= x; i++ {
		res += i
	}
	return res
}

func add(putNumChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		x, ok := <-putNumChan
		if !ok {
			exitChan <- true
			return
		}
		a := dealNum(x)
		m := make(map[int]int, 0)
		m[x] = a
		resChan <- m
	}
}

func output(resChan chan map[int]int, exitChan chan bool) {
	for i := 0; i < 8; i++ {
		<-exitChan
	}
	close(exitChan)
	close(resChan)
}

func main() {
	n := 2000
	putNumChan := make(chan int)
	go putNum(n, putNumChan)
	resChan := make(chan map[int]int)
	exitChan := make(chan bool)
	for i := 0; i < 8; i++ {
		go add(putNumChan, resChan, exitChan)
	}
	go output(resChan, exitChan)
	for v := range resChan {
		fmt.Println(v)
	}
}
