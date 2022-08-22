package Problem_0098_ValidateBinarySearchTree

import "testing"

func TestMorris(t *testing.T) {
	rootInts := []int{5, 1, 4, -1 << 63, -1 << 63, 3, 6}
	root := Ints2TreeNode(rootInts)
	want := false
	got := isValidBST(root)
	if got != want {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
