package __dynamicProgramming

import (
	"fmt"
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	var sum, zeroNum int
	for _, v := range nums {
		sum += v
	}
	if sum < int(math.Abs(float64(target))) {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	positive := (sum + target) / 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, positive+1)
	}
	if nums[0] <= positive {
		dp[0][nums[0]] = 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNum++
		}
		dp[i][0] = 1 << zeroNum
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= positive; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][positive]
}

func FindTargetSumWays() {
	fmt.Println("findTargetSumWays: ", findTargetSumWays([]int{100}, -200))
}
