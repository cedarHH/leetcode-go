package assessment

import (
	"fmt"
	"sort"
)

func dewuAss1(arr []int, arrLen int) bool {
	sort.Ints(arr)
	if arr[arrLen-1]-arr[0]-arr[1] < 0 {
		return false
	} else {
		return true
	}
}

func dewuAss2(arr []int, arrLen int) bool {
	fmt.Println(arr, arrLen)
	if arrLen < 3 {
		return false
	}
	for idx := 2; idx < arrLen-1; idx++ {
		if arr[idx] <= arr[idx-2] {
			return true
		}
	}

	return false
}

func dewuAss3(num int) (res int) {
	if num < 9 {
		return
	}
	nums := make(map[int]bool)
	for i := 1; i < num; i++ {
		nums[i*i] = true
	}

	isExists := func(a int) bool {
		_, exists := nums[a]
		return exists
	}

	doubleTuple := make(map[[2]int]bool)
	for i := 1; i <= num; i++ {
		for j := i + 1; j <= num; j++ {
			if isExists(i * j) {
				doubleTuple[[2]int{i, j}] = true
			}
		}
	}

	isExistsM := func(a [2]int) bool {
		_, exists := doubleTuple[a]
		return exists
	}

	for key, _ := range doubleTuple {
		for i := key[1] + 1; i <= num; i++ {
			if isExistsM([2]int{key[0], i}) && isExistsM([2]int{key[1], i}) {
				res++
			}
		}
	}

	return res
}

func ass12In() {
	nGroup := 0
	arrLen := 0
	fmt.Scan(&nGroup)
	var ans = make([]bool, nGroup)
	for i := 0; i < nGroup; i++ {
		fmt.Scan(&arrLen)
		var arr = make([]int, arrLen)
		for j := 0; j < arrLen; j++ {
			x := 0
			fmt.Scan(&x)
			arr[j] = x
		}
		// ans[i] = dewuAss1(arr, arrLen)
		ans[i] = dewuAss2(arr, arrLen)
	}
	for _, val := range ans {
		if val {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
}

func ass3In() {
	a := 0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", dewuAss3(a))
		}
	}
}

func DewuAssessment() {
	//ass12In()
	ass3In()
}
