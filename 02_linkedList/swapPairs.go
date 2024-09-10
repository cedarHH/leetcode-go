package _2_linkedList

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head.Next
	head.Next = prev.Next
	prev.Next = head
	head = prev
	prev = prev.Next
	curr := prev.Next
	for curr != nil && curr.Next != nil {
		temp := curr.Next
		curr.Next = curr.Next.Next
		temp.Next = curr
		prev.Next = temp
		prev = curr
		curr = prev.Next
	}
	return head
}

func SwapPairs() {
	head := &ListNode{Val: 1, Next: nil}
	tail := head

	elements := []int{2, 3, 4, 5}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	swapPairs(head)
}
