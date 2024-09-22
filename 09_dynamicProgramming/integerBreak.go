package __dynamicProgramming

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	}
	power, remainder := n/3, n%3
	if remainder == 1 {
		return int(math.Pow(3, float64(power-1))) * 4
	} else if remainder == 0 {
		return int(math.Pow(3, float64(power)))
	} else {
		return int(math.Pow(3, float64(power))) * 2
	}
}

func IntegerBreak() {
	fmt.Println("integerBreak: ", integerBreak(10))
}
