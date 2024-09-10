package __hashmap

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nums12 := map[int]int{}
	result := 0
	for _, item1 := range nums1 {
		for _, item2 := range nums2 {
			nums12[item1+item2]++
		}
	}
	for _, item3 := range nums3 {
		for _, item4 := range nums4 {
			if value, exist := nums12[-item3-item4]; exist {
				result += value
			}
		}
	}
	return result
}

func FourSumCount() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
