package __dynamicProgramming

import "fmt"

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Fib() {
	fmt.Println("fib: ", fib(3))
}
