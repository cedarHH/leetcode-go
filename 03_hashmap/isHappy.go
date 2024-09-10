package __hashmap

func isHappy(n int) bool {
	numSet := map[int]struct{}{}
	for {
		if _, exist := numSet[n]; exist {
			return false
		}
		numSet[n] = struct{}{}
		if n == 1 {
			return true
		}
		res := 0
		for n > 0 {
			res += (n % 10) * (n % 10)
			n /= 10
		}
		n = res

	}
}

func IsHappy() {
	n := 19
	isHappy(n)
}
