package main

type PriorityValue interface {
	Priority() float64
	Compare(a, b PriorityValue) bool
	//Index() int
}

const (
	QUEUE_EMPTY = "queue length:0"
)

const (
	NULL_VALUE = -1
)

type PriorityQueue[T PriorityValue] struct {
	queue []T
}

func Init[T PriorityValue](queue []T) *PriorityQueue[T] {
	list := make([]T, len(queue))
	copy(list, queue)
	//todo
	pq := &PriorityQueue[T]{
		list,
	}
	for i := len(list)/2 - 1; i >= 0; i-- {
		pq.down(i, len(list))
	}
	return pq
}

func (pq *PriorityQueue[T]) Len() int { return len(pq.queue) }

func (pq *PriorityQueue[T]) less(i, j int) bool {
	return pq.queue[i].Compare(pq.queue[i], pq.queue[j])
	//return pq.queue[i].Priority() > pq.queue[j].Priority()
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQueue[T]) FirstValue() *T {
	if pq.Len() == 0 {
		return nil
	}
	return &pq.queue[pq.Len()-1]
}

func (pq *PriorityQueue[T]) Pop() T {
	if pq.Len() == 0 {
		panic(QUEUE_EMPTY)
	}
	n := pq.Len() - 1
	pq.swap(0, n)
	pq.down(0, n)

	item := pq.queue[pq.Len()-1]
	pq.queue = pq.queue[0 : pq.Len()-1]
	return item
}

func (pq *PriorityQueue[T]) Push(v T) {
	if pq.Len() == cap(pq.queue) {
		pq.Pop()
	}
	pq.queue = append(pq.queue, v)
	pq.up(pq.Len() - 1)
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		i = j
	}
	return i > i0
}
