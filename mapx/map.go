package mapx

import (
	"errors"
)

// Keys 返回 map 中所有的键
func Keys[K comparable, V any](target map[K]V) []K {
	res := make([]K, 0, len(target))

	for k := range target {
		res = append(res, k)
	}

	return res
}

// Values 返回 map 中所有的值
func Values[K comparable, V any](target map[K]V) []V {
	res := make([]V, 0, len(target))

	for _, v := range target {
		res = append(res, v)
	}

	return res
}

// KeysAndValues 返回 map 中所有的键和值
func KeysAndValues[K comparable, V any](target map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(target))
	values := make([]V, 0, len(target))

	for k := range target {
		keys = append(keys, k)
		values = append(values, target[k])
	}

	return keys, values
}

// ToMap 将 K V 切片转 map
func ToMap[K comparable, V any](keys []K, values []V) (map[K]V, error) {

	if keys == nil || values == nil {
		return nil, errors.New("keys 与 values 均不可为nil")
	}

	if len(keys) != len(values) {
		return nil, errors.New("传入参数不符合要求")
	}

	mp := map[K]V{}
	for i := 0; i < len(keys); i++ {
		mp[keys[i]] = values[i]
	}

	return mp, nil
}
