package __greedy

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	childIndex, cookieIndex := 0, 0

	for childIndex < len(g) && cookieIndex < len(s) {
		if s[cookieIndex] >= g[childIndex] {
			childIndex++
		}
		cookieIndex++
	}

	return childIndex
}

func FindContentChildren() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
