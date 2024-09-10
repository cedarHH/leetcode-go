package __hashmap

func canConstruct(ransomNote string, magazine string) bool {
	charSet := map[rune]int{}
	for _, char := range magazine {
		charSet[char]++
	}
	for _, char := range ransomNote {
		charSet[char]--
		if charSet[char] == -1 {
			return false
		}
	}
	return true
}

func CanConstruct() {
	ransomNote := "aa"
	magazine := "aab"
	canConstruct(ransomNote, magazine)
}
