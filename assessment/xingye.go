package assessment

import (
	"fmt"
)

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

func xingyeAss1(a, b []int) (res float32) {
	kthSmallest := func(k int) int {
		index1, index2 := 0, 0
		for {
			if index1 == len(a) {
				return b[index2+k-1]
			}
			if index2 == len(b) {
				return a[index1+k-1]
			}
			if k == 1 {
				return min(a[index1], b[index2])
			}
			pivot1 := min(len(a)-1, index1+k/2-1)
			pivot2 := min(len(b)-1, index2+k/2-1)
			if a[pivot1] < b[pivot2] {
				k -= pivot1 - index1 + 1
				index1 = pivot1 + 1
			} else {
				k -= pivot2 - index2 + 1
				index2 = pivot2 + 1
			}
		}
	}

	num := len(a) + len(b)
	if num%2 == 1 {
		return float32(kthSmallest(num>>1 + 1))
	} else {
		return float32(kthSmallest(num>>1)+kthSmallest(num>>1+1)) / 2
	}
}
