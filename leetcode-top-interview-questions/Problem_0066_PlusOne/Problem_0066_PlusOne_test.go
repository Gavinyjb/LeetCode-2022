package Problem_0066_PlusOne

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {0}, {4, 3, 2, 1}, {9}, {8, 9, 9, 9}}
	for _, input := range inputs {
		fmt.Println(input)
		got := plusOne(input)
		fmt.Println(got)
		fmt.Println("----------------")
	}
}
