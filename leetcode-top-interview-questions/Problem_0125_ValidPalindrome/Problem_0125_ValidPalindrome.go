package Problem_0125_ValidPalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for !check(s[i]) && i < j {
			i++
		}
		for !check(s[j]) && i < j {
			j--
		}
		if i < j && s[i] != s[j] {
			return false
		}
	}
	return true
}
func check(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}
