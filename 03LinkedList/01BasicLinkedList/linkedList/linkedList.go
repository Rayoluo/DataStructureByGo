package linkedList

import (
	"bytes"
	"fmt"
)

type Node struct {
	elem interface{}
	next *Node
}

type LinkedList struct {
	dummyHead *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}

// 获取链表中的元素个数
func (list *LinkedList) GetSize() int {
	return list.size
}

// 返回链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// 在链表的index(0-size)位置添加元素
func (list *LinkedList) Add(index int, elem interface{}) {
	if index < 0 || index > list.size {
		panic("index illegal. Add failed!")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &Node{
		elem: elem,
		next: prev.next,
	}
	list.size++
}

// 在链表头添加元素
func (list *LinkedList) AddFirst(elem interface{}) {
	list.Add(0, elem)
}

// 在链表尾添加元素
func (list *LinkedList) AddLast(elem interface{}) {
	list.Add(list.size, elem)
}

// 获取链表的第index个位置的元素
func (list *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("index illegal. Get failed!")
	}
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.elem
}

// 获取链表头的元素
func (list *LinkedList) GetFirst() interface{} {
	return list.Get(0)
}

// 获取链表尾的元素
func (list *LinkedList) GetLast() interface{} {
	return list.Get(list.size-1)
}

// 修改链表的第index元素
func (list *LinkedList) Set(index int, elem interface{}) {
	if index < 0 || index >= list.size {
		panic("index illegal. Set failed!")
	}
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.elem = elem
}

// 查找链表是否存在元素e
func (list *LinkedList) Contains(elem interface{}) bool {
	cur := list.dummyHead
	for cur.next != nil {
		cur = cur.next
		if cur.elem == elem {
			return true
		}
	}
	return false
}

// 从链表中删除index位置的元素
func (list *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("index illegal. Remove failed!")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	ret := prev.next
	prev.next = ret.next
	ret.next = nil
	list.size--
	return ret.elem
}

// 从链表中删除第一个元素
func (list *LinkedList) RemoveFirst() interface{} {
	return list.Remove(0)
}

// 从链表中删除最后一个元素
func (list *LinkedList) RemoveLast() interface{} {
	return list.Remove(list.size-1)
}

// 将链表以字符串的形式打印
func (list *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := list.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.elem) + "->")
		cur = cur.next
	}
	buffer.WriteString("NULL")

	return buffer.String()
}

