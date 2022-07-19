package leetcode

func search(nums []int, target int) int {
	//1.初始化start,end
	start := 0
	end := len(nums) - 1
	//2. 处理for循环
	for start+1 < end {
		mid := start + (end-start)>>1
		//3. 比较a[mid]和target值
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
