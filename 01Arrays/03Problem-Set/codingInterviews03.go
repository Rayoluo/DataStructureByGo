package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber2(nums))
}

// 统计所有数出现的次数
func findRepeatNumber(nums []int) int {
	times := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		times[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if times[i] > 1 {
			return i
		}
	}
	return -1
}

// 使用集合
func findRepeatNumber1(nums []int) int {
	var mp map[int]bool = make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = true
		} else {
			return nums[i]
		}
	}
	return -1
}

// 原地交换方法, 鸽巢原理
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		} else {
			val := nums[i]
			if nums[val] == val {
				return val
			} else {
				nums[i], nums[val] = nums[val], nums[i]
			}
		}
	}
	return -1
}
