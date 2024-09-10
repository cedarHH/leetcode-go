package __linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head.Val == val {
		head = head.Next
		if head == nil {
			return head
		}
	}
	prev, cur := head, head.Next
	for cur != nil {
		cur = cur.Next
		if prev.Next.Val == val {
			prev.Next = cur
		} else {
			prev = prev.Next
		}
	}
	return head
}

func RemoveElements() {
	head := &ListNode{Val: 1, Next: nil}
	tail := head

	elements := []int{2, 6, 3, 4, 5, 6}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	removeElements(head, 6)
}
