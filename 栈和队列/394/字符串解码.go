package leetcdoe

import "strings"

var (
	src string
	ptr int
)

func decodeString(str string) string {
	src = str
	ptr = 0
	return getString()
}
func getString() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur <= '9' && cur >= '0' {
		repTime = getDigites()
		ptr++
		str := getString()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString()
}

func getDigites() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}
