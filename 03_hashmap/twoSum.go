package __hashmap

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, item := range nums {
		if value, exist := numMap[item]; exist {
			return []int{value, index}
		}
		numMap[target-item] = index
	}
	return []int{}
}

func TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Println(ret)
}
