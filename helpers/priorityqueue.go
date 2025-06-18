package helpers

import "container/heap"

type Item[T any] struct {
	Value    T
	Priority float64
	index    int
}

// PriorityQueue is a min-heap priority queue for items of type T.
type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// New creates a new priority queue.
func New[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}

// PushItem adds a new item to the priority queue.
func (pq *PriorityQueue[T]) PushItem(value T, priority float64) {
	heap.Push(pq, &Item[T]{Value: value, Priority: priority})
}

// PopItem removes and returns the item with the highest priority.
func (pq *PriorityQueue[T]) PopItem() T {
	item := heap.Pop(pq).(*Item[T])
	return item.Value
}

// IsEmpty returns true if the queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}
