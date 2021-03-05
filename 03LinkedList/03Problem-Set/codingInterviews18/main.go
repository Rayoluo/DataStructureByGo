package main

import "fmt"

/*
 * 题目描述：删除链表中重复的节点
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

func deleteNode(head *ListNode, val int) *ListNode {
	var (
		dummyHead, cur *ListNode
	)
	dummyHead = &ListNode{
		Val: -1,
		Next: head,
	}
	cur = dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	cur := deleteNode(NewLinkedList([]int{4, 5, 1, 9, 5}), 5)
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("NULL")
}