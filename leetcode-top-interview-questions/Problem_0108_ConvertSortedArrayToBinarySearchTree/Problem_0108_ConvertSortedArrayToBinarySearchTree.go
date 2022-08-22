package Problem_0108_ConvertSortedArrayToBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return process(nums, 0, len(nums)-1)
}
func process(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}
	mid := l + (r-l)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = process(nums, l, mid-1)
	root.Right = process(nums, mid+1, r)
	return root

}
