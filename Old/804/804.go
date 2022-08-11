package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	str:=[]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	m:=make(map[string]string,0)
	for i, s := range str {
		m[string(i+'a')]=s
	}
	//fmt.Println(m)
	result:=make(map[string]bool,0)
	for _, word := range words {
		track:=""
		for len(word)>0{
			key:=word[:1]
			word=word[1:]
			track=track+m[key]
		}
		//fmt.Println(track)
		result[track]=true
	}
	return len(result)
}
func main() {
	words:=[]string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
