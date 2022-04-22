package main

func combinationSum3(k int, n int) [][]int {
	nums := make([]int, 9)
	for i := 0; i < 9; i++ {
		nums[i] = i + 1
	}
	ret := make([][]int, 0)
	track := make([]int, 0)

	var backtrack func(track []int, start int)
	backtrack = func(track []int, start int) {
		//终止条件
		if len(track) == k {
			temp := make([]int, len(track))
			copy(temp, track)
			sum := 0
			for _, v := range temp {
				sum += v
			}
			if sum == n {
				ret = append(ret, temp)
			}
		}
		if len(track)+n-start+1 < k {
			return
		}
		for i := start; i < len(nums); i++ {
			//做选择
			track = append(track, nums[i])
			backtrack(track, i+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, 0)
	return ret
}

func main() {

}
