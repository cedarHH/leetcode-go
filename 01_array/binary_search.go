package __array

import "fmt"

func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		index int
	)
	for {
		index = left + (right-left)/2 //Overflow prevention
		switch {
		case left > right:
			return -1
		case nums[index] == target:
			return index
		case nums[index] < target:
			left = index + 1
		case nums[index] > target:
			right = index - 1
		}
	}
}

func Search() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 1
	fmt.Println(search(nums, target))
}
