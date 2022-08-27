package Problem_0912_SortAnArray

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	n := 20
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(20))
	}
	want := make([]int, len(nums))
	copy(want, nums)
	quickSortMine(nums, 0, len(nums)-1)
	quickSort(want, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println(want)

}

func checkSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
