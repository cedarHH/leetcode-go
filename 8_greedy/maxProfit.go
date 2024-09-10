package __greedy

import "fmt"

func maxProfit(prices []int) int {
	prev, result := prices[0], 0
	for _, price := range prices[1:] {
		if price > prev {
			result += price - prev
			prev = price
		} else {
			prev = price
		}
	}
	return result
}

func MaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
