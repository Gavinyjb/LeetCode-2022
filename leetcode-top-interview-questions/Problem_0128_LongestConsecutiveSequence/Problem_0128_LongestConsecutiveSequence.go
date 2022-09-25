package Problem_0128_LongestConsecutiveSequence

func longestConsecutive(nums []int) int {
	//map 查找 动规思想
	//数字要连续 因此重复不计算其中
	mp := make(map[int]int)
	res := 0

	for _, value := range nums {
		if mp[value] == 0 {
			//sum是现阶段数组最长序列
			//left 是 左连续序列 长度 right是右连续序列长度
			left, right, sum := 0, 0, 0

			if mp[value-1] > 0 {
				//有前缀值
				left = mp[value-1]
			} else {
				left = 0
			}

			if mp[value+1] > 0 {
				//有后缀值
				right = mp[value+1]
			} else {
				right = 0
			}

			//左数字个数 加 有数字加本身
			sum = left + right + 1

			//更新mp中此時 value 的值
			mp[value] = sum

			res = max(res, sum)

			//value-left 是这个数组的初始起点
			mp[value-left] = sum
			mp[value+right] = sum
		} else {
			continue
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
