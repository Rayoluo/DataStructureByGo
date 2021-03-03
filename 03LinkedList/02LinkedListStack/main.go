package main

import (
	"02LinkedListStack/linkedListStack"
	"02LinkedListStack/stack"
	"fmt"
)

func main() {
	var s stack.Stack
	s = linkedListStack.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		s.Push(i)
		fmt.Println(s.LinkedListStackToString())
	}
	fmt.Println(s.Pop())
	fmt.Println(s.LinkedListStackToString())
}