package __stack_queue

import (
	"container/list"
	"fmt"
)

type Deque struct {
	items *list.List
}

func NewDeque() *Deque {
	return &Deque{items: list.New()}
}

func (d *Deque) PushBack(item int) {
	d.items.PushBack(item)
}

func (d *Deque) PopFront() int {
	if d.items.Len() == 0 {
		return -1
	}
	front := d.items.Front()
	d.items.Remove(front)
	return front.Value.(int)
}

func (d *Deque) IsEmpty() bool {
	return d.items.Len() == 0
}

type MyStack struct {
	queue *Deque
}

func Constructor() MyStack {
	return MyStack{
		queue: NewDeque(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyStack) Pop() int {
	n := this.queue.items.Len() - 1
	for n != 0 {
		val := this.queue.PopFront()
		this.queue.PushBack(val)
		n--
	}
	val := this.queue.PopFront()
	return val

}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue.PushBack(val)
	return val
}

func (this *MyStack) Empty() bool {
	return this.queue.items.Len() == 0
}

func MyStackTest() {
	obj := Constructor()
	obj.Push(1)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
