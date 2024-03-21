package main

import "fmt"

type UFS struct {
	father []int
}

func NewUFS(size int) *UFS {
	return &UFS{
		father: make([]int, size),
	}
}

func (u *UFS) Init() {
	for i := 0; i < len(u.father); i++ {
		u.father[i] = i
	}
}

func (u *UFS) FindFather(x int) int {
	a := x
	for x != u.father[x] {
		x = u.father[x]
	}
	for a != u.father[a] {
		z := a
		a = u.father[a]
		u.father[z] = x
	}
	return x
}

func (u *UFS) Union(a, b int) {
	faA := u.FindFather(a)
	faB := u.FindFather(b)
	if faA != faB {
		u.father[faA] = faB
	}
}

func (u *UFS) Print() {
	fathermap := make(map[int][]int, 0)
	for i := 0; i < len(u.father); i++ {
		father := u.FindFather(i)
		_, ok := fathermap[father]
		if ok {
			fathermap[father] = append(fathermap[father], i)
		} else {
			fathermap[father] = make([]int, 1)
			fathermap[father][0] = i
		}
	}
	for k, v := range fathermap {
		fmt.Println(k, v)
	}
}

func main() {
	u := NewUFS(10)
	u.Init()
	u.Union(0, 1)
	u.Union(2, 3)
	u.Union(3, 5)
	u.Union(7, 5)
	u.Union(4, 6)
	u.Print()
}
