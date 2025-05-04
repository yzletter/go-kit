package treex

import (
	"golang.org/x/exp/constraints"
)

type SegmentTree[T constraints.Ordered] struct {
	length int //	树的大小
	tr     []T
	lazy   []T
}

func NewSegmentTree[T constraints.Ordered](length int) *SegmentTree[T] {
	return &SegmentTree[T]{
		length: length,
		tr:     make([]T, length*4),
		lazy:   make([]T, length*4),
	}
}

// BuildTreeFromSlice 从切片中构建树
func (tree *SegmentTree[T]) BuildTreeFromSlice(src []T) {
	tree.buildTreeFromSlice(src, 1, tree.length, 1)
}

// BuildNilTree 构建空树
func (tree *SegmentTree[T]) BuildNilTree() {
	tree.buildNilTree(1, tree.length, 1)
}

func (tree *SegmentTree[T]) buildTreeFromSlice(src []T, st, ed, u int) {
	// 递归出口
	if st == ed {
		tree.tr[u] = src[st]
		return
	}

	mid := (st + ed) >> 1
	tree.buildTreeFromSlice(src, st, mid, u<<1)
	tree.buildTreeFromSlice(src, mid+1, ed, u<<1|1)
	tree.pushUp(u)
}

func (tree *SegmentTree[T]) buildNilTree(st, ed, u int) {
	// 递归出口
	if st == ed {
		var t T
		tree.tr[u] = t
		return
	}

	mid := (st + ed) >> 1
	tree.buildNilTree(st, mid, u<<1)
	tree.buildNilTree(mid+1, ed, u<<1|1)
	tree.pushUp(u)
}

func (tree *SegmentTree[T]) update(st, ed, u int, x T) {
	tree.lazy[u] = x
}

func (tree *SegmentTree[T]) pushUp(u int) {
	tree.tr[u] = max(tree.tr[u<<1], tree.tr[u<<1|1])
}

func (tree *SegmentTree[T]) pushDown(st, ed, u int) {

	mid := (st + ed) >> 1
	tree.update(st, mid, u<<1, tree.lazy[u])
	tree.update(mid+1, ed, u<<1|1, tree.lazy[u])
}

func (tree *SegmentTree[T]) Modify() {

}

func (tree *SegmentTree[T]) Query() {

}
