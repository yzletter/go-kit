package stackx

import "errors"

type stack[T comparable] interface {
	Push(val T)
	Top() (T, error)
	Pop()
	Size() int
	Empty() bool
	Clear()
}

type Stack[T comparable] struct {
	data []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push 新元素入栈
func (stk *Stack[T]) Push(val T) {
	stk.data = append(stk.data, val)
}

// Top 取栈顶
func (stk *Stack[T]) Top() (T, error) {
	if !stk.Empty() {
		return stk.data[stk.Size()-1], nil
	}
	var t T
	return t, errors.New("栈为空")
}

// Pop 栈顶弹出
func (stk *Stack[T]) Pop() {
	if !stk.Empty() {
		stk.data = stk.data[:stk.Size()-1]
	}
}

// Size 栈内元素个数
func (stk *Stack[T]) Size() int {
	return len(stk.data)
}

// Empty 判断栈是否为空
func (stk *Stack[T]) Empty() bool {
	return len(stk.data) == 0
}

// Clear 清空栈
func (stk *Stack[T]) Clear() {
	stk.data = make([]T, 0)
}
