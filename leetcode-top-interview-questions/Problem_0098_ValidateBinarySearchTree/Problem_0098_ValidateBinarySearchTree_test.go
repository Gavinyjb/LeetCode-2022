package Problem_0098_ValidateBinarySearchTree

import "testing"

func TestMorris(t *testing.T) {
	rootInts := []int{5, 1, 4, -1 << 63, -1 << 63, 3, 6}
	root := Ints2TreeNode(rootInts)
	want := false
	got := isValidBST3(root)
	if got != want {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
func TestBST3(t *testing.T) {
	//rootInts := []int{1, 1}
	rootInts := []int{5, 4, 6, -1 << 63, -1 << 63, 3, 7}
	//rootInts := []int{5, 1, 4, -1 << 63, -1 << 63, 3, 6}
	root := Ints2TreeNode(rootInts)
	want := false
	got := isValidBST3(root)
	if got != want {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
