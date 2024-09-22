package __dynamicProgramming

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(cost)]
}

func MinCostClimbingStairs() {
	fmt.Println("minCostClimbingStairs: ", minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
