package Problem_0078_Subsets

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	want := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	got := subsets(nums)
	fmt.Println(want)
	fmt.Println(got)
}
