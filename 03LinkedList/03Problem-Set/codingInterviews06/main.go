package main

import (
	"fmt"
)

/**
 * 题目描述：输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func NewLinkedList(input []int) *ListNode {
	if input == nil || len(input) == 0 {
		return nil
	}
	dummyHead := &ListNode{-1, nil}
	cur := dummyHead
	for i := 0; i < len(input); i++ {
		cur.Next = &ListNode{input[i], nil}
		cur = cur.Next
	}
	return dummyHead.Next
}

// 顺序遍历存到数组中，然后逆序输出
func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = arr[len(arr) - 1 - i]
	}
	return ret
}

// 使用栈
func reversePrint1(head *ListNode) []int {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	ret := make([]int, len(stack))
	for i := 0; i < len(ret); i++ {
		ret[i] = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return ret
}

// 使用递归
func reversePrint2(head *ListNode) []int {
	return subFunc(head)
}

func subFunc(head *ListNode) []int {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	if head.Next == nil {
		arr = append(arr, head.Val)
		return arr
	}
	arr = append(arr, subFunc(head.Next)...)
	arr = append(arr, head.Val)
	return arr
}

func main() {
	fmt.Println(reversePrint2(NewLinkedList([]int{})))
}
