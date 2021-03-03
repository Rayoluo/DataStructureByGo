package linkedListStack

import (
	list "01BasicLinkedList/linkedList"
	"bytes"
)

type LinkedListStack struct {
	*list.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		list.NewLinkedList(),
	}
}

func (s *LinkedListStack) GetStackSize() int {
	return s.GetSize()
}

func (s *LinkedListStack) IsStackEmpty() bool {
	return s.IsEmpty()
}

func (s *LinkedListStack) Push(elem interface{}) {
	s.AddFirst(elem)
}

func (s *LinkedListStack) Peek() interface{} {
	return s.GetFirst()
}

func (s *LinkedListStack) Pop() interface{} {
	return s.RemoveFirst()
}

func (s *LinkedListStack) LinkedListStackToString() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Stack: top ")
	buffer.WriteString(s.String())
	return buffer.String()
}