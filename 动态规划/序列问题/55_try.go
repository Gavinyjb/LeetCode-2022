package leetcode

func canJump(nums []int) bool {
	l:=len(nums)
	next:=0
	for i := 0; i < l-1; i++ {
		next=max(next,nums[i]+i)
		if next<=i{
			return false
		}
	}
	return next>=l-1
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
