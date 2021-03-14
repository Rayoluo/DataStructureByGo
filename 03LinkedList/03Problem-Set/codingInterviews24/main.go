package main

/*
 * 题目描述：反转链表，输入一个链表的头节点，反转该链表并输出反转后链表的头节点
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代的方式反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		prev, cur, post *ListNode
	)
	prev = nil
	cur = head
	for cur.Next != nil {
		post = cur.Next
		cur.Next = prev
		prev = cur
		cur = post
	}
	cur.Next = prev
	return cur
}
