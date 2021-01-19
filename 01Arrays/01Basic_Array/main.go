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
}
