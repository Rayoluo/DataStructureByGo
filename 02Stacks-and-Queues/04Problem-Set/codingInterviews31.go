package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences1([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	mp := make(map[int]bool) // 标记元素是否已经压入栈
	stack := make([]int, 0)
	visitIndex := 0
	for i := 0; i < len(popped); i++ {
		// 如果该元素没有被压入栈，则压入序列中该元素之前的所有元素都将被压入栈中并被标记
		if _, ok := mp[popped[i]]; !ok {
			for j := visitIndex; j < len(pushed) && pushed[j] != popped[i]; j++ {
				stack = append(stack, pushed[j])
				mp[pushed[j]] = true
				visitIndex++
			}
			mp[popped[i]] = true
			visitIndex++
			continue
		} else {
			// 如果元素已被压入栈中，但不是栈顶元素，则返回false
			if topElem := stack[len(stack)-1]; topElem != popped[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return true
}

// 优化空间复杂度
func validateStackSequences1(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	visitIndex := 0
	for i := 0; i < len(popped); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
		} else {
			for j := visitIndex; j < len(pushed) && pushed[j] != popped[i]; j++ {
				stack = append(stack, pushed[j])
				visitIndex++
			}
			if visitIndex == len(pushed) {
				return false
			}
			visitIndex++
		}
	}
	return true
}
