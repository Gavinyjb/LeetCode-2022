package 排序

/**
 * @Author: yirufeng
 * @Date: 2021/4/13 8:28 下午
 * @Desc:

时间复杂度：O(n^2)
空间复杂度：O(1)
稳定的：因为是到前面找到一个小于等于自己的后面那个位置
 **/

func InsertSort(nums []int) {
	insertSort(nums)
}

func insertSort(nums []int) {
	var j int
	//第一个元素已经有序，只需要插入后面的元素即可
	for i := 1; i < len(nums); i++ {
		//当前要插入的元素目前位于i位置
		cur := nums[i]
		//从后往前找，直到前面元素小于当前元素就停，并且遍历的过程中不断将前面的元素移动到后面
		for j = i - 1; j >= 0 && nums[j] > cur; j-- {
			nums[j+1] = nums[j]
		}
		//此时j指向的元素就是前面那个小于或者等于当前元素的位置
		//但是我们应该在j+1位置插入
		nums[j+1] = cur
	}
}
