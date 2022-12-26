package main

import (
	"fmt"
	"math"
)

func judgeprime(n int, num chan int) {
	if n == 0 || n == 1 {
		return
	}
	maxnum := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxnum; i += 1 {
		if n%i == 0 {
			return
		}
	}
	num <- n
}

func judge(nmin, nmax int, num, exit chan int) {
	for i := nmin; i < nmax; i++ {
		judgeprime(i, num)
	}
	exit <- nmin
}

// func main() {
	num := make(chan int)
	exit := make(chan int)
	total := 8000
	for i := 0; i < 4; i++ {
		go judge(total/4*i, total/4*(i+1), num, exit)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exit
		}
		close(num)
	}()
	for v := range num {
		fmt.Printf("%d ", v)
	}
}
