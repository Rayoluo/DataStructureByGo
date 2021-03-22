package main

/*
 * 题目描述：给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
 */
import (
	"container/heap"
	"fmt"
)

// index存放数据在堆（底层为切片）中的位置（索引）
type Item struct {
	val, index int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return (*h[i]).val > (*h[j]).val }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	(*h[i]).index, (*h[j]).index = (*h[j]).index, (*h[i]).index
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	queue := make([]*Item, 0)
	maxInEachWin := make([]int, 0)
	h := make(MaxHeap, k)
	// 初始化
	for i := 0; i < k; i++ {
		item := &Item{val: nums[i], index: i}
		h[i] = item
		queue = append(queue, item)
	}
	heap.Init(&h)
	maxInEachWin = append(maxInEachWin, (*h[0]).val)
	for i := k; i < len(nums); i++ {
		// 取出队首元素，即滑动窗口最左边的元素
		x := queue[0]
		x.val = nums[i]
		queue = queue[1:]
		queue = append(queue, x)
		heap.Fix(&h, x.index)
		maxInEachWin = append(maxInEachWin, (*h[0]).val)
	}
	return maxInEachWin
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
