package __hashmap

func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, r := range s {
		record[r-'a']++
	}
	for _, r := range t {
		record[r-'a']--
	}
	return record == [26]int{}
}

func IsAnagram() {
	s := "qwe"
	t := "asd"
	isAnagram(s, t)
}
