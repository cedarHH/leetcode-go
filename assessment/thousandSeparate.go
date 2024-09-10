package assessment

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	reversed := reverseString(strconv.Itoa(n))

	re := regexp.MustCompile(`(\d{3})`)

	result := re.ReplaceAllString(reversed, "$1,")

	result = strings.TrimSuffix(result, ",")

	return reverseString(result)
}

func ThousandSeparator() {
	n := 1234567890
	fmt.Printf("%s\n", thousandSeparator(n))
}
