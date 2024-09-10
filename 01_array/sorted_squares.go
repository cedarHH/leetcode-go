package __array

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	resultIndex := len(nums) - 1
	result := make([]int, len(nums))

	for resultIndex != -1 {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			result[resultIndex] = nums[right] * nums[right]
			resultIndex--
			right--
		} else {
			result[resultIndex] = nums[left] * nums[left]
			resultIndex--
			left++
		}
	}
	return result
}

func SortedSquares() {
	nums := []int{-5, -3, -2, -1}
	fmt.Println(sortedSquares(nums))
}
