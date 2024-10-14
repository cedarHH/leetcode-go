package __dynamicProgramming

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	var dfs func(int, int) bool
	memo := map[int]bool{}
	dfs = func(currState, total int) bool {
		if _, exists := memo[currState]; exists == false {
			res := false
			for i := 0; i < maxChoosableInteger; i++ {
				if currState&(1<<i) == 0 && (total+i+1 >= desiredTotal || !dfs(currState|(1<<i), total+i+1)) {
					res = true
					break
				}
			}
			memo[currState] = res
		}
		return memo[currState]
	}
	return dfs(0, 0)
}

func CanIWin() {
	fmt.Println("can I win: ", canIWin(10, 11))
}
