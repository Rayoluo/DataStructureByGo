package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// []int(nil)这种写法和var slice []int 效果相同
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var slice [][]int
	slice = findPath(root, target, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, target int, slice [][]int, stack []int) [][]int {
	stack = append(stack, n.Val)
	if n.Left == nil && n.Right == nil { // 到达了叶子节点
		if n.Val == target {
			slice = append(slice, append([]int{}, stack...))
		}
	}
	if n.Left != nil {
		slice = findPath(n.Left, target-n.Val, slice, stack)
	}
	if n.Right != nil {
		slice = findPath(n.Right, target-n.Val, slice, stack)
	}
	return slice
}
