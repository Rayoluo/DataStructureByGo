package main

import (
	"fmt"
	"01BasicLinkedList/linkedList"
)

func main() {
	list := linkedList.NewLinkedList()
	for i := 0; i < 5; i++ {
		list.AddFirst(i)
		fmt.Println(list)
	}
	list.Add(2, 666)
	fmt.Println(list)
	list.Remove(2)
	fmt.Println(list)
	list.RemoveFirst()
	fmt.Println(list)
	list.RemoveLast()
	fmt.Println(list)
}