package datastructures

// Implementation with slices. Would be more efficient with ring buffer
// -- front ------- back --
type Deque[T comparable] struct {
	items []T
}

func NewDeque[T comparable](items []T) *Deque[T] {
	deque := &Deque[T]{
		items: append([]T{}, items...),
	}

	return deque
}

func (d *Deque[T]) Len() int {
	return len(d.items)
}

func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) PopBack() (T, bool) {
	var result T
	if len(d.items) == 0 {
		return result, false
	}
	result = d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return result, true
}

func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) PopFront() (T, bool) {
	var result T
	if len(d.items) == 0 {
		return result, false
	}
	result = d.items[0]
	d.items = d.items[1:]
	return result, true
}
