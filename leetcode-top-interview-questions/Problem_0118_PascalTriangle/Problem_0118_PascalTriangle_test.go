package Problem_0118_PascalTriangle

import (
	"fmt"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	numRows := 5
	got := generate(numRows)
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	fmt.Println(got)
	fmt.Println(want)
}
