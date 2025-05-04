package gokit

type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type SegmentTreeNumber interface {
	~int | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
