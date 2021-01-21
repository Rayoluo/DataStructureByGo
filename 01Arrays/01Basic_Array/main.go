package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/01Arrays/01Basic_Array/array"
)

func main() {
	arr := array.NewArray(10)
	fmt.Println(arr.GetCapacity())
	fmt.Println(arr.GetSize())
	fmt.Println(arr.IsEmpty())

	for i := 0; i < 6; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr.ArrayToString())

	// 在第2个元素位置添加一个元素6
	arr.AddIndex(6, 2)
	fmt.Println(arr.ArrayToString())
	// 在第0个元素位置添加一个元素9
	arr.AddFirst(9)
	fmt.Println(arr.ArrayToString())

	// 获取第6个位置的元素值
	fmt.Println("第6个位置的元素值是:", arr.Get(6))
	// 修改第6个位置的元素值为-1
	arr.Set(6, -1)
	fmt.Println(arr.ArrayToString())

	// 判断数组是否包含元素1
	fmt.Println(arr.Contains(1))
	// 查找数组中元素3的索引值
	fmt.Println(arr.Find(3))
	// 删除数组中的元素3
	// arr.DeleteIndex(arr.Find(3))
	fmt.Println("删除的元素是:", arr.DeleteIndex(arr.Find(3)))
	fmt.Println(arr.ArrayToString())
	fmt.Println("删除数组第一个元素:", arr.DeleteFirst())
	fmt.Println("删除数组最后一个元素:", arr.DeleteLast())
	fmt.Println(arr.ArrayToString())
}
