package queuex

import "errors"

type queue[T comparable] interface {
	Pop()
	Front() (T, error)
	Push(val T)
	Size() int
	Back() (T, error)
	Empty() bool
	Clear()
}

type Queue[T comparable] struct {
	data []T
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// Pop 队头出队
func (q *Queue[T]) Pop() {
	if !q.Empty() {
		q.data = q.data[1:]
	}
}

// Front 取队头
func (q *Queue[T]) Front() (T, error) {
	if !q.Empty() {
		return q.data[0], nil
	}
	var t T
	return t, errors.New("go-kit: 队列为空")
}

// Push 新元素入队
func (q *Queue[T]) Push(val T) {
	q.data = append(q.data, val)
}

// Size 队列中元素个数
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// Back 取队尾
func (q *Queue[T]) Back() (T, error) {
	if !q.Empty() {
		return q.data[q.Size()-1], nil
	}
	var t T
	return t, errors.New("go-kit: 队列为空")
}

// Empty 判断队空
func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

// Clear 清空队列
func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}
