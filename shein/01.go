package main

import "strings"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param haystack string字符串
 * @param needle string字符串
 * @return int整型
 */
func strStr(haystack string, needle string) int {
	// write code here
	return strings.Index(haystack, needle)
}
