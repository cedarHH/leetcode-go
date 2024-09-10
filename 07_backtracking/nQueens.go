package __backtracking

import (
	"fmt"
	"math"
	"strings"
)

func solveNQueens(n int) [][]string {
	combinations := make([]string, n)
	combination := make([]int, n)
	var result [][]string
	for i := range combinations {
		combinations[i] = strings.Repeat(".", i) + "Q" + strings.Repeat(".", n-i-1)
	}

	isValid := func(combination []int, rowNum, idx int) bool {
		if math.Abs(float64(combination[rowNum-1]-idx)) < 1.5 {
			return false
		}
		for _, row := range combination[:rowNum] {
			if row == idx {
				return false
			}
		}
		for row, i := rowNum-1, idx+1; row >= 0 && i < n; row, i = row-1, i+1 {
			if combination[row] == i {
				return false
			}
		}
		for row, i := rowNum-1, idx-1; row >= 0 && i >= 0; row, i = row-1, i-1 {
			if combination[row] == i {
				return false
			}
		}

		return true
	}

	var backtrack func(int)
	backtrack = func(rowNum int) {
		if rowNum == n {
			var chessboard []string
			for _, j := range combination {
				chessboard = append(chessboard, combinations[j])
			}
			result = append(result, chessboard)
			return
		}
		for i := 0; i < n; i++ {
			if rowNum == 0 {
				combination[rowNum] = i
				backtrack(rowNum + 1)
			} else if isValid(combination, rowNum, i) {
				combination[rowNum] = i
				backtrack(rowNum + 1)
			}
		}
	}

	backtrack(0)
	return result
}

func SolveNQueens() {
	result := solveNQueens(5)
	for _, row := range result {
		fmt.Println(row)
	}
}
