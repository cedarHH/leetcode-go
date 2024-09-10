package __strings

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	reverse := func(s []rune) {
		for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
			s[left], s[right] = s[right], s[left]
		}
	}

	runes := []rune(s)
	length := len(runes)
	i := 0
	for ; i+2*k < length; i += 2 * k {
		reverse(runes[i : i+k])
	}
	if i+k < length {
		reverse(runes[i : i+k])
	} else {
		reverse(runes[i:length])
	}
	return string(runes)
}

func ReverseStr() {
	s := "abcd"
	reverseS := reverseStr(s, 4)
	fmt.Println(reverseS)
}
