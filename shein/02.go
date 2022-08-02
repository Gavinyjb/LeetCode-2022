package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return int整型
 */
func numDecoding(s string) int {
	// write code here

	n := len(s)
	list := make([]int, n+1)
	list[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			list[i] += list[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			list[i] += list[i-2]
		}
	}
	return list[n]
}
