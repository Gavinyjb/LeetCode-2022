package Problem_0056_MergeIntervals

import (
	"sort"
)

type Range struct {
	start int
	end   int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	arr := make([]Range, len(intervals))
	//初始化[]Range
	for i, interval := range intervals {
		arr[i] = Range{
			interval[0],
			interval[1],
		}
	}
	//按照开始时间排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].start < arr[j].start
	})

	ans := make([]Range, 0)

	s := arr[0].start
	e := arr[0].end
	for i := 1; i < len(arr); i++ {
		if arr[i].start > e {
			ans = append(ans, Range{s, e})
			//更新 s e
			s = arr[i].start
			e = arr[i].end
		} else {
			e = max(e, arr[i].end)
		}
	}
	ans = append(ans, Range{s, e})

	return generateMatrix(ans)
}
func generateMatrix(input []Range) [][]int {
	ret := make([][]int, len(input))
	for i := 0; i < len(ret); i++ {
		ret[i] = []int{input[i].start, input[i].end}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
