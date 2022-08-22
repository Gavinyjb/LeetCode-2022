package Problem_0078_Subsets

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums, []int{}}
	}
	path := make([]int, 0)
	var process func(nums []int, idx int, path []int)
	process = func(nums []int, idx int, path []int) {
		if idx == len(nums) {
			ans := make([]int, len(path))
			copy(ans, path)
			ret = append(ret, ans)
		} else {
			//不使用idx
			process(nums, idx+1, path)
			//使用idx
			path = append(path, nums[idx])
			process(nums, idx+1, path)
			path = path[:len(path)-1]
		}
	}
	process(nums, 0, path)
	return ret
}
