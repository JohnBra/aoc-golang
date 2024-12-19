package datastructures

/*
	Heap with struct implementation
*/

type HeapItem[T any] struct {
	sort  int
	value T
}

// Min heap; 'sort' (int) field in HeapItem is used for sorting
type Heap[T any] []HeapItem[T]

func (h Heap[T]) Len() int {
	return len(h)
}

func (h Heap[T]) Less(i, j int) bool {
	return h[i].sort < h[j].sort
}

func (h Heap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(HeapItem[T]))
}

func (h *Heap[T]) Pop() any {
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
