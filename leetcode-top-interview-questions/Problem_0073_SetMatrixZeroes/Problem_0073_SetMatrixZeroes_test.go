package Problem_0073_SetMatrixZeroes

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	showMatrix(matrix)
	setZeroes(matrix)
	showMatrix(matrix)
}
func showMatrix(matrix [][]int) {
	for _, ints := range matrix {
		fmt.Println(ints)
	}
}
