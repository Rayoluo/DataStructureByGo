package main

/*
 * 题目描述：最小的k个数
 */

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || arr == nil {
		return nil
	}
	var h IntHeap = make(IntHeap, len(arr))
	for i := 0; i < len(arr); i++ {
		h[i] = arr[i]
	}
	heap.Init(&h)
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = heap.Pop(&h).(int)
	}
	return ret
}

// 快排的方法也能做，先留个坑

func main() {
	// todo
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 3))
}
