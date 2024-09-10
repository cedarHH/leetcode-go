package assessment

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func minAreaRect1(points [][]int) int {
	type Line struct {
		Y, L int
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	var groupedPoints [][][]int
	var currentGroup [][]int
	for i, point := range points {
		if i == 0 || point[0] == points[i-1][0] {
			currentGroup = append(currentGroup, point)
		} else {
			groupedPoints = append(groupedPoints, currentGroup)
			currentGroup = [][]int{point}
		}
	}
	groupedPoints = append(groupedPoints, currentGroup)
	minArea := math.MaxInt
	lineMap := map[Line][]int{}

	sliceMax := func(slice []int) int {
		maxVal := -1
		for _, v := range slice {
			maxVal = max(v, maxVal)
		}
		return maxVal
	}

	for _, pointsSlice := range groupedPoints {
		for j := 0; j < len(pointsSlice); j++ {
			for k := j + 1; k < len(pointsSlice); k++ {
				line := Line{
					Y: pointsSlice[j][1],
					L: pointsSlice[k][1] - pointsSlice[j][1],
				}
				if lines, exists := lineMap[line]; exists {
					xMax := sliceMax(lines)
					if xMax != -1 {
						area := (pointsSlice[j][0] - xMax) * line.L
						minArea = min(minArea, area)
					}
					lineMap[line] = append(lineMap[line], pointsSlice[j][0])
				} else {
					lineMap[line] = []int{pointsSlice[j][0]}
				}

			}
		}
	}
	if minArea == math.MaxInt {
		minArea = 0
	}
	return minArea
}

func minAreaRect(points [][]int) int {
	col := map[int]map[int]bool{}
	for _, point := range points {
		if col[point[0]] == nil {
			col[point[0]] = make(map[int]bool)
		}
		col[point[0]][point[1]] = true
	}
	var keys []int
	for key := range col {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	minArea := math.MaxInt
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			intersection := make(map[int]bool)
			for key := range col[keys[i]] {
				if _, exists := col[keys[j]][key]; exists {
					intersection[key] = true
				}
			}
			var yKeys []int
			for key := range intersection {
				yKeys = append(yKeys, key)
			}
			sort.Ints(yKeys)
			minDiff := math.MaxInt
			for k := 1; k < len(yKeys); k++ {
				diff := yKeys[k] - yKeys[k-1]
				if diff < minDiff {
					minDiff = diff
				}
			}
			if minDiff != math.MaxInt {
				minArea = min(minArea, minDiff*(keys[j]-keys[i]))
			}
		}
	}

	if minArea == math.MaxInt {
		minArea = 0
	}
	return minArea
}

func MinAreaRect() {
	var points [][]int
	input := "[[1,1],[1,3],[3,1],[3,3],[2,2]]"

	re := regexp.MustCompile(`\[(\d+),(\d+)\]`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		points = append(points, []int{a, b})
	}

	fmt.Printf("%d\n", minAreaRect(points))
}
