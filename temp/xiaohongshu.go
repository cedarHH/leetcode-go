package temp

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func redBookAss1() {

	avg := func(arr []int) int {
		if len(arr) == 0 {
			return 0
		}
		if len(arr) == 1 {
			return arr[0]
		}
		sum := 0
		for i := 1; i < len(arr)-1; i++ {
			sum += arr[i]
		}
		return int(math.Ceil(float64(sum) / float64(len(arr)-2)))
	}

	n, m := 0, 0
	//_, _ = fmt.Scan(&n, &m)
	n, m = 4, 5
	scores := make([][]int, m)
	for i := 0; i < m; i++ {
		scores[i] = make([]int, n)
	}
	mock := []string{"abaKm", "abAfe", "czaCd", "DbAVa"}
	for i := 0; i < n; i++ {
		str := ""
		//_, _ = fmt.Scan(&str)
		str = mock[i]
		str = strings.ToLower(str)
		for j := 0; j < m; j++ {
			scores[j][i] = int(str[j] - 'a')
		}
	}
	for i := 0; i < m; i++ {
		sort.Ints(scores[i])
	}
	result := make([]int, m)
	orders := make(map[int][]string)
	for i := range result {
		result[i] = avg(scores[i])
		orders[result[i]] = append(orders[result[i]], string(uint8(i)+'1'))
	}
	sortedResult := make([]int, m)
	copy(sortedResult, result)
	sort.Ints(sortedResult)
	chars := make([]string, m)
	for i := range sortedResult {
		chars[i] = string(uint8(sortedResult[i]) + 'a')
	}

	orderArr := make([]string, 0)
	for i := range sortedResult {
		if i == 0 || sortedResult[i-1] != sortedResult[i] {
			orderArr = append(orderArr, orders[sortedResult[i]]...)
		}
	}

	fmt.Println(strings.Join(orderArr, " "))
	fmt.Println(strings.Join(chars, " "))
}

func redBookAss2() {

}

func RedBookAssessment() {
	redBookAss1()
}
