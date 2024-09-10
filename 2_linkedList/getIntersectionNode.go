package __linkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	currA := headA
	for currA != nil {
		currA = currA.Next
		lenA++
	}
	currB := headB
	for currB != nil {
		currB = currB.Next
		lenB++
	}
	currA, currB = headA, headB
	if lenA > lenB {
		n := lenA - lenB
		for n != 0 {
			currA = currA.Next
			n--
		}
	} else {
		n := lenB - lenA
		for n != 0 {
			currB = currB.Next
			n--
		}
	}
	for currA != currB {
		currA = currA.Next
		currB = currB.Next
	}
	return currA
}

func GetIntersectionNode() {
	public := &ListNode{Val: 6, Next: nil}
	tail := public

	elements := []int{7, 8, 9, 10}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}

	headA := &ListNode{Val: 1, Next: nil}
	tail = headA

	elements = []int{2, 3, 4, 5}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	tail.Next = public

	headB := &ListNode{Val: 11, Next: nil}
	tail = headB

	elements = []int{12, 13}
	for _, element := range elements {
		tail.Next = new(ListNode)
		tail = tail.Next
		tail.Val = element
	}
	tail.Next = public

	getIntersectionNode(headA, headB)
}
