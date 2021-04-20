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
	leftToRight := true
	for len(queue) != 0 {
		count := 0
		line := make([]int, 0)
		stack := make([]*TreeNode, 0)
		// 将当前行的元素打印, 并将其子节点加入到队列中，统计下一行元素数目
		for currentLineNums != 0 {
			node := queue[0]
			line = append(line, node.Val)
			if leftToRight { // 从左到右打印当前行，则依次按照左孩子右孩子的顺序将下一行的子节点压入到栈中
				if node.Left != nil {
					stack = append(stack, node.Left)
					count++
				}
				if node.Right != nil {
					stack = append(stack, node.Right)
					count++
				}
			} else { // 从右到左打印当前行，则依次按照右孩子左孩子的新婚徐将下一行的子节点压入到栈中
				if node.Right != nil {
					stack = append(stack, node.Right)
					count++
				}
				if node.Left != nil {
					stack = append(stack, node.Left)
					count++
				}
			}
			// 弹出队列的首元素
			queue = queue[1:]
			currentLineNums--
		}
		lines = append(lines, line)
		currentLineNums = count
		if leftToRight {
			leftToRight = false
		} else {
			leftToRight = true
		}
		// 将栈中的元素弹出到队列中
		for len(stack) != 0 {
			queue = append(queue, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return lines
}
