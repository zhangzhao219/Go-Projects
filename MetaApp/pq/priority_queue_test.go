package main

import (
	"fmt"
	"testing"
)

type testS struct {
	value float64
}

func (s *testS) Priority() float64 {
	return s.value
}

func (s *testS) Compare(a, b PriorityValue) bool {
	return s.Priority() > b.Priority()
}

func main() {
	TestPriorityQueue_Push(&testing.T{})
}

var fs = []float64{0.6046602879796196, 0.9405090880450124, 0.6645600532184904, 0.4377141871869802, 0.4246374970712657, 0.6868230728671094, 0.06563701921747622, 0.15651925473279124, 0.09696951891448456, 0.30091186058528707}

func TestPriorityQueue_Init(t *testing.T) {
	l := make([]*testS, 0, 10)
	for i := 0; i < 10; i++ {
		l = append(l, &testS{
			fs[i],
		})
	}
	for _, elem := range l {
		fmt.Print(elem.value, " ")
	}
	fmt.Println()
	h := Init[*testS](l)
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println()
}

func TestPriorityQueue_Push(t *testing.T) {
	l := make([]*testS, 0, 10)
	for i := 0; i < 10; i++ {
		l = append(l, &testS{
			fs[i],
		})
	}
	for _, elem := range l {
		fmt.Print(elem.value, " ")
	}
	fmt.Println()
	h := Init[*testS](l)
	h.Push(&testS{2})
	h.Push(&testS{0.03})
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println()
}

func TestPriorityQueue_Pop(t *testing.T) {
	l := make([]*testS, 0, 10)
	for i := 0; i < 10; i++ {
		l = append(l, &testS{
			fs[i],
		})
	}
	for _, elem := range l {
		fmt.Print(elem.value, " ")
	}
	fmt.Println()
	h := Init[*testS](l)
	fmt.Println(h.Pop().value)
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println(h.Pop().value)
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println(h.Pop().value)
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println(h.Pop().value)
	for h.Len() != 0 {
		t.Log(h.Pop().value)
	}
	fmt.Println()
}
