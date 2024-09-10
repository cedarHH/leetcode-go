package __strings

import "fmt"

func rightRotation(strByte []byte, n int) string {
	reverse := func(l, r int) {
		for l < r {
			strByte[l], strByte[r] = strByte[r], strByte[l]
			l++
			r--
		}
	}
	l := len(strByte)
	reverse(0, l-n-1)
	reverse(l-n, l-1)
	reverse(0, l-1)
	return string(strByte)
}

func RightRotation() {
	s := "abcdefg"
	fmt.Printf("%s\n", rightRotation([]byte(s), 2))
}
