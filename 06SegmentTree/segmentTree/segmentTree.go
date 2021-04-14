package segmentTree

import (
	"bytes"
	"fmt"
)

type SegmentTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(interface{}, interface{}) interface{}
}

// 构建一个线段树
func New(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		tree:   make([]interface{}, 4*len(arr)),
		data:   make([]interface{}, len(arr)),
		merger: merger,
	}
	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}
	// 构建二叉树
	segmentTree.buildSegmentTree(0, 0, len(arr)-1)
	return segmentTree
}

// 在treeIndex位置构建表示区间[l,r]的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	mid := l + (r-l)>>1
	st.buildSegmentTree(leftChild(treeIndex), l, mid)
	st.buildSegmentTree(rightChild(treeIndex), mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftChild(treeIndex)], st.tree[rightChild(treeIndex)])
}

// 获取线段树的大小
func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

// 获取线段树index索引处的元素
func (st *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("index is illegal.")
	}
	return st.data[index]
}

// 返回区间[queryL, queryR]的值
func (st *SegmentTree) Query(queryL, queryR int) interface{} {
	if queryL < 0 || queryL >= len(st.data) || queryR < 0 || queryR >= len(st.data) || queryL > queryR {
		panic("query value is illegal.")
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以treeIndex为根的线段树中[l,r]范围内，搜索区间[queryL, queryR]的值
func (st *SegmentTree) query(treeIndex int, l int, r int, queryL int, queryR int) interface{} {
	if queryL == l && queryR == r {
		return st.tree[treeIndex]
	}
	mid := l + (r-l)>>1
	leftchild := leftChild(treeIndex)
	rightchild := rightChild(treeIndex)
	if queryR <= mid {
		return st.query(leftchild, l, mid, queryL, queryR)
	} else if queryL > mid {
		return st.query(rightchild, mid+1, r, queryL, queryR)
	} else {
		leftValue := st.query(leftchild, l, mid, queryL, mid)
		rightValue := st.query(rightchild, mid+1, r, mid+1, queryR)
		return st.merger(leftValue, rightValue)
	}
}

// 线段树的更新
func (st *SegmentTree) Update(index int, e interface{}) {
	if index < 0 || index >= len(st.data) {
		panic("Update: index is illegal.")
	}
	st.data[index] = e
	st.update(0, 0, len(st.data)-1, index, e)
}

func (st *SegmentTree) update(treeIndex int, l int, r int, index int, e interface{}) {
	if l == r {
		st.tree[treeIndex] = e
		return
	}
	mid := l + (r-l)>>1
	leftchild := leftChild(treeIndex)
	rightchild := rightChild(treeIndex)
	if index <= mid {
		st.update(leftchild, l, mid, index, e)
	} else {
		st.update(rightchild, mid+1, r, index, e)
	}
	st.tree[treeIndex] = st.merger(st.tree[leftchild], st.tree[rightchild])
}

func (st *SegmentTree) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < len(st.tree); i++ {
		if st.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(st.tree[i]))
		} else {
			buffer.WriteString("nil")
		}
		if i != len(st.tree)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func parent(index int) int {
	return (index - 1) / 2
}
