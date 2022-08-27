package Problem_0075_SortColors

func sortColors(nums []int) {
	p0, p2 := -1, len(nums)

	//all in [0,p0]  ==0
	//all in (p0,cur)  ==1
	//all in [p2,len]  ==2

	cur := 0
	for cur < p2 {
		if nums[cur] == 0 {
			p0++
			swap(nums, p0, cur)
			cur++
		} else if nums[cur] == 1 {
			cur++
		} else {
			p2--
			swap(nums, cur, p2)
		}
	}
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//leetcode 答案
func sortColorsLeetCode(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
