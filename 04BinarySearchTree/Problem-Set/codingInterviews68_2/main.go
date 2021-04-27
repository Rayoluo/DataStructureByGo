package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pi, qi   int
	midSlice []int
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	midSlice = make([]int, 0)
	midOrderTraverse(root, &midSlice)
	// ri = findIndex(midSlice, root.Val)
	pi = findIndex(midSlice, p.Val)
	qi = findIndex(midSlice, q.Val)
	return lca(root)
}

func lca(root *TreeNode) *TreeNode {
	ri := findIndex(midSlice, root.Val)
	if ri > pi && ri > qi {
		return lca(root.Left)
	}
	if ri < pi && ri < qi {
		return lca(root.Right)
	}
	return root
}

func findIndex(arr []int, val int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func midOrderTraverse(root *TreeNode, midSlice *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		midOrderTraverse(root.Left, midSlice)
	}
	*midSlice = append(*midSlice, root.Val)
	if root.Right != nil {
		midOrderTraverse(root.Right, midSlice)
	}
}
