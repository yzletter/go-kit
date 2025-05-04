package kitx

func ToPtr[T any](t T) *T {
	return &t
}

// Pair 类型
type Pair[K any, V any] struct {
	First  K
	Second V
}
