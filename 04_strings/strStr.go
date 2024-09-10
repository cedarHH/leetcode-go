package __strings

import "fmt"

func strStr(haystack string, needle string) int {

	next := make([]int, len(needle))

	for i := 1; i < len(needle); i++ {
		index := i - 1
		for {
			if needle[i] == needle[next[index]] {
				next[i] = next[index] + 1
				break
			} else if index == 0 {
				break
			} else {
				index = next[index-1]
			}
		}
	}
	j := 0
	n := len(needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func StrStr() {
	haystack := "aabaaabaaac"
	needle := "aabaaac"
	fmt.Printf("%d\n", strStr2(haystack, needle))
}

func strStr2(haystack string, needle string) int {
	next := make([]int, len(needle))
	var index int
	index = 0
	for i := 1; i < len(needle); i++ {
		for index > 0 && needle[i] != needle[index] {
			index = next[index-1]
		}
		if needle[i] == needle[index] {
			index++
		}
		next[i] = index
	}
	index = 0
	for i := 0; i < len(haystack); i++ {
		for index > 0 && haystack[i] != needle[index] {
			index = next[index-1]
		}
		if haystack[i] == needle[index] {
			index++
		}
		if index == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
