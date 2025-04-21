package errs

import "errors"

var (
	ErrInvalidIndex     = errors.New("go-kit: 传入下标错误")
	ErrEmpty            = errors.New("go-kit: 目标数据结构内没有元素")
	ErrNotExist         = errors.New("go-kit: 元素不存在")
	ErrInvalidParameter = errors.New("go-kit: keys 与 values 均不可为 nil")
)
