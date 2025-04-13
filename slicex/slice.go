package slicex

import (
	"errors"
	"go-kit/setx"
	"golang.org/x/exp/constraints"
)

// Insert 插入新元素
func Insert[T any](target []T, val T, idx int) ([]T, error) {

	if idx < 0 || idx > len(target) {
		return nil, errors.New("插入非法位置")
	}

	var t T
	target = append(target, t)
	// 将 idx + 1 开始的所有元素后移一位
	for i := len(target) - 1; i > idx; i-- {
		target[i] = target[i-1]
	}
	// 将 val 插入到 idx
	target[idx] = val
	return target, nil
}

// UpperBound 返回 >= x 的最小值的下标
func UpperBound[T constraints.Ordered](target []T, val T) int {
	l, r := 0, len(target)-1
	for l < r {
		mid := (l + r) >> 1
		if target[mid] < val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if target[l] == val {
		return l
	} else {
		return -1
	}
}

// LowerBound 返回 > x 的最小值的下标
func LowerBound[T constraints.Ordered](target []T, val T) int {
	l, r := 0, len(target)-1
	for l < r {
		mid := (l + r) >> 1
		if target[mid] <= val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if target[l] > val {
		return l
	} else {
		return -1
	}
}

// Unique 去重
func Unique[T comparable](target []T) []T {
	hash := setx.NewSet[T]()
	ans := make([]T, 0)
	for _, v := range target {
		hash.Insert(v)
	}
	ans = append(ans, hash.Keys()...)
	return ans
}

// Reverse 返回翻转后的 slice
func Reverse[T any](target []T) []T {
	ans := make([]T, 0)
	for i := len(target) - 1; i >= 0; i-- {
		ans = append(ans, target[i])
	}
	return ans
}

// ReverseItSelf 返回翻转后的自身
func ReverseItSelf[T any](target []T) {
	for i, j := 0, len(target)-1; i < j; i, j = i+1, j-1 {
		target[i], target[j] = target[j], target[i]
	}
}

// Find 返回要查找元素的第一个下标位置
func Find[T comparable](target []T, val T) (int, error) {
	for i := 0; i < len(target)-1; i++ {
		if target[i] == val {
			return i, nil
		}
	}
	return -1, errors.New("元素不存在")
}
