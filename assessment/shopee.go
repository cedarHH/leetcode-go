package assessment

import "fmt"

func ass1(arr []int) (res int) {
	//length := len(arr)
	//if length < 2 {
	//	return 0
	//}
	//
	//dp := make([]int, length)
	//
	//for idx := 1; idx < length; idx++ {
	//	if dp[idx] > dp[idx-1] {
	//		dp[idx] = 2*dp[idx-1] + 1
	//	} else if dp[idx] == dp[idx-1] {
	//		dp[idx] = dp[idx-1] + 1
	//	} else {
	//		for j := idx - 1; j >= 0; j-- {
	//			if dp[idx] >
	//		}
	//	}
	//}
	//
	//for _, val := range dp {
	//	res += val
	//}
	return
}

func jump(steps []int) int {
	// write code here
	minInt := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	length := len(steps)
	dp := make([]int, length)
	for idx := range dp {
		dp[idx] = 100001
	}
	dp[0] = 0
	for idx := range dp {
		for j := range dp[idx:minInt(idx+steps[idx]+1, length)] {
			dp[idx+j] = minInt(dp[idx+j], dp[idx]+1)
		}
	}
	return dp[length-1]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (l *ListNode) append(val int) {
	curr := l
	for ; curr.Next != nil; curr = curr.Next {
	}
	curr.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (l *ListNode) showList() {

	for curr := l; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}

func removeNode(node *ListNode, n int) *ListNode {
	// write code here
	slow, fast := node, node
	for _ = range n - 1 {
		fast = fast.Next
		if fast == nil {
			return node
		}
	}
	if fast.Next == nil {
		return node.Next
	}
	fast = fast.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return node
}

func ShopeeAssessment() {
	fmt.Println("ass1:", ass1([]int{0, 2, 7}))
	fmt.Println("ass2", jump([]int{3, 1, 2, 2, 3, 3, 2, 4, 0, 2, 2, 0}))
	arr := []int{1, 2, 3, 4, 5}
	l := NewList(arr[0])
	for _, val := range arr[1:] {
		l.append(val)
	}
	l.showList()
	l2 := removeNode(l, 2)
	l2.showList()
}
