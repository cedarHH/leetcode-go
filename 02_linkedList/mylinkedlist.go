package _2_linkedList

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	head := new(MyLinkedList)
	return *head
}

func (this *MyLinkedList) Get(index int) int {
	cur := this
	for i := 0; i < index; i++ {
		cur = cur.Next
		if cur == nil {
			return -1
		}
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Next = &MyLinkedList{
		Val:  val,
		Next: this.Next,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	prev := this
	for ; prev.Next != nil; prev = prev.Next {
	}
	prev.Next = &MyLinkedList{
		Val:  val,
		Next: nil,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this
	for i := 0; i < index; i++ {
		cur = cur.Next
		if cur == nil {
			return
		}
	}
	newNode := MyLinkedList{
		Val:  val,
		Next: cur.Next,
	}
	cur.Next = &newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev := this
	cur := this.Next
	for i := 0; i < index; i++ {
		prev = cur
		cur = cur.Next
		if cur == nil {
			return
		}
	}
	prev.Next = cur.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func LinkedList() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.Get(1)
	obj.DeleteAtIndex(1)
	obj.Get(1)
	obj.Get(3)
	obj.DeleteAtIndex(3)
	obj.DeleteAtIndex(0)
	obj.Get(0)
	obj.DeleteAtIndex(0)
	obj.Get(0)
}
