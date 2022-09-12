package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDemo(t *testing.T) {
	n := 20
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(20))
	}
	want := make([]int, len(nums))
	copy(want, nums)
	got := HeapSort(nums)
	fmt.Println(got)
	fmt.Println(want)
}
