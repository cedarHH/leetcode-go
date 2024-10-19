package assessment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baiduAss1Handle(n, m int) int {
	maxScores := (n + 1) >> 2
	if m <= maxScores {
		return m
	} else {
		return 2*maxScores - m + (n+1)%2
	}
}

func baiduAss1() {
	nRow := 0
	scanner := bufio.NewScanner(os.Stdin)
	nRow, _ = strconv.Atoi(scanner.Text())
	res := make([]string, nRow)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < nRow; i++ {
		strArr := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(strArr[0])
		m, _ := strconv.Atoi(strArr[1])
		res[i] = strconv.Itoa(baiduAss1Handle(n, m))
	}
	fmt.Println(strings.Join(res, "\n"))
}

func baiduAss2() {
	str := ""
	_, _ = fmt.Scan(&str)
	swap := func(idx int) {
		str = str[:idx] + str[idx+1:] + string(str[idx])
	}
	for i := 0; i < len(str); i++ {
		swap(i)
	}
	fmt.Println(str)
}

func BaiduAssessment() {
	// baiduAss1()
	baiduAss2()
}
