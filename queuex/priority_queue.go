package queuex

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	data        []T
	compareFunc func(a, b T) bool
}

func NewPriorityQueue[T constraints.Ordered](compareFunc func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data:        make([]T, 0),
		compareFunc: compareFunc,
	}
}

// Pop 弹出堆顶
func (heap *PriorityQueue[T]) Pop() {
	if heap.Empty() {
		return
	}
	heap.data[0] = heap.data[heap.Size()-1] // 将最后一个数放到堆顶
	heap.data = heap.data[:heap.Size()-1]   // 删除最后一个数
	// 更新堆
	heap.update(0)
}

// Top 取堆顶元素
func (heap *PriorityQueue[T]) Top() (T, error) {
	if !heap.Empty() {
		return heap.data[0], nil
	}

	var t T
	return t, errors.New("go-kit: 队列为空")
}

// Push 插入新元素
func (heap *PriorityQueue[T]) Push(val T) {
	// 将节点插到堆最后一个元素
	heap.data = append(heap.data, val)

	children := heap.Size() - 1
	parent := (children - 1) >> 1

	// 沿着当前子节点向上，若小于父节点，则交换位置
	for parent >= 0 && heap.compareFunc(heap.data[children], heap.data[parent]) {
		//for parent >= 0 && heap.data[children] < heap.data[parent] {
		heap.data[parent], heap.data[children] = heap.data[children], heap.data[parent]
		children = parent
		parent = (children - 1) >> 1
	}
}

// Size 返回堆中元素个数
func (heap *PriorityQueue[T]) Size() int {
	return len(heap.data)
}

// Empty 判断堆空
func (heap *PriorityQueue[T]) Empty() bool {
	return len(heap.data) == 0
}

// Clear 清空堆
func (heap *PriorityQueue[T]) Clear() {
	heap.data = make([]T, 0)
}

// update 将 idx 号结点沉底
func (heap *PriorityQueue[T]) update(idx int) {
	n := heap.Size()

	target := idx // 最小的节点下标
	for {
		// 左孩子
		if leftchild := idx*2 + 1; leftchild <= n-1 && heap.compareFunc(heap.data[leftchild], heap.data[target]) {
			target = leftchild
		}
		// 右孩子
		if rightchild := idx*2 + 2; rightchild <= n-1 && heap.compareFunc(heap.data[rightchild], heap.data[target]) {
			target = rightchild
		}

		if target == idx {
			break
		}

		// 交换
		heap.data[idx], heap.data[target] = heap.data[target], heap.data[idx]
		idx = target
	}
}
