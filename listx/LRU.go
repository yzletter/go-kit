package listx

import (
	"github.com/yzletter/go-kit/gokit/errs"
	"github.com/yzletter/go-kit/kitx"
)

// LRUCache 满足 LRU (最近最少使用) 缓存约束的数据结构
type LRUCache[K comparable, V any] struct {
	capacity int
	l        *LinkedList[kitx.Pair[K, V]]
	mp       map[K]*node[kitx.Pair[K, V]]
}

// NewLRUCache 构造函数 每一个 Node 包含一个 K-V 键值对
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {

	return &LRUCache[K, V]{
		capacity: capacity,
		l:        NewLinkedList[kitx.Pair[K, V]](),
		mp:       make(map[K]*node[kitx.Pair[K, V]]),
	}
}

// Get 根据 key 返回缓存中对应的值，若不在缓存中，则返回错误
func (cache *LRUCache[K, V]) Get(key K) (kitx.Pair[K, V], error) {
	/// 查询是否在缓存中
	node, ok := cache.mp[key]
	if !ok {
		var t kitx.Pair[K, V]
		return t, errs.ErrNotExist
	}
	// 将 node 位置更新
	cache.moveToTop(node)
	// 返回结果
	return node.val, nil
}

// Put 修改 key 的值为 val, 若不存在，则进行插入
func (cache *LRUCache[K, V]) Put(key K, val V) {
	// 查询是否在缓存中
	node, ok := cache.mp[key]

	// 不在缓存中
	if !ok {
		newNode := newNode[kitx.Pair[K, V]](kitx.Pair[K, V]{First: key, Second: val})
		cache.mp[key] = newNode
		insertBefore(newNode, cache.l.head.next)

		// 容量超限了，删除最后一个节点
		if len(cache.mp) > cache.capacity {
			lastNode := cache.l.lastNode()
			delete(cache.mp, lastNode.val.First) // 先从 map 删除末尾节点的 key
			deleteNode(lastNode)                 // 再将 LRU 中的末尾节点删除
		}

		return
	}

	// 在缓存中，修改值，更新位置
	node.val = kitx.Pair[K, V]{First: key, Second: val}
	cache.moveToTop(node)
}

// moveToTop 将当前 node 放到最前
func (cache *LRUCache[K, V]) moveToTop(node *node[kitx.Pair[K, V]]) {
	deleteNode(node)
	insertBefore(node, cache.l.head.next)
}
