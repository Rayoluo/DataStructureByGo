package main

import (
	"bytes"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据前序遍历和中序遍历结果重建二叉树, 前序遍历和中序遍历的结果都不包含重复数字
// 犹豫不决，先写注释
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || preorder == nil {
		return nil
	}
	// 如果只有一个节点，就直接返回该节点
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	// 否则，先拿出子树的根节点，再递归构建该节点的左右子树
	var (
		root *TreeNode
		k    int
	)
	// 在中序遍历的结果中查找该节点, 返回索引值k，则左子树有k个节点,右子树有len(inorder)-k-1个节点
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			k = i
			break
		}
	}
	root = &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:k+1], inorder[:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])
	return root
}

func main() {
	// test := []int{1, 2, 3}
	// fmt.Println(test[1:1]) // []

	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	var buffer bytes.Buffer
	generateBSTSting(root, 0, &buffer)
	fmt.Print(buffer.String())
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *TreeNode, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.Val) + "\n")
	generateBSTSting(n.Left, depth+1, buffer)
	generateBSTSting(n.Right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
