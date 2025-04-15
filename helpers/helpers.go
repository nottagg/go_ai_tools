package helpers

import (
	"container/heap"
)

// Contains helper functions used across the toolset

type Coordinate struct {
	X int
	Y int
}

// MapKeysToSlice takes a generic map and returns the keys as a slice
func MapKeysToSlice[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValuesToSlice takes a generic map and returns the values as a slice
func MapValuesToSlice[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetGridDirections(allowDiagonal bool) [][2]int {
	if allowDiagonal {
		return [][2]int{
			{0, -1},  // Up
			{1, 0},   // Right
			{0, 1},   // Down
			{-1, 0},  // Left
			{-1, -1}, // Up-Left
			{1, -1},  // Up-Right
			{1, 1},   // Down-Right
			{-1, 1},  // Down-Left
		}
	}
	return [][2]int{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
}

type Item[T any] struct {
	Value    T
	Priority float64
	Index    int
}

type PriorityQueue[T any] []*Item[T]

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq

}

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.Index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Enqueue(value T, priority float64) {
	item := &Item[T]{
		Value:    value,
		Priority: priority,
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) Dequeue() T {
	if pq.Len() == 0 {
		var zeroValue T
		return zeroValue
	}
	item := heap.Pop(pq).(*Item[T])
	return item.Value
}
