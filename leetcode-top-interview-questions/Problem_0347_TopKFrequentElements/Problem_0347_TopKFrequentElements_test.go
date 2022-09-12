package Problem_0347_TopKFrequentElements

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	nums := []int{3, 0, 1, 0}
	k := 1
	//nums := []int{3, 0, 1, 0}
	//k := 1
	fmt.Println(topKFrequent(nums, k))
}
