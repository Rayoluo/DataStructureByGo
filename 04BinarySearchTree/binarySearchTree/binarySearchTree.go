package binarySearchTree

import (
	"04BinarySearchTree/util"
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

// attention: 默认二叉搜索树中没有重复的元素
type BST struct {
	root *Node
	size int
}

func newNode(e interface{}) *Node {
	return &Node{e: e}
}

// 新建一个二叉搜索树
func NewBST() *BST {
	return &BST{}
}

func (b *BST) GetSize() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0

}

// 向二分搜索树插入元素，一开始创建的是空结构体，root为nil
func (b *BST) Add(e interface{}) {
	b.root = addElement(b.root, e)
	b.size++
}

// addElement：向二叉搜索树中插入元素值，返回该二叉搜索树的根节点
func addElement(n *Node, e interface{}) *Node {
	if n == nil {
		return &Node{e: e}
	}
	if util.Compare(e, n.e) < 0 {
		n.left = addElement(n.left, e)
	} else if util.Compare(e, n.e) > 0 {
		n.right = addElement(n.right, e)
	}
	return n
}

// 查看二叉搜索树中是否包含元素e
func (b *BST) Search(e interface{}) bool {
	return searchElement(b.root, e)
}

// 在以n为根节点的二叉搜索树中查找元素e
func searchElement(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}
	if util.Compare(e, n.e) < 0 {
		return searchElement(n.left, e)
	} else if util.Compare(e, n.e) > 0 {
		return searchElement(n.right, e)
	}
	return true
}

// 实现二叉搜索树的前序遍历
func (b *BST) PreOrder() {
	preOrder(b.root)
}

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.e)
	preOrder(n.left)
	preOrder(n.right)
}

// 二叉搜素树前序遍历的非递归实现, 基于栈实现
func (b *BST) PreOrderNonRecursion() {
	if b.root == nil {
		fmt.Println("binary search tree's root is nil")
		return
	}
	stack := make([]*Node, 0)
	stack = append(stack, b.root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.e)
		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

// 实现二叉搜索树的后序遍历
func (b *BST) PostOrder() {
	postOrder(b.root)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.e)
}

// 实现二叉搜素树的中序遍历
func (b *BST) MidOrder() {
	midOrder(b.root)
}

func midOrder(n *Node) {
	if n == nil {
		return
	}
	midOrder(n.left)
	fmt.Println(n.e)
	midOrder(n.right)
}

// 二叉搜索树的层序遍历实现, 基于队列
func (b *BST) LevelOrder() {
	if b.root == nil {
		fmt.Println("binary search tree's root is nil")
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, b.root)
	for len(queue) != 0 {
		cur := queue[0]
		fmt.Println(cur.e)
		queue = queue[1:]
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}

// 返回二叉搜索树的最小值
func (b *BST) Minimal() *Node {
	return minimal(b.root)
}

func minimal(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n
	}
	return minimal(n.left)
}

// 移除二叉搜索树的最小值, 并返回该最小值
func (b *BST) RemoveMin() interface{} {
	// 首先获取元素的最小值
	min := b.Minimal()
	b.root = removeMin(b.root)
	b.size--
	return min.e
}

// 移除给定根节点的二叉搜索树的最小值，并返回该二叉搜索树的根节点
func removeMin(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n.right
	}
	n.left = removeMin(n.left)
	return n
}

// 移除二叉搜素树中的指定元素
func (b *BST) Remove(e interface{}) {
	b.root = remove(b.root, e)
	b.size--
}

// 删除二叉搜索树中的指定元素，并返回该二叉搜索树的根节点
func remove(n *Node, e interface{}) *Node {
	if n == nil {
		return nil
	}
	if util.Compare(e, n.e) < 0 {
		n.left = remove(n.left, e)
		return n
	} else if util.Compare(e, n.e) > 0 {
		n.right = remove(n.right, e)
		return n
	}
	// 当前元素即为要被删除的元素，分为两种情况
	// 1. 只有左子树或者右子树，这时候只需要返回左子树或右子树即可
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		return rightNode
	}
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		return leftNode
	}
	// 2. 同时有左子树和右子树, 接下来的第一步是先找到右子树的最小值
	newNode := minimal(n.right)
	newNode.right = removeMin(n.right)
	newNode.left = n.left
	return newNode
}

func (b *BST) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("size of BST: " + strconv.Itoa(b.size) + "\n")
	generateBSTSting(b.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
