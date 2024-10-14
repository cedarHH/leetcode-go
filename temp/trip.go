package temp

import "fmt"

func tripAss1(variable string) string {
	var res []rune
	var tag int
	for _, chr := range variable {
		if chr == '_' {
			tag = 1
		} else if tag == 1 {
			tag = 0
			res = append(res, chr-32)
		} else {
			res = append(res, chr)
		}
	}
	res[0] = res[0] - 32
	return string(res)
}

func ass1Handle() {
	n := 0
	var ans []string

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		variable := ""
		fmt.Scan(&variable)
		ans = append(ans, tripAss1(variable))
	}
	for _, str := range ans {
		fmt.Println(str)
	}
}

func build() [1000001]int {
	n := 1000000
	var primes [1000001]int
	primes[0], primes[1] = 1, 1
	for idx := 2; idx <= n; idx++ {
		for curr := idx * idx; curr <= n; curr += idx {
			primes[curr] = 1
		}
	}
	return primes
}

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func tripAss2(n int, primes *[1000001]int, dayMap *map[int]int) (days int) {
	arr := make([]int, 0)
	for n > 0 {
		if val, exists := (*dayMap)[n]; exists {
			return days + val
		}
		arr = append(arr, n)
		if n < 1000001 {
			if primes[n] == 0 {
				n -= n/3 + 1
			} else {
				n -= n/2 + 1
			}
		} else {
			if isPrime(n) {
				n -= n/3 + 1
			} else {
				n -= n/2 + 1
			}
		}
		days++
	}
	for idx, val := range arr {
		(*dayMap)[val] = days - idx
	}
	return days
}

func ass2Handle() {
	a := 0
	primes := build()
	dayMap := make(map[int]int)
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", tripAss2(a, &primes, &dayMap))
		}
	}
}

func gcd(arr []int) int {
	minIdx := func(arr []int) (idx int) {
		for i, v := range arr {
			if v < arr[idx] {
				idx = i
			}
		}
		return
	}
	nums := 1
	for nums <= len(arr) {
		mIdx := minIdx(arr)
		for i := range arr {
			if arr[i] != arr[mIdx] {
				m := arr[i] % arr[mIdx]
				if m != 0 {
					arr[i] = m
				} else {
					arr[i] = arr[mIdx]
					nums++
				}
			}
		}
	}
	return arr[0]
}

func tripAss3(n, m int, arr []int) int {
	return 0
}

func ass3Handle() {
	//n, m := 0, 0
	//_, _ = fmt.Scan(&n, &m)
	//arr := make([]int, n)
	//for j := 0; j < n; j++ {
	//	x := 0
	//	fmt.Scan(&x)
	//	arr[j] = x
	//}
	//fmt.Printf("%d\n", tripAss3(n, m, arr))
	fmt.Println(gcd([]int{12, 6, 9, 15}))
}

func TripAssessment() {
	//ass1Handle()
	//ass2Handle()
	ass3Handle()
}
