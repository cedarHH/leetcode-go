package assessment

import "fmt"

func mt1(n int) int {
	if n < 2 {
		return n
	}
	count, res := 2, 2
	for {
		if res >= n {
			return count
		}
		res = res*2 - 1
		count++
	}
}

func mt2(n, m int, a, b, c []int) (res int) {
	labelItems := make(map[int][]int)
	for _, v := range c {
		res += v
	}
	for idx, label := range a {
		labelItems[label] = append(labelItems[label], idx)
	}
	for idx := range labelItems {
		maxBenefit := 0
		for _, item := range labelItems[idx] {
			if b[item]-c[item] > maxBenefit {
				maxBenefit = b[item] - c[item]
			}
		}
		res += maxBenefit
	}
	return
}

func Temp() {
	fmt.Println(mt1(12))

	n, m := 5, 3
	a := []int{1, 1, 2, 2, 3}
	b := []int{10, 20, 30, 40, 50}
	c := []int{5, 5, 10, 10, 15}
	fmt.Println(mt2(n, m, a, b, c))
}
