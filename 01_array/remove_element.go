package __array

import "fmt"

func removeElement(nums []int, val int) int {
	var (
		index  = 0
		length = len(nums)
	)
	if length == 0 {
		return 0
	}
	for {
		if index == length {
			return length
		}
		if nums[index] == val {
			nums[index] = nums[length-1]
			length--
		} else {
			index++
		}
	}
}

func RemoveElement() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
