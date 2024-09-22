package __dynamicProgramming

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := range dp {
		for j := range dp[0] {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func UniquePathsWithObstacles() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println("uniquePathsWithObstacles: ", uniquePathsWithObstacles(obstacleGrid))
}
