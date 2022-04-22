package main

func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	ret := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(track []int, start int)
	backtrack = func(track []int, start int) {
		//终止条件
		if len(track) == k {
			temp := make([]int, len(track))
			copy(temp, track)
			ret = append(ret, temp)
		}
		if len(track)+n-start+1 < k {
			return
		}
		//for 选择列表 组合问题，选过了的不再重复选
		for i := start; i < n; i++ {
			//做选择
			track = append(track, nums[i])
			//进入下一层决策树
			backtrack(track, i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, 0)
	return ret
}

func main() {

}
