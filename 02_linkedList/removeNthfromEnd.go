package _2_linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return slow.Next
	}
	fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func RemoveNthFromEnd() {
	head := &ListNode{Val: 1, Next: nil}
	tail := head

	elements := []int{2, 3, 4, 5}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	removeNthFromEnd(head, 2)
}
