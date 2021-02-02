package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
	fmt.Println(spiralOrder([][]int{{}}))
}

// 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	column, row := len(matrix[0]), len(matrix)
	var retList []int = make([]int, 0)
	ix, iy := 0, 0
	for column > 0 && row > 0 {
		for i := 0; i < column; i++ {
			retList = append(retList, matrix[ix][iy+i])
		}
		for i := 1; i < row; i++ {
			retList = append(retList, matrix[ix+i][iy+column-1])
		}
		if column != 1 && row != 1 {
			for i := 1; i < column; i++ {
				retList = append(retList, matrix[ix+row-1][iy+column-1-i])
			}
			for i := 1; i < row-1; i++ {
				retList = append(retList, matrix[ix+row-1-i][iy])
			}
		}
		ix++
		iy++
		column -= 2
		row -= 2
	}
	return retList
}
