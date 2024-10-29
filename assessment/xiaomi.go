package assessment

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func xiaomiAss1Handle(data [3][3]int) int {
	digits := make([]int, 10)
	for i := range data {
		for j := range data[i] {
			digits[data[i][j]] = 1
		}
	}
	remainDigits := make([]int, 0)
	for i := range digits {
		if digits[i] == 0 {
			remainDigits = append(remainDigits, i)
		}
	}

	locations := make([][2]int, 0)
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 0 {
				locations = append(locations, [2]int{i, j})
			}
		}
	}

	check := func(a, b int) bool {
		if a-b == 1 || b-a == 1 {
			return false
		} else {
			return true
		}
	}
	res := 0
	var dfs func(int, []int)
	dfs = func(location int, remainNum []int) {
		if location == len(locations) {
			res++
			return
		}
		i, j := locations[location][0], locations[location][1]
		for idx := range remainNum {
			data[i][j] = remainNum[idx]
			remaining := append(remainNum[:idx], remainNum[idx+1:]...)
			if check(data[1][1], data[0][1]) &&
				check(data[1][1], data[2][1]) &&
				check(data[1][1], data[1][0]) &&
				check(data[1][1], data[1][2]) {
				dfs(location+1, remaining)
			}
		}
	}

	dfs(0, remainDigits)
	return res
}

func xiaomiAss1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		nGroup, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < nGroup; i++ {
			var data [3][3]int
			idx := 0
			for idx < 9 && scanner.Scan() {
				num, _ := strconv.Atoi(scanner.Text())
				data[idx/3][idx%3] = num
				idx++
			}
			fmt.Println(xiaomiAss1Handle(data))
		}
	}
	//data := [3][3]int{{1, 2, 3}, {0, 5, 8}, {0, 7, 0}}
	//fmt.Println(xiaomiAss1Handle(data))
}

func xiaomiAss2(a, n int) int {
	rotateRight := func(num int) int {
		x, k := num, 1
		for x/10 != 0 {
			k++
			x /= 10
		}
		leftPart := num % 10
		for i := 0; i < k-1; i++ {
			leftPart *= 10
		}
		return num/10 + leftPart
	}

	if n == 1 {
		return 0
	}

	queue := list.New()
	queue.PushBack([2]int{1, 0})

	visited := make(map[int]bool)
	visited[1] = true

	for queue.Len() > 0 {
		front := queue.Front()
		current, steps := front.Value.([2]int)[0], front.Value.([2]int)[1]
		queue.Remove(front)

		nextMult := current * a
		if nextMult == n {
			return steps + 1
		}
		if nextMult <= n*a && !visited[nextMult] {
			visited[nextMult] = true
			queue.PushBack([2]int{nextMult, steps + 1})
		}

		nextRot := rotateRight(current)
		if nextRot == n {
			return steps + 1
		}
		if !visited[nextRot] {
			visited[nextRot] = true
			queue.PushBack([2]int{nextRot, steps + 1})
		}
	}

	return -1
}

func XiaomiAssessment() {
	// xiaomiAss1()
	fmt.Println(xiaomiAss2(3, 621))
}
