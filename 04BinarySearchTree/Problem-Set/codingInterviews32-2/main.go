package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lines := make([][]int, 0)
	// 表示当前行有一个元素
	currentLineNums := 1
	for len(queue) != 0 {
		count := 0
		line := make([]int, 0)
		// 将当前行的元素打印, 并将其子节点加入到队列中，统计下一行元素数目
		for currentLineNums != 0 {
			node := queue[0]
			line = append(line, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				count++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				count++
			}
			// 弹出队列的首元素
			queue = queue[1:]
			currentLineNums--
		}
		lines = append(lines, line)
		currentLineNums = count
	}
	return lines
}
