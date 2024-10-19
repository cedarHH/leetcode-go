// Package assessment provides ...
package assessment

import (
	"fmt"
)

func JdAssessment() {
	n := 0

	fmt.Scan(&n)
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		x := ""
		fmt.Scan(&x)
		strs[i] = x
	}
	lcp := strs[0]
	fmt.Println(lcp)
	for i := 1; i < n; i++ {
		if len(lcp) == 0 {
			fmt.Println(-1)
			continue
		}
		j := 0
		for ; j < len(lcp) && j < len(strs[i]) && lcp[j] == strs[i][j]; j++ {
		}
		lcp = lcp[:j]
		if len(lcp) == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(lcp)
		}
	}
}
