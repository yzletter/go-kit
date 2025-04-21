package main

import (
	"fmt"
	"github.com/yzletter/go-kit/listx"
)

func main() {

	l := listx.NewLinkedList[int]()

	{
		l.InsertToHead(6)
		l.InsertToHead(5)
		l.InsertToHead(4)
		l.InsertToTail(7)
		l.InsertToTail(8)
		l.InsertToTail(5)
		l.Insert(0, 1)
		l.Insert(0, 1)
		l.Insert(0, 1)
		l.Insert(l.Len(), 2)
		l.Insert(l.Len(), 2)
		l.Insert(l.Len(), 2)
		l.Insert(2, 6)
		l.Insert(2, 6)
		l.Insert(2, 6)
		l.Delete(0)
		l.Delete(0)
		l.Delete(0)
		l.Delete(0)
		l.Delete(l.Len() - 1)
		l.Delete(l.Len() - 1)
		l.Delete(l.Len() - 1)
		l.Delete(l.Len() - 1)
		l.Update(0, 5)
		l.Update(2, 10)
		l.Update(6, 5)
		l.Update(3, -1)
	}
	listx.Sort(l)

	b := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 6, 6, 6, 10, 10, 10}
	t := listx.NewLinkedListFromSlice(b)
	listx.UniqueByOrder(t)
	fmt.Println(t.Values())

	d := []int{0, 1, 0, 1, 2, 1, 2, 4, 2, 6, 3, 6, 5, 10, 10}
	tt := listx.NewLinkedListFromSlice(d)
	newTT := listx.Unique(tt)
	fmt.Println(newTT.Values())
}
