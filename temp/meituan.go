package temp

import (
	"fmt"
	"sort"
)

func mtAss1Handle(arr []int) (res int) {

	isGoodNum := func(num int) bool {
		odd := false
		sum := 0
		for num > 0 {
			digit := num % 10
			num /= 10
			sum += digit
			if digit%2 == 1 {
				odd = true
			}
		}
		if sum%2%2 == 0 && odd {
			return true
		}
		return false
	}
	for i := range arr {
		if isGoodNum(arr[i]) {
			res++
		}
	}

	return res
}

func mtAss1() {
	n := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		arr[i] = x
	}
	fmt.Println(mtAss1Handle(arr))
}

func mtAss2Handle(n int, str string) {
	arr := make([]int, 0)
	num := 0
	for idx, chr := range str {
		if chr >= '0' && chr <= '9' {
			num *= 10
			num += int(chr - '0')
		} else if idx > 0 && str[idx-1] >= '0' && str[idx-1] <= '9' {
			arr = append(arr, num)
			num = 0
		}
	}
	if str[len(str)-1] >= '0' && str[len(str)-1] <= '9' {
		arr = append(arr, num)
	}
	if len(arr) < n {
		fmt.Println("N")
	} else {
		sort.Ints(arr)
		fmt.Println(arr)
		fmt.Println(arr[len(arr)-n])
	}
}

func mtAss2() {
	n, str := 3, "a12b03ccc00002c123c3d0d0d4"
	//_, _ = fmt.Scan(&n)
	//_, _ = fmt.Scan(&str)
	mtAss2Handle(n, str)
}

func mtAss3Handle(arrA, arrB []int) {
	isPrime1 := func(num int) bool {
		if num <= 1 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	findIdx := func(num int, arr []int) int {
		for i := range arr {
			if arr[i] == num {
				arr[i] = -1
				return i
			}
		}
		return -1
	}

	min1 := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	max1 := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	var arrAP, arrANp, arrBP, arrBNp []int
	res := make([]int, len(arrA))
	sum := 0
	for _, val := range arrA {
		if isPrime1(val) {
			arrAP = append(arrAP, val)
		} else {
			arrANp = append(arrANp, val)
		}
	}
	for _, val := range arrB {
		if isPrime1(val) {
			arrBP = append(arrBP, val)
		} else {
			arrBNp = append(arrBNp, val)
		}
	}
	sort.Ints(arrAP)
	sort.Ints(arrBP)
	sort.Ints(arrANp)
	sort.Ints(arrBNp)
	var extraANp, extraBNp []int
	if len(arrANp) < len(arrBNp) {
		for idx, val := range arrANp {
			sum += val
			sum += arrBNp[idx]
			res[findIdx(val, arrA)] = findIdx(arrBNp[idx], arrB)
		}
		extraBNp = append(extraBNp, arrBNp[len(arrANp):]...)
	} else {
		for idx, val := range arrBNp {
			sum += val
			sum += arrANp[idx]
			res[findIdx(arrANp[idx], arrA)] = findIdx(val, arrB)
		}
		extraANp = append(extraANp, arrANp[len(arrBNp):]...)
	}
	if len(arrAP) < len(arrBP) {
		for idx := 0; idx < len(arrAP); idx++ {
			sum = sum + (arrAP[len(arrAP)-idx-1]+arrBP[len(arrBP)-idx-1])*2
			res[findIdx(arrAP[len(arrAP)-idx-1], arrA)] = findIdx(arrBP[len(arrBP)-idx-1], arrB)
		}
	} else {
		for idx := 0; idx < len(arrBP); idx++ {
			sum = sum + (arrAP[len(arrAP)-idx-1]+arrBP[len(arrBP)-idx-1])*2
			res[findIdx(arrAP[len(arrAP)-idx-1], arrA)] = findIdx(arrBP[len(arrBP)-idx-1], arrB)
		}
	}
	fmt.Println(sum)
	fmt.Println(res)
}

func mtAss3() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		arrA[i] = x
	}
	arrB := make([]int, m)
	for i := 0; i < m; i++ {
		x := 0
		fmt.Scan(&x)
		arrB[i] = x
	}
	mtAss3Handle(arrA, arrB)
}

func MeituanAssessment() {
	// mtAss1()
	// mtAss2()
	mtAss3()
}
