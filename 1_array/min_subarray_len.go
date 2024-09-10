package __array

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	left, right, sum, res := 0, 0, nums[0], length+1

	for right < length-1 {
		if sum < target {
			right++
			sum += nums[right]
		} else {
			res = min(res, right-left+1)
			if sum-nums[left] >= target {
				sum -= nums[left]
				left++
			} else {
				right++
				sum += nums[right]
			}
		}
	}
	if sum < target {
		return 0
	}
	for sum-nums[left] >= target {
		sum -= nums[left]
		left++
	}
	res = min(res, right-left+1)
	return res
}

func MinSubArrayLen() {
	target := 15
	nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	fmt.Println(minSubArrayLen(target, nums))
}
