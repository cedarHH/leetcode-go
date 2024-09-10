package assessment

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	nums := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		nums[i] = 10001
		for _, j := range coins {
			if i >= j {
				nums[i] = min(nums[i], nums[i-j]+1)
			}
		}
	}
	if nums[amount] == 10001 {
		nums[amount] = -1
	}
	return nums[amount]
}

func CoinChange() {
	coins := []int{2}
	fmt.Printf("%d\n", coinChange(coins, 3))
}
