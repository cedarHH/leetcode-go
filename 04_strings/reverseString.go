package __strings

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
)

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func ReverseString() {
	strings := []string{"h", "e", "l", "l", "o"}
	s := stream.Map(stream.Of(strings...), func(it string) byte {
		return it[0]
	}).ToSlice()

	reverseString(s)
	fmt.Println(s)
}
