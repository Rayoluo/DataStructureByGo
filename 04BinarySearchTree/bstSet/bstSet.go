package bstSet

import (
	"04BinarySearchTree/binarySearchTree"
)

type BstSet struct {
	*binarySearchTree.BST
}

func NewBstSet() *BstSet {
	return &BstSet{
		binarySearchTree.NewBST(),
	}
}

func (s *BstSet) Contains(e interface{}) bool {
	return s.Search(e)
}

func (s *BstSet) IsEmpty() bool {
	return s.BST.IsEmpty()
}

func (s *BstSet) GetSize() int {
	return s.BST.GetSize()
}

func (s *BstSet) Add(e interface{}) {
	s.BST.Add(e)
}

func (s *BstSet) Remove(e interface{}) {
	s.BST.Remove(e)
}
