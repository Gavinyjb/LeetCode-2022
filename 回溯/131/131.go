package main

func partition(s string) [][]string {
	ret := make([][]string, 0)
	track := make([]string, 0)
	var backtrack func(str string, start int)
	backtrack = func(str string, start int) {
		if start >= len(str) {
			ret = append(ret, track)
		}
		for i := start; i < len(str); i++ {
			if isPalindrome(str, start, i) { // 是回文子串
				// 获取[startIndex,i]在s中的子串
				temp := str[start : i+1]
				ret = append(ret, temp)
			}
		}
	}
	backtrack(s, 0)
}

func isPalindrome(str string, start, end int) bool {
	for i, j := start, end; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j++
	}
	return true
}

func main() {

}
