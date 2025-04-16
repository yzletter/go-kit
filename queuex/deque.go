package queuex

import (
	"errors"
	"github.com/yzletter/go-kit/slicex"
)

type deque[T any] struct {
	data []T
}

func NewDeque[T any]() *deque[T] {
	return &deque[T]{
		data: make([]T, 0),
	}
}

// PopFront 弹出队头
func (dq *deque[T]) PopFront() {
	dq.data = dq.data[1:]
	dq.data = slicex.Shrink(dq.data)
}

// PopBack 弹出队尾
func (dq *deque[T]) PopBack() {
	dq.data = dq.data[:dq.Size()-1]
	dq.data = slicex.Shrink(dq.data)
}

// PushFront 插入队头
func (dq *deque[T]) PushFront(val T) error {
	var err error
	dq.data, err = slicex.Insert[T](dq.data, val, 0)
	return err
}

// PushBack 插入队尾
func (dq *deque[T]) PushBack(val T) error {
	var err error
	dq.data, err = slicex.Insert[T](dq.data, val, dq.Size())
	return err
}

// Front 取队头
func (dq *deque[T]) Front() (T, error) {
	if dq.Empty() {
		var t T
		return t, errors.New("队列为空")
	}
	return dq.data[0], nil
}

// Back 取队尾
func (dq *deque[T]) Back() (T, error) {
	if dq.Empty() {
		var t T
		return t, errors.New("队列为空")
	}
	return dq.data[dq.Size()-1], nil
}

// Empty 判断队空
func (dq *deque[T]) Empty() bool {
	return dq.Size() == 0
}

// Size 返回队列中元素个数
func (dq *deque[T]) Size() int {
	return len(dq.data)
}

// Clear 清空队列
func (dq *deque[T]) Clear() {
	dq.data = make([]T, 0)
}
