package Problem_0239_maxSlidingWindow

import (
	"fmt"
	"testing"
)

func TestMine(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	got := maxSlidingWindow(nums, k)
	want := []int{3, 3, 5, 5, 6, 7}
	fmt.Println("want", want)
	fmt.Println("got", got)
}
func TestMine2(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	got := maxSlidingWindow(nums, k)
	want := []int{3, 3, 5, 5, 6, 7}
	fmt.Println("want", want)
	fmt.Println("got", got)
}
