package main

/*
 * 题目描述：请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 暴力方法，不用动脑子
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	valueList := make([]int, 0)
	addrList := make([]*Node, 0)
	randomIndex := make([]int, 0)
	cur := head
	for cur != nil {
		valueList = append(valueList, cur.Val)
		addrList = append(addrList, cur)
		cur = cur.Next
	}
	for cur = head; cur != nil; cur = cur.Next {
		if cur.Random == nil {
			randomIndex = append(randomIndex, -1)
		} else {
			for i := 0; i < len(addrList); i++ {
				if addrList[i] == cur.Random {
					randomIndex = append(randomIndex, i)
					break
				}
			}
		}
	}
	dummyHead := &Node{-1, nil, nil}
	newAddrList := make([]*Node, 0)
	cur = dummyHead
	for i := 0; i < len(valueList); i++ {
		cur.Next = &Node{valueList[i], nil, nil}
		newAddrList = append(newAddrList, cur.Next)
		cur = cur.Next
	}
	cur = dummyHead
	for i := 0; i < len(valueList); i++ {
		if randomIndex[i] == -1 {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = newAddrList[randomIndex[i]]
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

// 另一种方法使用哈希表，考虑构建原链表节点和新链表对应节点的键值对映射关系
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	var (
		dummyHead, cur, newCur *Node
	)
	mp := make(map[*Node]*Node)
	dummyHead = &Node{-1, nil, nil}
	cur, newCur = head, dummyHead
	for cur != nil {
		newCur.Next = &Node{cur.Val, nil, nil}
		newCur = newCur.Next
		mp[cur] = newCur
		cur = cur.Next
	}
	cur, newCur = head, dummyHead.Next
	for cur != nil {
		newCur.Random = mp[cur.Random]
		cur, newCur = cur.Next, newCur.Next
	}
	return dummyHead.Next
}

// 第三种方法是拼接+拆分, 知道思路就可以了，这里就不实现了，主要是太懒了，参考笔记
