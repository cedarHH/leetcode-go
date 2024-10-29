package assessment

import (
	"fmt"
	"strings"
)

func TyyAssessment() {
	fmt.Println(getValue(1, 1))
	fmt.Println(countStr("dianxindianxindianx"))
	fmt.Println(minOperation([]int{1, 2, 3}, []int{2, 3, 5}))
}

func getValue(a int, b int) float64 {
	// write code here
	if a == 0 && b != 1 {
		return -1.0
	}
	if a == 0 && b == 1 {
		return 99999.0
	}
	return float64((int((1 - float64(b)) / float64(a) * 10))) / 10.0
}

func countStr(str string) int {
	dx := "dianx"
	dictionary := make(map[uint8]int)
	str = strings.ToLower(str)
	for i := range str {
		dictionary[str[i]] += 1
	}
	if val, exists := dictionary['i']; exists {
		dictionary['i'] = val / 2
	}
	if val, exists := dictionary['n']; exists {
		dictionary['n'] = val / 2
	}
	res := 1000
	for i := range dx {
		if val, exists := dictionary[dx[i]]; exists {
			if val < res {
				res = val
			}
		} else {
			return 0
		}
	}
	return res
}

func minOperation(a []int, b []int) int {
	// write code here
	maxArr := func(arr []int) int {
		res := 0
		for _, v := range arr {
			if v > res {
				res = v
			}
		}
		return res
	}

	maxA := maxArr(a)
	res := 0
	for _, v := range b {
		if v < maxA {
			res += maxA - v
		}
	}
	return res
}
