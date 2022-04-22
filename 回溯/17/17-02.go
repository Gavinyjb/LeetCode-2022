package main

import "fmt"

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	track := ""
	if len(digits) == 0 {
		return ret
	}
	letterMap := map[string][]string{
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
			temp := ""
			temp = track
			ret = append(ret, temp)
			return
		}

		index := digits[0:1]
		digits = digits[1:]
		for i := 0; i < len(letterMap[index]); i++ {
			track = track + letterMap[index][i]
			backtrack(track, digits)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, digits)
	return ret
}
func main() {
	str := "1"
	fmt.Println(str[0:1])
}
