package main

import "fmt"

/*
 * 题目描述：给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 思路：首先判断链表是否有环，方法是快慢指针；
 * 接下来计算链表环中有多少个结点（count），方法是选定环中的一个结点然后计数，直到再次回到当前结点；
 * 最后是找到链表环的入口结点，方法是双指针，一个指针先走count步，然后以相同的速度往前走，
 * 最终两个指针将在链表环的入口处相遇；
 */
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	markNode := HasLoop(pHead)
	if markNode == nil {
		return nil
	}
	count := CountNodesInLoop(markNode)
	slow := pHead
	fast := pHead
	for i := 0; i < count; i++ {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 判断链表是否有环，有环则返回环中的节点，否则返回nil
func HasLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	slow := pHead
	fast := pHead
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

// 计算链表环中节点的个数
func CountNodesInLoop(markNode *ListNode) int {
	cur := markNode.Next
	count := 1
	for cur != markNode {
		count++
		cur = cur.Next
	}
	return count
}

// 构建一个带环的链表, 并指定输入数组的第k个元素作为环的入口结点, k从0开始
func NewLoopLinkedList(arr []int, k int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	if k < 0 || k >= len(arr)-1 {
		panic("illegal k value!")
	}
	dummyHead := &ListNode{-1, nil}
	var enNode, cur *ListNode
	cur = dummyHead
	for i := 0; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		cur = cur.Next
		if i == k {
			enNode = cur
		}
	}
	cur.Next = enNode
	return dummyHead.Next
}

// 生成不带环的链表
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
	fmt.Println("生成带环的链表，测试HasLoop函数：")
	fmt.Println(HasLoop(NewLoopLinkedList([]int{1, 2, 3, 4, 5, 6}, 2)).Val)
	fmt.Println("生成不带环的链表，测试HasLoop函数：")
	fmt.Println(HasLoop(NewLinkedList([]int{1, 2, 3, 4, 5, 6})))
	fmt.Println("带环的链表，测试EntryNodeOfLoop函数")
	fmt.Println(EntryNodeOfLoop(NewLoopLinkedList([]int{1, 2, 3, 4, 5, 6}, 2)))
	fmt.Println("不带环的链表，测试EntryNodeOfLoop函数")
	fmt.Println(EntryNodeOfLoop(NewLinkedList([]int{1, 2, 3, 4, 5, 6})))
}
