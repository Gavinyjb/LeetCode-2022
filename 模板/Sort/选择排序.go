package 排序

/**
 * @Author: yirufeng
 * @Date: 2021/4/13 8:28 下午
 * @Desc: 选择排序

时间复杂度：假设数组中有n个数字
	大循环需要遍历n-2趟，然后每一个大循环中，
							我们都需要从开始的数字一直到最后选出最小的数字(时间复杂度：O(n))，然后填充到当期那还没有排好序的起始位置

最后总的时间复杂度：O(n^2)，
空间复杂度：O(1)。
稳定的排序算法（因为两个相等的数字前面那个一定是排序的时候第一个会被放置到前面的）
 **/

func SelectedSort(nums []int) {
	selectedSort(nums)
}

func selectedSort(nums []int) {
	var min int
	for i := 0; i < len(nums)-1; i++ {
		//每次从nums[i]开始起
		min = i
		//依次与后面比对来选择最小的
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		//交换nums[min]与nums[i]
		nums[i], nums[min] = nums[min], nums[i]
	}
}
