package exercises

import (
	"fmt"
	"sort"
)

type segTree []struct {
	l, r, val int
}

func (t segTree) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t segTree) inc(o, i int) {
	if t[o].l == t[o].r {
		t[o].val++
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.inc(o<<1, i)
	} else {
		t.inc(o<<1|1, i)
	}
	t[o].val = t[o<<1].val + t[o<<1|1].val
}

func (t segTree) query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	n := len(nums)

	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}

	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	t := make(segTree, 4*k)
	t.build(1, 1, k)
	t.inc(1, kth[0])
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum-lower]
		cnt += t.query(1, left, right)
		t.inc(1, kth[sum])
	}
	return
}

func countRangeSumMerge(nums []int, lower, upper int) int {
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}

		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2)

		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}
	return mergeCount(prefixSum)
}

func CountRangeSum() {
	nums := []int{-2, 5, -1}
	lower, upper := -2, 2
	fmt.Println(countRangeSum(nums, lower, upper))
}
