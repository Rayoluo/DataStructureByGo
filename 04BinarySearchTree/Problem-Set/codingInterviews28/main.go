package main

/*
 * 题目描述：判断二叉树是否为对称的
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 检查左子树的镜像是否和右子树相同
	return check(mirrorTree(root.Left), root.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	t1 := mirrorTree(root.Left)
	t2 := mirrorTree(root.Right)
	root.Left = t2
	root.Right = t1
	return root
}

// 判断两棵树是否完全相同
func check(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Left) && check(a.Right, b.Right)
}
