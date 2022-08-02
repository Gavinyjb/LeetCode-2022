package 模板

import "math/rand"

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l < r {
		//如果这里加上一步：将随机选中的一个元素与最后一个元素交换就是我们的随机快排
		index := rand.Intn(r-l+1) + l
		nums[index], nums[r] = nums[r], nums[index]
		l1, r1 := partition(nums, l, r)
		quickSort(nums, l, l1-1)
		quickSort(nums, r1+1, len(nums)-1)
	}
}

//使用nums[R]作为划分
func partition(nums []int, L, R int) (int, int) {
	//最开始没有小于等于范围没有数字，
	lessIndex := L - 1
	largerIndex := R
	num := nums[R]
	cur := L
	for cur < largerIndex {
		if nums[cur] < num {
			nums[lessIndex+1], nums[cur] = nums[cur], nums[lessIndex+1]
			lessIndex++
			cur++
		} else if nums[cur] > num {
			nums[largerIndex-1], nums[cur] = nums[cur], nums[largerIndex-1]
			largerIndex--
		} else if nums[cur] == num {
			cur++
		}
	}
	//因为我们最后一个是没有排序的，所以大于的索引位置处与最后一个位置的元素交换即可
	nums[largerIndex], nums[R] = nums[R], nums[largerIndex]
	return lessIndex + 1, largerIndex - 1
}
