package main

/*
 * 题目描述：数据流的中位数
 */

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	obj := Constructor()
	// test case 1
	// obj.AddNum(2)
	// fmt.Println(obj.FindMedian())
	// obj.AddNum(3)
	// fmt.Println(obj.FindMedian())
	// test case 2
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}

type MedianFinder struct {
	h1 MinHeap
	h2 MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		h1: make(MinHeap, 0),
		h2: make(MaxHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	m := len(this.h1)
	n := len(this.h2)
	// 如果最小堆和最大堆的元素数量相同，则先向最大堆中添加元素，并将最大堆的堆顶元素加入到最小堆中
	if m == n {
		heap.Push(&this.h2, num)
		x := heap.Pop(&this.h2)
		heap.Push(&this.h1, x.(int))
	} else if m == n+1 {
		heap.Push(&this.h1, num)
		x := heap.Pop(&this.h1)
		heap.Push(&this.h2, x.(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	m := len(this.h1)
	n := len(this.h2)
	if m == 0 || n == 0 {
		if m != 0 {
			return float64(this.h1[0])
		} else {
			return float64(this.h2[0])
		}
	}
	if m == n+1 {
		return float64(this.h1[0])
	} else {
		// fmt.Println(this.h1[0])
		// fmt.Println(this.h2[0])
		// return float64((this.h1[0] + this.h2[0]) / 2.0) // 这怕不是个坑。。
		// fmt.Println(float64(this.h1[0]+this.h2[0]) / 2)
		return float64(this.h1[0]+this.h2[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
