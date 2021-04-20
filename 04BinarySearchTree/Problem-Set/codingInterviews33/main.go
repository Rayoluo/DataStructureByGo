package main

func verifyPostorder(postorder []int) bool {
	return verifyFunc(postorder, 0, len(postorder)-1)
}

// 递归函数
func verifyFunc(postorder []int, left int, right int) bool {
	if left >= right {
		return true
	}
	// 取出最后一个元素作为二叉搜索树的根节点
	root := postorder[right]
	mid := left
	for ; mid < right; mid++ {
		if postorder[mid] > root {
			break
		}
	}
	if postorder[mid] > root {
		for j := mid + 1; j < right; j++ {
			if postorder[j] < root {
				return false
			}
		}
		return verifyFunc(postorder, left, mid-1) && verifyFunc(postorder, mid, right-1)
	}
	// postorder[mid] < root
	return verifyFunc(postorder, left, right-1)
}
