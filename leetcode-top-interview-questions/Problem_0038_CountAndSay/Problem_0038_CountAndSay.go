package Problem_0038_CountAndSay

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	last := []byte(countAndSay(n - 1))
	ans := new(strings.Builder)
	times := 1
	for i := 1; i < len(last); i++ {
		if last[i] == last[i-1] {
			times++
		} else {
			ans.WriteString(strconv.Itoa(times))
			ans.WriteByte(last[i-1])
			times = 1
		}
	}
	ans.WriteString(strconv.Itoa(times))
	ans.WriteByte(last[len(last)-1])
	return ans.String()
}
