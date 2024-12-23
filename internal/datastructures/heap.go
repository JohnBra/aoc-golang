package datastructures

import "container/heap"

/*
	Heap with struct implementation
*/

type HeapItem interface {
	Priority() int
}

// Min heap; Arbitrary heap item must implement priority() int
type Heap []HeapItem

func NewHeap(vals ...HeapItem) *Heap {
	h := &Heap{}
	*h = append(*h, vals...)
	heap.Init(h)
	return h
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Priority() < h[j].Priority()
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(HeapItem))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
	Int Heap implementation
*/

// Min heap; First item in each int slice is used for sorting
type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
