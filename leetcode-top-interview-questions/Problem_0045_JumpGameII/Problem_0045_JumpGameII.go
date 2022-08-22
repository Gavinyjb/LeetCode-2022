package Problem_0045_JumpGameII

//输入: nums = [2,3,1,1,4]
//输出: 2

func jump(nums []int) int {
	return dpFunc(nums, 0, len(nums)-1, 0)
}
func dpFunc(nums []int, idx int, end int, count int) int {
	dp := make([]int, len(nums))
	dp[len(dp)-1] = 0
	for idx := len(dp) - 2; idx >= 0; idx-- {
		if nums[idx] == 0 {
			dp[idx] = len(nums) //不可达
		}
		res := len(nums) - 1
		for i := 1; i <= nums[idx] && idx+i < len(nums); i++ {
			res = min(1+dp[idx+i], res)
		}
		dp[idx] = res
	}
	return dp[0]
}

//暴力尝试
func jump2(nums []int) int {
	return try(nums, 0, len(nums)-1)
}
func try(nums []int, idx int, end int) int {
	if idx == end {
		return 0
	} else {
		if nums[idx] == 0 {
			return len(nums) //说明不可达  如果可达最多也只需要len(nums)-1
		}
		res := len(nums) - 1
		for i := 1; i <= nums[idx] && idx+i < len(nums); i++ {
			res = min(1+try(nums, idx+i, end), res)
		}
		return res
	}
}

//func try(nums []int,idx int,end int,count int)int  {
//	if idx==end{
//		return count
//	}else {
//		if nums[idx]==0{
//			return len(nums)-1
//		}
//		res:=len(nums)-1
//		for i := 1; i <= nums[idx]&&idx+i<len(nums); i++ {
//			res=min(try(nums,idx+i,end,count+1),res)
//		}
//		return res
//	}
//}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
