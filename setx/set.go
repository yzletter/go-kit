package setx

type set[T comparable] interface {
	Insert(key T)
	Size() int
	Empty() bool
	Clear()
	Erase(key T)
	Exist(key T) bool
	Keys() []T
}

type Set[T comparable] struct {
	mp map[T]struct{}
}

// NewSet 构造函数
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		mp: make(map[T]struct{}),
	}
}

// Insert 插入元素
func (s *Set[T]) Insert(key T) {
	s.mp[key] = struct{}{}
}

// Size 返回元素个数
func (s *Set[T]) Size() int {
	return len(s.mp)
}

// Empty 判断集合是否为空集
func (s *Set[T]) Empty() bool {
	return len(s.mp) == 0
}

// Clear 清空集合
func (s *Set[T]) Clear() {
	s.mp = make(map[T]struct{})
}

// Erase 删除元素
func (s *Set[T]) Erase(key T) {
	delete(s.mp, key)
}

// Exist 查询元素
func (s *Set[T]) Exist(key T) bool {
	_, ok := s.mp[key]
	return ok
}

// Keys 返回所有元素（无序）
func (s *Set[T]) Keys() []T {
	result := make([]T, 0, len(s.mp))
	for k := range s.mp {
		result = append(result, k)
	}
	return result
}
