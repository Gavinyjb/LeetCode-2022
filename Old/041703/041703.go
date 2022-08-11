package main

import (
	"fmt"
	"sort"
)

//二分搜索

func getMin(indexs []int) int {
	//遍历 最小间隔一定在相邻位置之间产生
	min := indexs[len(indexs)-1] - indexs[0]
	//长度至少大于1
	if len(indexs) <= 1 {
		return 0
	}
	for i := 1; i < len(indexs); i++ { // 遍历n-1 次
		if indexs[i]-indexs[i-1] < min {
			min = indexs[i] - indexs[i-1]
		}
	}
	return min
}

func getMax(indexs []int) int {
	//已经排序  最大-最小   最后一位-第一位
	return indexs[len(indexs)-1] - indexs[0]
}

//验证x是否可行
func help(indexs []int, ret *[]int, m *int) {
	if *m == 1 {
		*ret = append(*ret, indexs[0])
		return
	}
	if *m >= 2 {
		if indexs[len(indexs)-1]-indexs[len(indexs)-2] > indexs[1]-indexs[0] {
			//取尾端
			*ret = append(*ret, indexs[len(indexs)-1])
			indexs = indexs[:len(indexs)-1]
		} else {
			//取首端
			*ret = append(*ret, indexs[0])
			indexs = indexs[1:]
		}
		*m--
		help(indexs, ret, m)
	}
}
func main() {
	//n, m := 5, 3
	m := 3
	indexs := []int{1, 9, 8, 4, 2}

	sort.Ints(indexs) //排序数组，便于得到间距的最大值与最小值
	ret := make([]int, 0)
	help(indexs, &ret, &m)
	fmt.Println(ret)
}
