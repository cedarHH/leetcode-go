package assessment

import (
	"fmt"
	"math"
	"sort"
)

func sfAss1() {
	min2 := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	area := 0
	_, _ = fmt.Scan(&area)
	dp := make([]int, area+1)
	for i := 1; i <= area; i++ {
		a := int(math.Sqrt(float64(i)))
		if a*a == i {
			dp[i] = a
			continue
		}
		dp[i] = min2(dp[i-1]+1, dp[a*a]+dp[i-a*a])
		if a > 3 {
			n := i / ((a - 1) * (a - 1))
			r := i % ((a - 1) * (a - 1))
			dp[i] = min2(dp[i], n*dp[(a-1)*(a-1)]+dp[r])
		}
	}
	fmt.Println(4 * dp[area])
}

//n, m := 5, 100
//team := []int{82, 37, 43, 121, 55}

func sfAss2() {
	n, m := 0, 0
	_, _ = fmt.Scan(&n, &m)
	team := make([]int, n)
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		team[i] = x
	}

	sort.Ints(team)
	left, right, res, num := 0, n-1, 0, 1
	for left+num-1 <= right {
		if team[right]*num >= m {
			res += 1
			right--
			left += num - 1
		} else {
			num++
		}
	}
	remain := 0
	for left <= right {
		remain += team[left]
		left++
	}
	if remain >= m {
		res++
	}
	fmt.Println(res)
}

func SfAssessment() {
	// sfAss1()
	sfAss2()
}
