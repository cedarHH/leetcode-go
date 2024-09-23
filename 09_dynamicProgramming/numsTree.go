package __dynamicProgramming

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] * 2 * (2*i + 1) / (i + 2)
	}
	return dp[n]
}

func NumTrees() {
	fmt.Println("NumTrees: ", numTrees(4))
}
