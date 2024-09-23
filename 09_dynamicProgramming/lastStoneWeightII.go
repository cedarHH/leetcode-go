package __dynamicProgramming

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]
}

func LastStoneWeightII() {
	fmt.Println("LastStoneWeight: ", lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
