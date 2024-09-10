package __array

import "fmt"

func Pow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func rangeSum(nums []int, n int, left int, right int) int {
	modulo := 1e9 + 7
	length := len(nums)
	prefixSum := make([]int, length+1)
	prefixPrefixSums := make([]int, length+1)

	for idx, val := range nums {
		prefixSum[idx+1] = prefixSum[idx] + val
	}
	for idx, val := range prefixSum[1:] {
		prefixPrefixSums[idx+1] = prefixPrefixSums[idx] + val
	}

	getCount := func(val int) int {
		count, j := 0, 1
		for i := range nums {
			for j <= length && prefixSum[j]-prefixSum[i] <= val {
				j++
			}
			j--
			count += j - i
		}
		return count
	}

	getKth := func(k int) int {
		low, high := 0, prefixSum[n]
		for low < high {
			mid := low + (high-low)>>1
			if getCount(mid) < k {
				low = mid + 1
			} else {
				high = mid
			}
		}
		return low
	}

	getSum := func(k int) int {
		total, count, j := 0, 0, 1
		num := getKth(k)
		for i := range nums {
			for j <= length && prefixSum[j]-prefixSum[i] < num {
				j++
			}
			j--
			total += prefixPrefixSums[j] - prefixPrefixSums[i] - (j-i)*prefixSum[i]
			count += j - i
		}
		total += num * (k - count)
		return total
	}

	return (getSum(right) - getSum(left-1)) % int(modulo)
}

func RangeSum() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(rangeSum(nums, 4, 1, 5))
}
