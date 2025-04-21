package listx

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort 链表排序
func Sort[T constraints.Ordered](l *LinkedList[T]) {

	// 将节点存入数组
	nodes := make([]*node[T], 0)
	for i := l.head.next; i != l.head; i = i.next {
		nodes = append(nodes, i)
	}

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
