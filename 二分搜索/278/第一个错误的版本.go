package leetcode

func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	//二分搜索
	start, end := 0, n
	for start+1 < end {
		mid := start + (end-start)>>1
		if isBadVersion(mid) {
			end = mid
		} else if isBadVersion(mid) == false {
			start = mid
		}
	}
	if isBadVersion(start) {
		return start
	}
	return end
}
