package Problem_0075_SortColors

import (
	"math/rand"
	"testing"
)

func TestSortColor(t *testing.T) {
	n := 20
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(2))
	}
	want := make([]int, len(nums))
	copy(want, nums)
	sortColors(nums)
	sortColorsLeetCode(want)
	for i, _ := range want {
		if want[i] != nums[i] {
			t.Errorf("want:%v\n got:%v\n", want, nums)
		}
	}
}
