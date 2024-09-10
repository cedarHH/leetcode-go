package __monostack

import "fmt"

func trap(height []int) int {
	var monoStack []int
	sum := 0
	for i, h := range height {
		if len(monoStack) == 0 || h < height[monoStack[len(monoStack)-1]] {
			monoStack = append(monoStack, i)
		} else {
			for len(monoStack) > 1 {
				sum += (min(h, height[monoStack[len(monoStack)-2]]) -
					height[monoStack[len(monoStack)-1]]) *
					(i - monoStack[len(monoStack)-2] - 1)
				monoStack = monoStack[:len(monoStack)-1]
				if height[monoStack[len(monoStack)-1]] >= h {
					break
				}
			}
			if len(monoStack) == 1 && height[monoStack[0]] < h {
				monoStack[0] = i
			} else if height[monoStack[len(monoStack)-1]] == h {
				monoStack[len(monoStack)-1] = i
			} else {
				monoStack = append(monoStack, i)
			}
		}
	}
	return sum
}

func Trap() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
