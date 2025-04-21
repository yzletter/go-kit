package listx

type node[T any] struct {
	next *node[T]
	prev *node[T]
	val  T
}

// insertBefore 前插
func insertBefore[T any](a, b *node[T]) {
	a.next = b
	a.prev = b.prev
	b.prev.next = a
	b.prev = a
}

// nodes 将链表节点存入切片
func nodesSlice[T any](l *LinkedList[T]) []*node[T] {
	nodes := make([]*node[T], 0)
	for i := l.head.next; i != l.head; i = i.next {
		nodes = append(nodes, i)
	}
	return nodes
}

// deleteNode 删除指定节点
func deleteNode[T any](node *node[T]) {
	// 删除节点
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
}

// newNode 创建一个新节点
func newNode[T any](val T) *node[T] {
	return &node[T]{val: val}
}
