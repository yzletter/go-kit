package listx

import (
	"github.com/yzletter/go-kit/slicex"
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort 链表排序
func Sort[T constraints.Ordered](l *LinkedList[T]) {
	// 将节点存入数组
	nodes := nodesSlice(l)
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

// Reverse 返回翻转后的新链表
func Reverse[T any](l *LinkedList[T]) *LinkedList[T] {
	// 将节点值存入数组
	nodesValue := l.Values()
	slicex.ReverseItSelf(nodesValue)

	res := NewLinkedListFromSlice(nodesValue)
	return res
}

// ReverseItSelf 翻转链表自身
func ReverseItSelf[T any](l *LinkedList[T]) {
	// 将节点存入数组
	nodes := nodesSlice(l)
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
