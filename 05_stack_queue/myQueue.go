package __stack_queue

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) Peek() int { return s.data[len(s.data)-1] }

func (s *Stack) Empty() bool { return len(s.data) == 0 }

type MyQueue struct {
	stackIn  Stack
	stackOut Stack
}

func ConstructorQ() MyQueue {
	return MyQueue{
		stackIn: Stack{
			data: nil,
		},
		stackOut: Stack{
			data: nil,
		},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			this.stackOut.Push(this.stackIn.Pop())
		}
	}
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			this.stackOut.Push(this.stackIn.Pop())
		}
	}
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

func MyQueueTest() {
	obj := ConstructorQ()
	obj.Push(1)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
