package main

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Length() int {
	return len(q.items)
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
