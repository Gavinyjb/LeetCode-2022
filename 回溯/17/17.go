package main

import "fmt"

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	track := ""
	if len(digits) == 0 {
		return res
	}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var backtrack func(track string, digits string)
	backtrack = func(track string, digits string) {
		//终止条件
		if len(digits) == 0 {
			res = append(res, track)
			return
		}
		key := digits[0:1]
		digits = digits[1:]
		for i := 0; i < len(m[key]); i++ {
			//做选择
			track = track + m[key][i]
			backtrack(track, digits)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, digits)
	return res
}
func main() {
	fmt.Println(letterCombinations("23"))
}
