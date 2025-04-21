package listx

import (
	"github.com/yzletter/go-kit/gokit/errs"
)

type list[T any] interface {
	Len() int
	Values() []T
	InsertToHead(val T)
	InsertToTail(val T)
	Insert(idx int, val T) error
	Delete(idx int) error
	Update(idx int, val T) error
	Get(idx int) (T, error)
	Clear()
}

// LinkedList 带头结点的双向循环链表
type LinkedList[T any] struct {
	head   *node[T] // 头结点
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	// 初始头结点
	head := &node[T]{}
	head.next, head.prev = head, head

	return &LinkedList[T]{
		head:   head,
		length: 0,
	}
}

// NewLinkedListFromSlice 将切片转为链表
func NewLinkedListFromSlice[T any](target []T) *LinkedList[T] {
	// 初始头结点
	head := &node[T]{}
	head.next, head.prev = head, head

	for i := len(target) - 1; i >= 0; i-- {
		insertBefore(newNode(target[i]), head.next)
	}

	return &LinkedList[T]{
		head:   head,
		length: len(target),
	}
}

// Len 返回当前链表长度
func (l *LinkedList[T]) Len() int {
	return l.length
}

// Values 返回链表中所有的值构成的切片
func (l *LinkedList[T]) Values() (ans []T) {
	for i := l.head.next; i != l.head; i = i.next {
		ans = append(ans, i.val)
	}
	return
}

// InsertToHead 头插法
func (l *LinkedList[T]) InsertToHead(val T) {
	newNode := newNode(val)            // 创建节点
	insertBefore(newNode, l.head.next) // 插入节点
	l.length++                         // 修改长度
}

// InsertToTail 尾插法
func (l *LinkedList[T]) InsertToTail(val T) {
	newNode := newNode(val)       // 创建节点
	insertBefore(newNode, l.head) // 插入节点
	l.length++                    // 修改长度
}

// Insert 插入到指定位置
func (l *LinkedList[T]) Insert(idx int, val T) error {
	// 判断下标
	if !l.checkInsertIndex(idx) {
		return errs.ErrInvalidIndex
	}

	if idx == 0 { // 头插
		l.InsertToHead(val)
		return nil
	} else if idx == l.length { // 尾插
		l.InsertToTail(val)
		return nil
	}

	nowNode := l.findNode(idx)     // 查找节点
	newNode := newNode(val)        // 创建节点
	insertBefore(newNode, nowNode) // 插入节点
	l.length++                     // 修改长度

	return nil
}

// Delete 删除指定下标节点
func (l *LinkedList[T]) Delete(idx int) error {
	if !l.checkDeleteIndex(idx) {
		return errs.ErrInvalidIndex
	}

	nowNode := l.findNode(idx) // 查找节点
	// 删除节点
	deleteNode(nowNode)

	l.length--
	return nil
}

// Update 更新指定下标节点
func (l *LinkedList[T]) Update(idx int, val T) error {
	if !l.checkDeleteIndex(idx) {
		return errs.ErrInvalidIndex
	}
	nowNode := l.findNode(idx)
	nowNode.val = val
	return nil
}

// Get 获取指定下标节点的 val
func (l *LinkedList[T]) Get(idx int) (T, error) {
	if !l.checkDeleteIndex(idx) {
		var t T
		return t, errs.ErrInvalidIndex
	}

	return l.findNode(idx).val, nil
}

// Clear 清空链表
func (l *LinkedList[T]) Clear() {
	l.head.next = l.head
	l.head.prev = l.head
}

// checkInsertIndex 判断插入下标合法性（0 ~ n)
func (l *LinkedList[T]) checkInsertIndex(idx int) bool {
	if idx >= 0 && idx <= l.Len() {
		return true
	}
	return false
}

// checkDeleteIndex 判断删除下标合法性（0 ~ n - 1)
func (l *LinkedList[T]) checkDeleteIndex(idx int) bool {
	if idx >= 0 && idx <= l.Len()-1 {
		return true
	}
	return false
}

// findNode 寻找下标为 idx 的节点
func (l *LinkedList[T]) findNode(idx int) *node[T] {
	nowNode := l.head.next
	for i := 0; i < idx; i++ {
		nowNode = nowNode.next
	}
	return nowNode
}

// lastNode 返回最后一个节点
func (l *LinkedList[T]) lastNode() *node[T] {
	return l.head.prev
}
