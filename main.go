package main

import (
	"fmt"
	"go-kit/queuex"
)

func main() {

	dq := queuex.NewDeque[int]()
	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushFront(2)
	dq.PushFront(3)
	//fmt.Println(dq.Front())
	//fmt.Println(dq.Back())

	for dq.Size() > 0 {
		fmt.Println(dq.Front())
		dq.PopFront()
	}

	dq.Clear()
	fmt.Println(dq.Front())

	//a := []int{13, 5, 3, 7, 8, 10}
	//a, _ = slicex.Insert(a, 1, 0)
	//a, _ = slicex.Insert(a, 1, 5)
	//a, _ = slicex.Insert(a, 1, len(a))
	//fmt.Println(a)
	//
	//b := []int{1, 2, 4, 5, 16, 18}
	//idx1 := slicex.UpperBound(b, 4)
	//idx2 := slicex.LowerBound(b, 4)
	//fmt.Println(idx1)
	//fmt.Println(idx2)
	//
	//c := []int{1, 4, 2, 2, 4, 5, 16, 2}
	//c = slicex.Unique(c)
	//fmt.Println(c)
	//
	//d := slicex.Reverse(c)
	//fmt.Println(d)
	//
	//slicex.ReverseItSelf(d)
	//fmt.Println(d)
	//
	//a = []int{13, 5, 3, 7, 8, 10}
	//idx3, _ := slicex.Find(a, 3)
	//idx4, _ := slicex.Find(a, 17)
	//fmt.Println(idx3)
	//fmt.Println(idx4)
}
