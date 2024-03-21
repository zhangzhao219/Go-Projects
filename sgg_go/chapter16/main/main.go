package main

import (
	"fmt"
	"sgg_go/chapter16/store"
)

func main() {
	monster1 := store.Newmonster("a", 1, "abc")
	a := monster1.Store("test1.txt")
	if a {
		fmt.Println("true")
	}
	b := monster1.ReStore("test1.txt")
	fmt.Println(b)
}
