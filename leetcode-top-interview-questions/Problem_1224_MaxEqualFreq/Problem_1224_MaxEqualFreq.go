package Problem_1224_MaxEqualFreq

func maxEqualFreq(nums []int) int {

}

//返回值：要删除的数字 和答案
func process(nums []int, idx int, del int, ans int) {
	if idx == 0 {
		return nums[0], 1
	}
	if nums[idx] == del {
		return
	}
}
