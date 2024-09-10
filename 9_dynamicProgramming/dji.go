package __dynamicProgramming

import (
	"fmt"
	"math/rand"
)

func dji(arr [][]int) int {
	rows, cols := len(arr), len(arr[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	dp[rows-1][cols-1] = 1
	for i := rows - 2; i >= 0; i-- {
		dp[i][cols-1] = max(arr[i+1][cols-1]+dp[i+1][cols-1], 1)
	}
	for j := cols - 2; j >= 0; j-- {
		dp[rows-1][j] = max(arr[rows-1][j+1]+dp[rows-1][j+1], 1)
	}
	for i := rows - 2; i >= 0; i-- {
		for j := cols - 2; j >= 0; j-- {
			dp[i][j] = max(min(arr[i+1][j]+dp[i+1][j],
				arr[i][j+1]+dp[i][j+1],
				arr[i+1][j+1]+dp[i+1][j+1]),
				1)
		}
	}
	//for _, row := range arr {
	//	fmt.Printf("%v\n", row)
	//}
	//for _, row := range dp {
	//	fmt.Printf("%v\n", row)
	//}
	return dp[0][0]
}

func Dji() {
	arr := make([][]int, 5)
	for i := range arr {
		arr[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			arr[i][j] = rand.Intn(100) - 10
		}
	}
	fmt.Printf("%d\n", dji(arr))
}
