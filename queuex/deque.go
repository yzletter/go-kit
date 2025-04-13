package queuex

import (
	"errors"
	"go-kit/slicex"
)

type deque[T any] struct {
	data []T
}

func NewDeque[T any]() *deque[T] {
	return &deque[T]{
		data: make([]T, 0),
	}
}

func (dq *deque[T]) PopFront() {
	dq.data = dq.data[1:]
}
func (dq *deque[T]) PopBack() {
	dq.data = dq.data[:dq.Size()-1]
}

func (dq *deque[T]) PushFront(val T) error {
	var err error
	dq.data, err = slicex.Insert[T](dq.data, val, 0)
	return err
}

func (dq *deque[T]) PushBack(val T) error {
	var err error
	dq.data, err = slicex.Insert[T](dq.data, val, dq.Size())
	return err
}

func (dq *deque[T]) Front() (T, error) {
	if dq.Empty() {
		var t T
		return t, errors.New("队列为空")
	}
	return dq.data[0], nil
}

func (dq *deque[T]) Back() (T, error) {
	if dq.Empty() {
		var t T
		return t, errors.New("队列为空")
	}
	return dq.data[dq.Size()-1], nil
}

func (dq *deque[T]) Empty() bool {
	return dq.Size() == 0
}

func (dq *deque[T]) Size() int {
	return len(dq.data)
}

func (dq *deque[T]) Clear() {
	dq.data = make([]T, 0)
}
