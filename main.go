package main

import (
	"fmt"
	"go-kit/stackx"
)

func main() {
	stk := stackx.NewStack[int]()

	for i := 0; i < 10; i++ {
		stk.Push(i)
	}

	fmt.Println(stk.Top())
	fmt.Println(stk.Size())
	stk.Pop()
	fmt.Println(stk.Size())
	stk.Clear()
	fmt.Println(stk.Size())

}
