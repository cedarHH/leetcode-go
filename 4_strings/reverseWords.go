package __strings

import "fmt"

func reverseWords(s string) string {
	b := []byte(s)
	slow, fast := 0, 0
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}

	for fast < len(b) {
		if b[fast] != ' ' {
			b[slow] = b[fast]
			slow++
			fast++
		} else {
			for fast < len(b) && b[fast] == ' ' {
				fast++
			}
			if fast < len(b) {
				b[slow] = ' '
				slow++
			}
		}
	}
	b = b[:slow]
	slow, fast = 0, len(b)-1

	reverse := func(start, end int) {
		for start < end {
			b[start], b[end] = b[end], b[start]
			start++
			end--
		}
	}

	reverse(slow, fast)

	slow, fast = 0, 0
	for fast < len(b) {
		for fast < len(b) && b[fast] != ' ' {
			fast++
		}
		reverse(slow, fast-1)
		fast++
		slow = fast
	}

	return string(b)
}

func ReverseWords() {
	s := " a good   example "
	fmt.Printf("%s\n", reverseWords(s))
}
