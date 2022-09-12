package Problem0028implementstrstr

//func strStr(haystack string, needle string) int {
//	if haystack == "" || needle == "" || len(needle) < 1 || len(haystack) < len(needle) {
//		return -1
//	}
//	s1 := []byte(haystack)
//	s2 := []byte(needle)
//	next := getNextArray(s2)
//	i, j := 0, 0
//	for i < len(s1) && j < len(s2) {
//		if s1[i] == s2[j] {
//			i++
//			j++
//		} else if next[j] == -1 {
//			i++
//		} else {
//			j = next[j]
//		}
//	}
//	if j == len(s2) {
//		return i - j
//	} else {
//		return -1
//	}
//}
//
//func getNextArray(ms []byte) []int {
//	if len(ms) == 1 {
//		return []int{-1}
//	}
//	next := make([]int, len(ms))
//	next[0] = -1 //任何一个字符串，0的位置上都是-1
//	next[1] = 0
//	i := 2 //目前在哪个位置上求next数组的值
//	//cn: cn位置的字符，是当前和i-1位置比较的字符
//	//永远是cn与i-1比较
//	cn := 0
//	for i < len(next) {
//		if ms[i-1] == ms[cn] { // i-1与cn 匹配
//			next[i] = cn + 1
//			i++
//			cn++
//		} else if cn > 0 { //暂时不匹配，但是cn还能往左跳
//			cn = next[cn]
//		} else {
//			next[i] = 0
//			i++
//		}
//	}
//	return next
//}

func strStr(haystack, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 || len(haystack) < len(needle) {
		return -1
	}
	s1 := []byte(haystack)
	s2 := []byte(needle)
	next := getNexts(s2)
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
			j++
		} else if next[j] == -1 {
			j = 0
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(s2) {
		return i - j
	} else {
		return -1
	}

	return -1
}
func getNexts(ms []byte) []int {
	if len(ms) == 1 {
		return []int{-1}
	}
	if len(ms) == 2 {
		return []int{-1, 0}
	}
	next := make([]int, len(ms))
	next[0] = -1
	next[1] = 0
	i := 2 //
	cn := 0
	for i < len(ms) {
		if ms[i-1] == ms[cn] {
			next[i] = cn + 1
			i++
			cn++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
