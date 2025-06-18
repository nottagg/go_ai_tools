package helpers

import (
	"container/heap"
	"math"
)

// Contains helper functions used across the toolset

type Coordinate struct {
	X float64
	Y float64
}

func IntegerAbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func EuclideanDistance(x1, x2, y1, y2 int) float64 {
	dx := IntegerAbsoluteValue(x1 - x2)
	dy := IntegerAbsoluteValue(y1 - y2)
	return math.Sqrt(float64(dx*dx + dy*dy))
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

// Straight ripped from GO Heap documentation
// https://pkg.go.dev/container/heap
type Item[T any] struct {
	Value    T       // The value of the item; arbitrary.
	Priority float64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority float64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
