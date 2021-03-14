package main

/*
 * 合并两个有序的链表：输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代的写法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		}
		return l1
	}
	head := &ListNode{-1, nil}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return head.Next
}
