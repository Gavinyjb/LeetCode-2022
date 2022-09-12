package offer40

import (
	"fmt"
	"testing"
)

func TestGetLessK(t *testing.T) {
	arr := []int{0, 1, 2, 1}
	k := 1
	got := getLeastNumbers(arr, k)
	fmt.Println(got)
}
