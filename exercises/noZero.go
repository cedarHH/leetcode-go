package exercises

import "fmt"

func getNoZeroIntegers(n int) []int {
	pow10 := func(n int) int {
		base := 10
		result := 1
		for n > 0 {
			if n&1 == 1 {
				result *= base
			}
			base *= base
			n >>= 1
		}
		return result
	}

	result := []int{0, 0}
	i := 0
	for n > 0 {
		remainder := n % 10
		n = n / 10
		if remainder <= 1 && n != 0 {
			n -= 1
			result[0] = result[0] + 2*pow10(i)
			result[1] = result[1] + (remainder+8)*pow10(i)
		} else {
			result[0] = result[0] + pow10(i)
			result[1] = result[1] + (remainder-1)*pow10(i)
		}
		i++
	}
	return result
}

func GetNoZeroIntegers() {
	n := 10000
	fmt.Printf("%v\n", getNoZeroIntegers(n))
}
