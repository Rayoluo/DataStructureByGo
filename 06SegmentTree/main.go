package main

import (
	"06SegmentTree/segmentTree"
	"fmt"
)

type NumArray struct {
	*segmentTree.SegmentTree
}

func Constructor(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	numArr := NumArray{
		segmentTree.New(data, func(i interface{}, j interface{}) interface{} {
			return i.(int) + j.(int)
		}),
	}
	return numArr
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(left, right).(int)
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numarr := Constructor(nums)
	fmt.Println(numarr.SumRange(0, 2))
	fmt.Println(numarr.SumRange(2, 5))
	fmt.Println(numarr.SumRange(0, 5))
	// fmt.Println(numarr.SumRange(0, 5))
}
