package Problem_0041_FirstMissingPositive

func firstMissingPositive(nums []int) int {
	l, r := 0, len(nums)-1 //l左侧是有序的  r右边是垃圾区
	for l <= r {
		//fmt.Println(nums)
		//fmt.Println(l,r)
		if nums[l] == l+1 { //不交换
			l++
		} else if nums[l] < l+1 || nums[l] > r+1 || nums[nums[l]-1] == nums[l] {
			//放到垃圾区
			swap(nums, l, r)
			r--
		} else {
			swap(nums, l, nums[l]-1)
		}
	}
	//fmt.Println(l,r)
	return l + 1
}
func swap(nums []int, l, r int) {
	nums[l], nums[r] = nums[r], nums[l]
}
