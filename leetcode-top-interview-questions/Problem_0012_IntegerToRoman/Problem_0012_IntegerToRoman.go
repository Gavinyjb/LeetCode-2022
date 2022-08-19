package problem0012integertoroman

import "strings"

func intToRoman(num int) string {
	c := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	var roman  strings.Builder
	roman.WriteString(c[3][num/1000%10])
	roman.WriteString(c[2][num/100%10])
	roman.WriteString(c[1][num/10%10])
	roman.WriteString(c[0][num%10])
	return roman.String()
}