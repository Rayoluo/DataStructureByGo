package maxHeap

import (
	"01MaxHeap/util"
)

// 实现大顶堆

// 大顶堆底层就是一个完全二叉树，用切片来存储， 从索引为0处开始存储
type MaxHeap struct {
	arr  []interface{}
	size int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		arr: make([]interface{}, 0),
	}
}

// 将一个数组整成一个heap
func NewMaxHeapFromArr(sarr []interface{}) *MaxHeap {
	heap := &MaxHeap{
		arr:  make([]interface{}, len(sarr)),
		size: len(sarr),
	}
	for i := 0; i < len(sarr); i++ {
		heap.arr[i] = sarr[i]
	}
	// 从第一个非叶子节点开始进行sift down操作, 第一个非叶子节点即最后一个节点的parent
	for i := parent(heap.GetSize() - 1); i >= 0; i-- {
		heap.siftDown(i)
	}
	return heap
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) GetSize() int {
	return h.size
}

// 获取指定位置父节点的索引
func parent(k int) int {
	if k == 0 {
		panic("root of max heap has no parent")
	}
	return (k - 1) / 2
}

// 获取指定位置节点的左孩子
func leftChild(k int) int {
	return 2*k + 1
}

// 获取指定位置节点的右孩子
func rightChild(k int) int {
	return 2*k + 2
}

func (h *MaxHeap) Add(e interface{}) {
	h.arr = append(h.arr, e)
	h.size++
	// sift up
	h.siftUp(h.GetSize() - 1)
}

func (h *MaxHeap) siftUp(k int) {
	for k > 0 && util.Compare(h.arr[k], h.arr[parent(k)]) > 0 {
		// swap
		tmp := h.arr[parent(k)]
		h.arr[parent(k)] = h.arr[k]
		h.arr[k] = tmp
		k = parent(k)
	}
}

// 查看堆中的最大元素 findMax
func (h *MaxHeap) FindMax() interface{} {
	if h.GetSize() == 0 {
		panic("Heap is empty! Can not find max element")
	}
	return h.arr[0]
}

// 取出堆的最大元素 extractMax
func (h *MaxHeap) ExtractMax() interface{} {
	ret := h.FindMax()
	h.arr[0] = h.arr[h.GetSize()-1]
	h.arr = h.arr[:h.GetSize()-1]
	h.size--
	// sift down
	h.siftDown(0)
	return ret
}

func (h *MaxHeap) siftDown(k int) {
	for leftChild(k) < h.GetSize() {
		j := leftChild(k)
		// 这里粗心写错了，比较的不是索引，而是对应的元素值: util.Compare(leftChild(k), rightChild(k)) < 0
		if j+1 < h.GetSize() && util.Compare(h.arr[leftChild(k)], h.arr[rightChild(k)]) < 0 {
			j = rightChild(k)
		}
		// j is the index of max member in children
		if util.Compare(h.arr[k], h.arr[j]) > 0 {
			break
		} else {
			tmp := h.arr[k]
			h.arr[k] = h.arr[j]
			h.arr[j] = tmp
			k = j
		}
	}
}

// 返回堆中的最大元素，并插入一个新的元素e
func (h *MaxHeap) Replace(e interface{}) interface{} {
	ret := h.FindMax()
	h.arr[0] = e
	h.siftDown(0)
	return ret
}
