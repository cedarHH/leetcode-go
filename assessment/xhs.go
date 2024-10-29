package assessment

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func XhsAssessment() {
	//xhsAss1()
	xhsAss2()
	//xhsAss3()
}

func xhsAss1() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	n, m := 0, 0
	if scanner.Scan() {
		size := strings.Split(scanner.Text(), " ")
		n, _ = strconv.Atoi(size[0])
		m, _ = strconv.Atoi(size[1])
	}
	world := make([][]int, n)
	for i := range world {
		world[i] = make([]int, m)
		if scanner.Scan() {
			row := scanner.Text()
			for j := range row {
				switch row[j] {
				case 'L':
					world[i][j] = 0
				case 'R':
					world[i][j] = 1
				case 'U':
					world[i][j] = 2
				case 'D':
					world[i][j] = 3
				default:
					world[i][j] = -1
				}
			}
		}
	}
	fmt.Println(xhsAss1Handle(n, m, world))
}

func xhsAss1Handle(n, m int, world [][]int) (remaining int) {
	existsRobot := make([][]int, n)
	for i := 0; i < n; i++ {
		existsRobot[i] = make([]int, m)
	}
	for i := range world {
		for j := range world[0] {
			visited := make([][]int, n)
			for k := 0; k < n; k++ {
				visited[k] = make([]int, m)
			}
			visited[i][j] = 1
			x, y := i, j
			for k := 0; k < 10e8; k++ {
				switch world[x][y] {
				case 0:
					y--
				case 1:
					y++
				case 2:
					x--
				case 3:
					x++
				}
				if x < 0 || y < 0 || x >= n || y >= m || existsRobot[x][y] == -1 {
					existsRobot[i][j] = -1
					break
				} else if visited[x][y] == 1 || existsRobot[x][y] == 1 {
					for a := range visited {
						for b := range visited[0] {
							if visited[a][b] == 1 {
								existsRobot[a][b] = 1
							}
						}
					}
					remaining++
					break
				} else {
					visited[x][y] = 1
				}
			}
		}
	}
	return remaining
}

func xhsAss2() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanWords)
	//n := 0
	//if scanner.Scan() {
	//	n, _ = strconv.Atoi(scanner.Text())
	//}
	//for i := 0; i < n && scanner.Scan(); i++ {
	//	num, _ := strconv.Atoi(scanner.Text())
	//	xhsAss2Handle(num)
	//}

	t0 := time.Now()
	for i := 10; i < 10e3; i++ {
		xhsAss2Handle(i)
	}
	fmt.Println(time.Now().Sub(t0).Seconds())
	t0 = time.Now()
	for i := 10; i < 10e4; i++ {
		findPythagoreanTriplet(i)
	}
	fmt.Println(time.Now().Sub(t0).Seconds())
}

func xhsAss2Handle(num int) (int, int, int) {
	maxSum := 3 * num
	res := [3]int{3 * num, 3 * num, 3 * num}
	for i := 1; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if c := math.Sqrt(float64(i*i + j*j)); c == float64(int(c)) {
				k := int(c)
				if i > res[0] && j > res[1] && k > res[2] {
					// fmt.Println(res[0], res[1], res[2])
					return res[0], res[1], res[2]
				}
				if i*i+j*j == k*k && num <= i+j+k {
					if i+j+k < maxSum {
						maxSum = i + j + k
						res[0], res[1], res[2] = i, j, k
					}
				}
			}
		}
	}
	//fmt.Println(res[0], res[1], res[2])
	return res[0], res[1], res[2]
}

func findPythagoreanTriplet(n int) (int, int, int) {
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}

	smallestSum := math.MaxInt
	var smallestA, smallestB, smallestC int

	m := 2
	for m = 2; ; m++ {
		for k := 1; k < m; k++ {
			// 确保 m 和 k 互质，且 m - k 为奇数
			if (m-k)%2 == 1 && gcd(m, k) == 1 {
				// 生成本原三元组
				a := m*m - k*k
				b := 2 * m * k
				c := m*m + k*k

				// 对本原三元组进行缩放
				scale := 1
				for {
					scaledA := scale * a
					scaledB := scale * b
					scaledC := scale * c
					sum := scaledA + scaledB + scaledC

					// 如果满足条件并且总和比当前的最小值更小，更新结果
					if sum > n && sum < smallestSum {
						smallestA = scaledA
						smallestB = scaledB
						smallestC = scaledC
						smallestSum = sum
					}

					// 防止无限循环，如果缩放过多，跳出循环
					if sum > 2*n {
						break
					}
					scale++
				}
			}
		}

		// 当找到的最小和的三元组已经比 n 大时，停止继续查找
		if smallestSum != math.MaxInt {
			break
		}
	}

	return smallestA, smallestB, smallestC
}

func xhsAss3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	n, b := 0, 0
	if scanner.Scan() {
		size := strings.Split(scanner.Text(), " ")
		n, _ = strconv.Atoi(size[0])
		b, _ = strconv.Atoi(size[1])
	}
	priceList := make([][3]int, n)
	if scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		priceList[0][0], _ = strconv.Atoi(row[0])
		priceList[0][1], _ = strconv.Atoi(row[1])
	}
	for i := 1; i < n && scanner.Scan(); i++ {
		row := strings.Split(scanner.Text(), " ")
		priceList[i][0], _ = strconv.Atoi(row[0])
		priceList[i][1], _ = strconv.Atoi(row[1])
		priceList[i][2], _ = strconv.Atoi(row[2])
	}
	fmt.Println(xhsAss3Handle(n, priceList, b))
}

func xhsAss3Handle(n int, priceList [][3]int, b int) (res int) {
	items := make([]bool, n)

	var dfs func(int, int, int, []bool)
	dfs = func(idx, price, maxNum int, itemList []bool) {
		if idx >= n || price == b {
			return
		}
		temp := make([]bool, n)
		copy(temp, itemList)
		if priceList[idx][2] == 0 && price+priceList[idx][0]-priceList[idx][1] <= b {
			temp[idx] = true
			if maxNum+1 > res {
				res = maxNum + 1
			}
			dfs(idx+1, price+priceList[idx][0]-priceList[idx][1], maxNum+1, temp)
		} else if itemList[priceList[idx][2]-1] && price+priceList[idx][0]-priceList[idx][1] <= b {
			temp[idx] = true
			if maxNum+1 > res {
				res = maxNum + 1
			}
			dfs(idx+1, price+priceList[idx][0]-priceList[idx][1], maxNum+1, temp)
		} else if !itemList[priceList[idx][2]-1] && price+priceList[idx][0] <= b {
			temp[idx] = true
			if maxNum+1 > res {
				res = maxNum + 1
			}
			dfs(idx+1, price+priceList[idx][0], maxNum+1, temp)
		}
		dfs(idx+1, price, maxNum, itemList)
	}

	dfs(0, 0, 0, items)
	return res
}
