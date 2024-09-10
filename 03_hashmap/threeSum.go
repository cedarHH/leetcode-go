package __hashmap

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var (
		result [][]int
		left   int
		right  int
	)
	length := len(nums)
	sort.Ints(nums)
	for index, value := range nums {
		if index > 0 && value == nums[index-1] {
			continue
		}
		if value > 0 || index == length-2 {
			break
		}
		left = index + 1
		right = length - 1
		for left < right {
			if value+nums[left]+nums[right] < 0 {
				left++
			} else if value+nums[left]+nums[right] > 0 {
				right--
			} else {
				result = append(result, []int{value, nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			}
		}
	}
	return result
}

func threeSum2(nums []int) [][]int {
	var result [][]int
	left, middle, right, length := 0, 1, 2, len(nums)
	sort.Ints(nums)
	for {
		if nums[left]+nums[middle]+nums[right] == 0 {
			result = append(result, []int{nums[left], nums[middle], nums[right]})
		}
		temp := nums[right]
		for {
			right++
			if right == length || nums[right] != temp {
				break
			}
		}
		if right == length {
			temp = nums[middle]
			for {
				middle++
				if middle == length-1 || nums[middle] != temp {
					break
				}
			}
			right = middle + 1
			if middle == length-1 {
				temp = nums[left]
				for {
					left++
					if left == length-2 || nums[left] != temp {
						break
					}
				}
				middle = left + 1
				right = middle + 1
				if left == length-2 {
					break
				}
			}
		}

	}
	return result
}

func threeSum1(nums []int) [][]int {
	sumMap := map[int][][2]int{}
	for i, item1 := range nums {
		for j, item2 := range nums {
			if j > i {
				sumMap[item1+item2] = append(sumMap[item1+item2], [2]int{i, j})
			}
		}
	}
	resSet := map[[3]int]struct{}{}
	for i, item := range nums {
		if value, exist := sumMap[-item]; exist {
			for _, pair := range value {
				if i != pair[0] && i != pair[1] {
					indexPair := [3]int{nums[pair[0]], nums[pair[1]], nums[i]}
					sort.Ints(indexPair[:])
					resSet[indexPair] = struct{}{}
				}
			}
		}
	}
	var result [][]int
	for key := range resSet {
		result = append(result, key[:])
	}
	return result
}

func ThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}
