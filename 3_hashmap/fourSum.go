package __hashmap

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		value1 := nums[i]
		if i > 0 && value1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			value2 := nums[j]
			if j > i+1 && value2 == nums[j-1] {
				continue
			}
			left, right := j+1, length-1
			for left < right {
				value3, value4 := nums[left], nums[right]
				sum := value1 + value2 + value3 + value4
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{value1, value2, value3, value4})
					for left < right && value3 == nums[left+1] {
						left++
					}
					for left < right && value4 == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return result
}

func FourSum() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fourSum(nums, target)
}
