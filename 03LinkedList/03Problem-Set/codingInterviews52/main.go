package main

/*
 * 题目描述：找到两个链表的第一个公共节点
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双栈的思路
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	s1 := make([]*ListNode, 0)
	s2 := make([]*ListNode, 0)
	var (
		cur, top1, top2, prev *ListNode
	)
	cur = headA
	for cur != nil {
		s1 = append(s1, cur)
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		s2 = append(s2, cur)
		cur = cur.Next
	}
	for len(s1) != 0 && len(s2) != 0 {
		top1 = s1[len(s1)-1]
		top2 = s2[len(s2)-1]
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
		if top1 != top2 {
			return prev
		} else {
			prev = top1
		}
	}
	return prev
}

// 双指针的思路 a + c + b = a + b + c
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pList1, pList2 := headA, headB
	for pList1 != pList2 {
		if pList1 != nil {
			pList1 = pList1.Next
		} else {
			pList1 = headB
		}
		if pList2 != nil {
			pList2 = pList2.Next
		} else {
			pList2 = headA
		}
	}
	return pList1
}
