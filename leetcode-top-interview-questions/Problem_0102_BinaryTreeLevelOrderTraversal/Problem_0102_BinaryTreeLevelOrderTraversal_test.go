package Problem_0102_BinaryTreeLevelOrderTraversal

import "testing"

func TestName(t *testing.T) {
	rootInts := []int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}
	root := Ints2TreeNode(rootInts)
	want := [][]int{{3}, {9, 20}, {15, 7}}
	got := levelOrder(root)
	if len(want) != len(got) {
		t.Errorf("want:%v got:%v\n", want, got)
	} else {
		for i, list := range got {
			for j, _ := range list {
				if got[i][j] != want[i][j] {
					t.Errorf("want:%v got:%v\n", want, got)
				}
			}
		}
	}
}
