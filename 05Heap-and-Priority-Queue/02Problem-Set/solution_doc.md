# 剑指offer堆类型题目总结

## 面试题41 数据流的中位数

题目链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/

题目描述；如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

设计一个支持以下两种操作的数据结构：

+ `void addNum(int num)`-从数据流中添加一个整数到数据结构中；
+ `double findMedian()`-返回目前所有元素的中位数

题解：

使用两个堆，一个大顶堆一个小顶堆，小顶堆元素数目为m，大顶堆元素数目为n，如果数据流元素总数为偶数，则m=n，如果数据流元素总数为奇数，则m=n+1；小顶堆中存放的是数据流中较大的一部分数字，大顶堆存放的是数据流中较小的一部分数字。

算法步骤：

1. 每插入一个数之前，先判断两个堆的元素数目是否相等；
2. 如果m=n，则先将这个数插入到大顶堆，然后将大顶堆的堆顶元素弹出并插入到小顶堆中。这样就可以保证小顶堆的所有数永远大于等于大顶堆的堆顶元素；
3. 如果m不等于n，即m=n+1，则先将这个数插入到小顶堆中，然后将小顶堆的堆顶元素弹出并插入到大顶堆中。这样就可以保证大顶堆的所有数永远小于等于小顶堆的堆顶元素；

代码实现：

```go
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
```

## 面试题59 滑动窗口的最大值

题目链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

题目描述：给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

<img src="https://gitee.com/oluoluo/pic_bed/raw/master/img/image-20210323110229050.png" alt="image-20210323110229050" style="zoom:67%;" />

题解：使用大顶堆来解决的思路比较简单，维护一个有k个元素的大顶堆（包含索引），同时维护一个队列，记录元素的先后顺序，每次滑动窗口向前滑动一个位置，取出队首元素得到堆中待移除元素的索引，然后将该元素替换成滑动窗口的后一个元素，之后再执行Fix操作。

代码实现：

```go
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
```

另外一种思路更为巧妙，即在栈和队列部分做过的题目，最大栈的思路。使用一个双端队列维护数据队列的每一步操作的最大值。如下图所示。

 [动画演示 单调队列 剑指 Offer 59 - I. 滑动窗口的最大值 - 滑动窗口的最大值 - 力扣（LeetCode）.ts](动画演示 单调队列 剑指 Offer 59 - I. 滑动窗口的最大值 - 滑动窗口的最大值 - 力扣（LeetCode）.ts) 