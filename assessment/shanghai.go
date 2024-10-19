package assessment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShanghaiAssessment() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if num, err := strconv.Atoi(scanner.Text()); err == nil {
			fmt.Println(shanghaiAss1Handle(num))
		}
	}
}

func shanghaiAss1Handle(num int) int {
	isPrime := make([]bool, num+1)
	dp := make([]int, num+1)
	for i := 2; i < num+1; i++ {
		isPrime[i] = true
		dp[i] = 1
	}
	for i := 2; i*i < num+1; i++ {
		for j := i * i; j < num+1; j += i {
			isPrime[j] = false
			dp[j] = 0
		}
	}
	primes := make([]int, 0)
	for idx, val := range isPrime {
		if val {
			primes = append(primes, idx)
		}
	}
	for idx, prime1 := range primes[:len(primes)-1] {
		if dp[prime1] > 1 && prime1*2 <= num {
			dp[prime1+prime1] += 1
		}
		for _, prime2 := range primes[idx+1:] {
			if prime1+prime2 <= num {
				dp[prime1+prime2] += dp[prime1]
			}
		}
	}
	fmt.Println(dp)
	return dp[num]
}
