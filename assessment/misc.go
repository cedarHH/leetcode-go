package assessment

import (
	"container/heap"
	"fmt"
)

type Person struct {
	age  int
	name string
}

type PersonHeap []*Person

func (p *PersonHeap) Len() int               { return len(*p) }
func (p *PersonHeap) Less(i int, j int) bool { return (*p)[i].age < (*p)[j].age }
func (p *PersonHeap) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *PersonHeap) Push(x any)             { *p = append(*p, x.(*Person)) }
func (p *PersonHeap) Pop() any {
	n := len(*p)
	item := (*p)[n-1]
	*p = (*p)[0 : n-1]
	return item
}

func Misc() {
	p := &PersonHeap{
		&Person{
			age:  4,
			name: "b",
		},
		&Person{
			age:  1,
			name: "c",
		},
		&Person{
			age:  2,
			name: "a",
		},
	}
	heap.Init(p)
	heap.Push(p, &Person{
		age:  3,
		name: "d",
	})
	for p.Len() > 0 {
		item := heap.Pop(p).(*Person)
		fmt.Println(item.age, ":", item.name)
	}
}
