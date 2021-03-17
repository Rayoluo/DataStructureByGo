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

// 二叉搜素树前序遍历的非递归实现

// 实现二叉搜索树的后序遍历

// 实现二叉搜素树的中序遍历

// 二叉搜索树的层序遍历实现

// 移除二叉搜索树的最大最小值

// 移除二叉搜素树中的指定元素

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
