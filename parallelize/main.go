package main

import (
	"fmt"
	"sync"
	"time"
)

var cond *sync.Cond
var wg sync.WaitGroup

// var ok bool

func waitGroupWaiting(i int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("waitGroup", i)
}

func condWaiting(i int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	fmt.Println(i, "正在等待")
	// for !ok {
	cond.Wait()
	// }
	fmt.Println("cond", i)
}

func main() {
	// Waitgroup一等多
	fmt.Println("开始进行Waitgroup测试")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go waitGroupWaiting(i)
	}
	wg.Wait()
	fmt.Println("Waitgroup测试完成")

	// Cond 多等一
	fmt.Println("开始进行Cond测试")
	cond = sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go condWaiting(i)
	}
	// ok = true
	time.Sleep(time.Second * 2)
	fmt.Println("Cond唤醒")
	cond.Broadcast()
	time.Sleep(time.Second * 5)
	fmt.Println("Cond测试完成")
	fmt.Println("程序运行结束！")
}
