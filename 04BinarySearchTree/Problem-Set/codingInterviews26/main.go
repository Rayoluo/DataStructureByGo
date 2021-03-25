package main

/*
 * 题目描述：输入两棵二叉树A和B，判断B是不是A的子结构
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isEqual(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isEqual(t1, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil || t1.Val != t2.Val {
		return false
	}
	return isEqual(t1.Left, t2.Left) && isEqual(t1.Right, t2.Right)
}
