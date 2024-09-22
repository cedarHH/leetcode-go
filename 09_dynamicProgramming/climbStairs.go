package __dynamicProgramming

import "fmt"

func climbStairs(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2

	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairs() {
	fmt.Println("climb stairs: ", climbStairs(3))
}
