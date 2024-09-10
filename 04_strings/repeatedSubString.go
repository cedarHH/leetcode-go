package __strings

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	next := make([]int, length)
	var index int
	index = 0
	for i := 1; i < length; i++ {
		for index > 0 && s[i] != s[index] {
			index = next[index-1]
		}
		if s[i] == s[index] {
			index++
		}
		next[i] = index
	}
	return next[length-1] != 0 && length%(length-next[length-1]) == 0
}

func RepeatedSubstringPattern() {
	fmt.Println(repeatedSubstringPattern("babbaaabbbbabbaaabbb"))
}
