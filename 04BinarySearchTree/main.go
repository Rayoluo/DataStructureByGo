package main

import (
	bst "04BinarySearchTree/binarySearchTree"
	"fmt"
)

func main() {
	b := bst.NewBST()
	nums := []int{2, 8, 5, 6, 1, 3, 9, 4, 7}
	for _, v := range nums {
		b.Add(v)
	}
	b.PreOrder()
	fmt.Println("==================")
	b.PreOrderNonRecursion()
	fmt.Println("==================")
	fmt.Println("==================")
	b.LevelOrder()
	fmt.Println("==================")
	fmt.Println(b)
	fmt.Println(b.Search(5))
	fmt.Println(b.Search(10))
}
