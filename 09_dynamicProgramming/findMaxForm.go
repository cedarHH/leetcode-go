package __dynamicProgramming

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	count10 := func(word string) (c0 int, c1 int) {
		c0, c1 = 0, 0
		for _, c := range word {
			if c == '0' {
				c0++
			} else {
				c1++
			}
		}
		return
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroNum, oneNum := count10(str)
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func FindMaxForm() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	fmt.Println("Find max form: ", findMaxForm(strs, 5, 3))
}
