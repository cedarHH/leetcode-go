package assessment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func dewu2Ass1() {
	t, n := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	if scanner.Scan() {
		t, _ = strconv.Atoi(scanner.Text())
	}
	var choice, correct string
	for i := 0; i < t; i++ {
		if scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
		}
		if scanner.Scan() {
			choice = scanner.Text()
		}
		if scanner.Scan() {
			correct = scanner.Text()
		}
		fmt.Println(dewu2Ass1Handle(n, choice, correct))
	}
}

func dewu2Ass1Handle(n int, choice, correct string) string {
	countCorrect := 0
	for i := range correct {
		if correct[i] == choice[i] {
			countCorrect += 1
		}
	}
	if 2*countCorrect > n {
		return "oh yes"
	} else if 2*countCorrect == n {
		return "(O.O)"
	} else {
		return "oh no"
	}
}

func dewu2Ass2() {
	n, m := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	if scanner.Scan() {
		m, _ = strconv.Atoi(scanner.Text())
	}
	menu := make([][2]int, n+1)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			menu[i][0], _ = strconv.Atoi(scanner.Text())
		}
		if scanner.Scan() {
			menu[i][1], _ = strconv.Atoi(scanner.Text())
		}
	}
	//n := 3
	//m := 6
	//menu := [][2]int{{1, 3}, {3, 6}, {1, 3}}
	fmt.Println(dewu2Ass2Handle(n, m, menu))
}

func dewu2Ass2Handle(n, m int, menu [][2]int) int {
	dp := make([]int, m+1)
	for i := range menu {
		for j := m; j >= menu[i][1]; j-- {
			profit := dp[j-menu[i][1]] + menu[i][1] - menu[i][0]
			if profit > dp[j] {
				dp[j] = profit
			}
		}
	}
	return dp[m]
}

func dewu2Ass3() {
	n, k := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	if scanner.Scan() {
		k, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(dewu2Ass3Handle(n, k))
}

func dewu2Ass3Handle(n, k int) (res int) {
	var dfs func([]int, int, int)
	dfs = func(memo []int, length, first int) {
		count := 0
		flag := true
		for i := n; i > 0; i-- {
			if memo[i] == 1 {
				count++
			} else {
				if length == 0 {
					first = i
				} else if length == 1 {
					if i < first {
						return
					}
				}
				memo[i] = 1
				flag = false
				memo[0] += count
				if memo[0] <= k {
					dfs(memo, length+1, first)
					memo[i] = 0
					memo[0] -= count
				} else {
					memo[i] = 0
					memo[0] -= count
					return
				}
			}
		}
		if flag {
			res++
		}
	}
	memo := make([]int, n+1)
	dfs(memo, 0, 0)
	return res % (10e9 + 7)
}

func DewuAssessment2() {
	//dewu2Ass1()
	//dewu2Ass2()
	dewu2Ass3()
}
