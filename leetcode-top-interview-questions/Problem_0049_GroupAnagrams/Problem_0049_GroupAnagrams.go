package Problem_0049_GroupAnagrams

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	myMap := make(map[string][]string)

	for _, str := range strs {
		str2Bytes := []byte(str)
		//将原本的字符串 按照字符排序
		sort.Slice(str2Bytes, func(i, j int) bool {
			return str2Bytes[i] < str2Bytes[j]
		})
		key := string(str2Bytes)
		if _, ok := myMap[key]; !ok {
			myMap[key] = make([]string, 0)
		}
		myMap[key] = append(myMap[key], str)
	}
	res := make([][]string, 0)
	for _, v := range myMap {
		res = append(res, v)
	}
	return res
}

//测试[]byte 排序
func demo(strs []string) {
	for _, str := range strs {
		str2Bytes := []byte(str)
		//将原本的字符串 按照字符排序
		sort.Slice(str2Bytes, func(i, j int) bool {
			return str2Bytes[i] < str2Bytes[j]
		})
		key := string(str2Bytes)
		fmt.Println(key)
	}
}
