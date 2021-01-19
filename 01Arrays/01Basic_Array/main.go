package main

import (
	"fmt"
	"github.com/algo/DataStructureByGo/01Arrays/01Basic_Array/array"
)

func main() {
	arr := array.NewArray(10)
	fmt.Printf("%T\n", arr)
	fmt.Println(arr.GetCapacity())
	fmt.Println(arr.GetSize())
	fmt.Println(arr.IsEmpty())
}
