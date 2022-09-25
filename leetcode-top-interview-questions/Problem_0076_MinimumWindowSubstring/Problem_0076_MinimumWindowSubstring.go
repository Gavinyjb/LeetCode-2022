package Problem_0076_MinimumWindowSubstring

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	//[l,r) 左闭右开
	s2b := []byte(s)
	t2b := []byte(t)

	restMap := restList(t2b)
	l, r := 0, 0
	all := len(t2b)

	minLen, ansL, ansR := len(s)+1, -1, -1
	for r < len(s2b) {
		restMap[s2b[r]]--
		if restMap[s2b[r]] >= 0 {
			all--
		}
		if all == 0 {
			for restMap[s2b[l]] < 0 {
				restMap[s2b[l]]++
				l++
			}
			if minLen > r-l+1 {
				ansL = l
				ansR = r
				minLen = r - l + 1
			}
			all++
			restMap[s2b[l]]++
			l++
		}
		r++
	}
	if minLen == len(s2b)+1 {
		return ""
	} else {
		return s[ansL : ansR+1]
	}
}
func restList(t []byte) map[byte]int {
	ret := make(map[byte]int)
	for _, b := range t {
		ret[b]++
	}
	return ret
}
