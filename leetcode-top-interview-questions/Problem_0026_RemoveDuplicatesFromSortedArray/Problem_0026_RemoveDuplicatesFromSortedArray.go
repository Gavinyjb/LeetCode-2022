package problem0026removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums)<2{
		return len(nums)
	}
	ret:=1
	for i := 1; i < len(nums); i++ {
		if nums[i]!=nums[i-1]{
			nums[ret]=nums[i]
			ret++
		}
	}
	return ret
}
