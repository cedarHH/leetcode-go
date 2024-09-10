package __linkedList

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func DetectCycle() {
	head := &ListNode{Val: 1, Next: nil}
	tail := head

	elements := []int{2, 3, 4, 5}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	entry := tail
	elements = []int{6, 7, 8, 9, 10}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	tail.Next = entry
	detectCycle(head)
}
