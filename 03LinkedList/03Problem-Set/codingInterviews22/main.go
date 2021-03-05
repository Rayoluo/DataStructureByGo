package main

import (
	"fmt"
)

/*
 * 题目描述：输入一个链表，输出该链表中倒数第k个节点，假定链表的尾节点是倒数第1个节点
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针方法
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	p1 := head
	p2 := head
	for i := 0; i < k-1 && p2 != nil; i++ {
		p2 = p2.Next
	}
	if p2 == nil {
		return nil
	}
	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1
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

func main() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	//fmt.Println("倒数第k个值为：" + strconv.Itoa(getKthFromEnd(head, 6).Val))
	fmt.Println(getKthFromEnd(head, 0))
}
