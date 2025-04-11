package main

import (
	"fmt"
	"go-kit/queuex"
)

func main() {
	heap := queuex.NewPriorityQueue[int](func(a, b int) bool {
		return a > b
	})

	a := []int{13, 5, 3, 7, 8, 10}

	for _, v := range a {
		heap.Push(v)
	}

	for heap.Size() > 0 {
		v, _ := heap.Top()
		fmt.Print(v, ", ")
		heap.Pop()
	}

	//fmt.Println(heap.Size())
	//fmt.Println(heap.Size())
	//fmt.Println(heap.Empty())
	//heap.Clear()
	//fmt.Println(heap.Size())

	//v, _ := heap.Top()
	//fmt.Println(v)
	//heap.Pop()
	//v, _ = heap.Top()
	//fmt.Println(v)
}
