package Problem_0145_BinaryTreePostOrderTraversal

import "testing"

func TestName(t *testing.T) {
	rootInts := []int{1, -1 << 63, 2, 3}
	root := Ints2TreeNode(rootInts)
	want := []int{3, 2, 1}
	got := postorderTraversal(root)
	if len(want) != len(got) {
		t.Errorf("want:%v got:%v\n", want, got)
	} else {
		for i, _ := range got {
			if got[i] != want[i] {
				t.Errorf("want:%v got:%v\n", want, got)
			}
		}
	}

}
