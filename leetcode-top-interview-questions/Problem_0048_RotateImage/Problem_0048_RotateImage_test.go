package Problem_0048_RotateImage

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	show(matrix)
	fmt.Println("<---------------->")
	rotate(matrix)
	show(matrix)
}
func show(nums [][]int) {
	for _, num := range nums {
		for _, v := range num {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
