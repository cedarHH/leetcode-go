package __array

import "fmt"

func generateMatrix(n int) [][]int {
	x, y, value, direction := 0, 0, 1, 0
	upper, lower, left, right := 0, n-1, 0, n-1
	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for value != n*n+1 {
		matrix[x][y] = value
		value++
		x += move[direction][0]
		y += move[direction][1]
		if x < upper || x > lower || y < left || y > right {
			switch direction {
			case 0:
				upper++
			case 1:
				right--
			case 2:
				lower--
			case 3:
				left++
			}
			x -= move[direction][0]
			y -= move[direction][1]
			direction = (direction + 1) % 4
			x += move[direction][0]
			y += move[direction][1]

		}
	}
	return matrix
}

func GenerateMatrix() {
	fmt.Println(generateMatrix(4))
}
