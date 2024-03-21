package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// const defaultWorkers = 10000000

// type (
// 	// FilterFunc defines the method to filter a Stream.
// 	FilterFunc func(item *AdsItem) bool
// 	WalkFunc   func(item *AdsItem, pipe chan<- *AdsItem)
// 	Stream     struct {
// 		source <-chan *AdsItem
// 	}
// )

// // drain drains the given channel.
// func drain(channel <-chan *AdsItem) {
// 	for range channel {
// 	}
// 	// count11 := 0
// 	// for item := range channel {
// 	// 	count11 += 1
// 	// 	if item.A == 1 || item.A == 2 || item.A == 3 || item.A == 4 {
// 	// 		fmt.Println(item.A, item.B)
// 	// 	}
// 	// 	if item.B == 5 || item.B == 6 || item.B == 7 {
// 	// 		fmt.Println(item.A, item.B)
// 	// 	}
// 	// }
// 	// fmt.Println(count11)
// }

// // Done waits all upstreaming operations to be done.
// func (s Stream) Done() {
// 	drain(s.source)
// }

// func Just(items []*AdsItem) Stream {
// 	source := make(chan *AdsItem, len(items))
// 	for _, item := range items {
// 		source <- item
// 	}
// 	close(source)

// 	return Range(source)
// }

// // Range converts the given channel to a Stream.
// func Range(source <-chan *AdsItem) Stream {
// 	return Stream{
// 		source: source,
// 	}
// }

// // Walk lets the callers handle each item, the caller may write zero, one or more items base on the given item.
// func (s Stream) Walk(fn WalkFunc) Stream {
// 	pipe := make(chan *AdsItem, defaultWorkers)

// 	go func() {
// 		// var wg sync.WaitGroup

// 		for item := range s.source {
// 			// // important, used in another goroutine
// 			val := item
// 			// wg.Add(1)
// 			// // better to safely run caller defined method
// 			// func() {
// 			// 	defer wg.Done()
// 			fn(val, pipe)
// 			// }()
// 		}

// 		// wg.Wait()
// 		close(pipe)
// 	}()

// 	return Range(pipe)
// }

// // Filter filters the items by the given FilterFunc.
// func (s Stream) Filter(fn FilterFunc) Stream {
// 	return s.Walk(func(item *AdsItem, pipe chan<- *AdsItem) {
// 		if fn(item) {
// 			pipe <- item
// 		}
// 	})
// }

// type AdsItem struct {
// 	A int
// 	B int
// }

// func main() {
// 	items := prepareFilterData()
// 	starttime := time.Now().UnixNano()

// 	Just(items).Filter(Filter1Single).Filter(Filter2Single).Filter(Filter3Single).Filter(Filter4Single).Filter(Filter5Single).Filter(Filter6Single).Filter(Filter7Single).Done()

// 	fmt.Println("forceContentFilter time ", time.Now().UnixNano()-starttime)
// 	// for item := range s.source {
// 	// 	if item.A == 1 || item.A == 2 || item.A == 3 || item.A == 4 {
// 	// 		fmt.Println(item.A, item.B)
// 	// 	}
// 	// 	if item.B == 5 || item.B == 6 || item.B == 7 {
// 	// 		fmt.Println(item.A, item.B)
// 	// 	}
// 	// }
// }

// func prepareFilterData() []*AdsItem {
// 	adsitem := make([]*AdsItem, 0)
// 	for i := 0; i < 100000000; i++ {
// 		temp := &AdsItem{
// 			A: rand.Intn(10),
// 			B: rand.Intn(10),
// 		}
// 		adsitem = append(adsitem, temp)
// 	}
// 	return adsitem
// }
