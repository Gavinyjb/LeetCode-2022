package Problem_0046_Permutations

//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 1 {
		ret = append(ret, nums)
		return ret
	}
	for i := 0; i < len(nums); i++ {
		swap(nums, 0, i)
		first := nums[0]
		sub := make([]int, len(nums)-1)
		copy(sub, nums[1:])
		subRet := permute(sub)
		for _, ints := range subRet {
			subAns := make([]int, len(nums))
			subAns[0] = first
			copy(subAns[1:], ints)
			ret = append(ret, subAns)
		}
	}
	return ret
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
