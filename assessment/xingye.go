package assessment

import "fmt"

func XingyeAssessment() {
	fmt.Println(maxLength([]int{1, 2, 3, 4, 5}, 3))
}

func maxLength(nums []int, k int) int {
	// write code here
	maxNums := func(m int) (sum int) {
		for i := range nums {
			sum += nums[i] / m
		}
		return
	}

	maxPipe := 0
	for i := range nums {
		if nums[i] > maxPipe {
			maxPipe = nums[i]
		}
	}

	left, right := 1, maxPipe
	mid := left + (right-left)/2
	midNums := maxNums(mid)
	if maxNums(left) < k {
		return -1
	}
	if maxNums(right) >= k {
		return maxPipe
	}
	for {
		if midNums >= k && maxNums(mid+1) < k {
			return mid
		}
		if midNums < k {
			right = mid
			mid = left + (right-left)/2
			midNums = maxNums(mid)
		} else {
			left = mid
			mid = left + (right-left)/2
			midNums = maxNums(mid)
		}
	}
}
