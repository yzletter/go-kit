package treex

import "github.com/yzletter/go-kit/gokit"

type SegmentTree[T gokit.Number] struct {
	length int //	树的大小
	tr     []T
	lazy   []T
}

func NewSegmentTree[T gokit.Number](length int) *SegmentTree[T] {
	return &SegmentTree[T]{
		length: length,
		tr:     make([]T, length*4),
		lazy:   make([]T, length*4),
	}
}

func (tree *SegmentTree[T]) BuildTree(src []T, st, ed, u int) {
	// 递归出口
	if st == ed {
		tree.tr[u] = src[st]
		return
	}

	mid := (st + ed) >> 1
	tree.BuildTree(src, st, mid, u<<1)
	tree.BuildTree(src, mid+1, ed, u<<1|1)
	tree.pushUp(u)
}

func (tree *SegmentTree[T]) update() {

}

func (tree *SegmentTree[T]) pushUp(u int) {
	tree.tr[u] = tree.tr[u<<1] + tree.tr[u<<1|1]
}

func (tree *SegmentTree[T]) pushDown() {

}

func (tree *SegmentTree[T]) Modify() {

}

func (tree *SegmentTree[T]) Query() {

}
