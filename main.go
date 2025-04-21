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
	}

	l.Delete(0)
	l.Delete(0)
	l.Delete(0)
	l.Delete(0)
	l.Delete(l.Len() - 1)
	l.Delete(l.Len() - 1)
	l.Delete(l.Len() - 1)
	l.Delete(l.Len() - 1)
	fmt.Println(l.Values())
	fmt.Println(l.Len())

	l.Update(0, 5)
	l.Update(2, 10)
	l.Update(6, 5)
	l.Update(3, -1)
	fmt.Println(l.Values())
	fmt.Println(l.Len())

	fmt.Println(l.Get(0))
	fmt.Println(l.Get(5))
	fmt.Println(l.Get(10))

	//l.Clear()
	//fmt.Println(l.Values())

	listx.Sort(l)
	fmt.Println(l.Values())
	fmt.Println(l.Len())

	a := []int{0, 1, 2, 3, 4, 5}
	ll := listx.NewLinkedListFromSlice(a)
	fmt.Println(ll.Values())
	fmt.Println(ll.Len())
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(5))
	fmt.Println(ll.Get(10))
}
