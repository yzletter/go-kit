package listx

import (
	"github.com/yzletter/go-kit/slicex"
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort 链表排序
func Sort[T constraints.Ordered](l *LinkedList[T]) {
	// 将节点存入数组
	nodes := nodes(l)
	// 排序
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val
	})
	// 更新指向
	for i := 0; i < len(nodes); i++ {
		if i != len(nodes)-1 {
			nodes[i].next = nodes[i+1]
		}
		if i != 0 {
			nodes[i].prev = nodes[i-1]
		}
	}
	// 构建循环链表
	nodes[0].prev = l.head
	l.head.next = nodes[0]
	nodes[len(nodes)-1].next = l.head
	l.head.prev = nodes[len(nodes)-1]
}

// ReverseItSelf 翻转链表自身
func ReverseItSelf[T any](l *LinkedList[T]) {
	// 将节点存入数组
	nodes := nodes(l)
	slicex.ReverseItSelf(nodes)
	// 更新指向
	for i := 0; i < len(nodes); i++ {
		if i != len(nodes)-1 {
			nodes[i].next = nodes[i+1]
		}
		if i != 0 {
			nodes[i].prev = nodes[i-1]
		}
	}
	// 构建循环链表
	nodes[0].prev = l.head
	l.head.next = nodes[0]
	nodes[len(nodes)-1].next = l.head
	l.head.prev = nodes[len(nodes)-1]
}

// Reverse 返回翻转后的新链表
func Reverse[T any](l *LinkedList[T]) *LinkedList[T] {
	// 将节点值存入数组
	nodesValue := l.Values()
	slicex.ReverseItSelf(nodesValue)

	res := NewLinkedListFromSlice(nodesValue)
	return res
}

// nodes 将链表节点存入切片
func nodes[T any](l *LinkedList[T]) []*node[T] {
	nodes := make([]*node[T], 0)
	for i := l.head.next; i != l.head; i = i.next {
		nodes = append(nodes, i)
	}
	return nodes
}

// UniqueByOrder 对排序后的链表进行去重
func UniqueByOrder[T comparable](l *LinkedList[T]) {
	nowNode := l.head

	// current.Next 当前未处理的第一个节点 current.Next.Next 当前未处理的第一个节点的下一个节点
	for nowNode.next != l.head && nowNode.next.next != l.head {

		nowVal := nowNode.next.val

		if nowNode.next.next.val == nowVal {
			// 删除重复节点
			for nowNode.next.next != l.head && nowNode.next.next.val == nowVal {
				deleteNode(nowNode.next.next)
			}
		} else {
			nowNode = nowNode.next
		}
	}
}

// Unique 对链表去重，返回顺序不一定
func Unique[T comparable](l *LinkedList[T]) *LinkedList[T] {
	nodes := l.Values()
	newNodes := slicex.Unique(nodes)

	res := NewLinkedListFromSlice(newNodes)
	return res
}

// insertBefore 前插
func insertBefore[T any](a, b *node[T]) {
	a.next = b
	a.prev = b.prev
	b.prev.next = a
	b.prev = a
}

// deleteNode 删除指定节点
func deleteNode[T any](node *node[T]) {
	// 删除节点
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
}
