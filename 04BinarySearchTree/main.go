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
	b.LevelOrder()
	fmt.Println("==================")
	fmt.Println(b)
	fmt.Println(b.Search(5))
	fmt.Println(b.Search(10))
	fmt.Println("=======测试Minimal函数==========")
	fmt.Println(b.Minimal())
	fmt.Println("=======测试RemoveMin函数==========")
	fmt.Println(b.RemoveMin())
	fmt.Println(b)
	fmt.Println("=======测试Remove函数(移除有单子树的节点)==========")
	b.Remove(6)
	fmt.Println(b)
	fmt.Println("=======测试Remove函数(移除有双子树的节点)==========")
	b.Remove(5)
	fmt.Println(b)
}
