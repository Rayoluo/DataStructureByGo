package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, res)
	*res = append(*res, root.Val)
	midOrder(root.Right, res)
}

func kthLargest(root *TreeNode, k int) int {
	res := make([]int, 0)
	midOrder(root, &res)
	return res[len(res)-k]
}
