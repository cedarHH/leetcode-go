package __hashmap

func intersection(nums1 []int, nums2 []int) []int {
	numMap := map[int]struct{}{}
	for _, key := range nums1 {
		numMap[key] = struct{}{}
	}
	var result []int
	for _, key := range nums2 {
		if _, exist := numMap[key]; exist {
			result = append(result, key)
			delete(numMap, key)
		}
	}
	return result
}

func Intersection() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{5, 2, 3, 6}
	intersection(nums1, nums2)
}
