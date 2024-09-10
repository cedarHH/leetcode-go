package __linkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev := head
	cur := head.Next
	prev.Next = nil
	for cur != nil {
		temp := cur
		cur = cur.Next
		temp.Next = prev
		prev = temp
	}
	return prev
}

func ReverseList() {
	head := &ListNode{Val: 1, Next: nil}
	tail := head

	elements := []int{2, 3, 4, 5}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	reverseList(head)
}
