package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

type Bitset struct {
	isFlip int
	count  int
	size   int
	bits   []uint64
}

func Constructor(size int) Bitset {
	words := (size + 63) / 64
	return Bitset{
		isFlip: 0,
		count:  0,
		size:   size,
		bits:   make([]uint64, words),
	}
}

func (this *Bitset) Fix(idx int) {
	word, bit := idx/64, idx%64
	if this.isFlip == 0 && ((this.bits[word]>>bit)&1) == 0 {
		this.bits[word] |= 1 << bit
		this.count++
	} else if this.isFlip == 1 && ((this.bits[word]>>bit)&1) == 1 {
		this.bits[word] &^= 1 << bit
		this.count++
	}
}

func (this *Bitset) Unfix(idx int) {
	word, bit := idx/64, idx%64
	if this.isFlip == 0 && ((this.bits[word]>>bit)&1) == 1 {
		this.bits[word] &^= 1 << bit
		this.count--
	} else if this.isFlip == 1 && ((this.bits[word]>>bit)&1) == 0 {
		this.bits[word] |= 1 << bit
		this.count--
	}
}

func (this *Bitset) Flip() {
	this.isFlip = (this.isFlip + 1) % 2
	this.count = this.size - this.count
}

func (this *Bitset) All() bool {
	return this.count == this.size
}

func (this *Bitset) One() bool {
	return this.count > 0
}

func (this *Bitset) Count() int {
	return this.count
}

func (this *Bitset) ToString() string {
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	if this.isFlip == 1 {
		this.isFlip = 0
		for i := 0; i < len(this.bits); i++ {
			this.bits[i] = ^this.bits[i]
		}
		this.bits[len(this.bits)-1] &= ^uint64(0) >> ((64 - this.size%64) % 64)
	}

	var result strings.Builder
	totalWords := len(this.bits)

	for i := 0; i < totalWords-1; i++ {
		binaryStr := strconv.FormatUint(this.bits[i], 2)
		paddedBinaryStr := fmt.Sprintf("%064s", binaryStr)
		reversedBinaryStr := reverseString(paddedBinaryStr)
		result.WriteString(reversedBinaryStr)
	}

	lastBits := this.size % 64
	if lastBits == 0 {
		lastBits = 64
	}
	lastBinaryStr := strconv.FormatUint(this.bits[totalWords-1], 2)
	paddedLastStr := fmt.Sprintf("%064s", lastBinaryStr)
	reversedLastStr := reverseString(paddedLastStr)
	result.WriteString(reversedLastStr[:lastBits])
	return result.String()
}

func testBitset(operations []string, arguments [][]int) {
	var bitset Bitset
	for i, operation := range operations {
		fmt.Printf("%d %s\n", i, operation)
		switch operation {
		case "Bitset":
			bitset = Constructor(arguments[i][0])
			fmt.Println("Bitset as string:", bitset.ToString())
		case "fix":
			bitset.Fix(arguments[i][0])
			fmt.Println("Bitset as string:", bitset.ToString())
		case "unfix":
			bitset.Unfix(arguments[i][0])
			fmt.Println("Bitset as string:", bitset.ToString())
		case "flip":
			bitset.Flip()
			fmt.Println("Bitset as string:", bitset.ToString())
		case "all":
			fmt.Printf("%t\n", bitset.All())
			fmt.Println("Bitset as string:", bitset.ToString())
		case "one":
			bitset.One()
			fmt.Println("Bitset as string:", bitset.ToString())
		case "count":
			count := bitset.Count()
			fmt.Println("Bitset as string:", bitset.ToString())
			fmt.Println("Count of set bits:", count)
		case "toString":
			str := bitset.ToString()
			fmt.Println("Bitset as string:", str)
		}
	}
}

func BitSetTest() {
	operations := []string{"Bitset", "flip", "unfix", "all", "fix", "fix", "unfix", "all", "count", "toString", "toString", "toString", "unfix", "flip", "all", "unfix", "one", "one", "all", "fix", "unfix"}
	arguments := [][]int{{2}, {}, {1}, {}, {1}, {1}, {1}, {}, {}, {}, {}, {}, {0}, {}, {}, {0}, {}, {}, {}, {0}, {0}}
	testBitset(operations, arguments)
}
