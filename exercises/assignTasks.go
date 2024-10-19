package exercises

import (
	"container/heap"
	"fmt"
)

type PriorityQueue[T any] struct {
	items []T
	less  func(i, j T) bool
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

func (pq *PriorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if len(pq.items) == 0 {
		return *new(T), false
	}
	return pq.items[0], true
}

func NewPriorityQueue[T any](lessFunc func(i, j T) bool, initialItems []T) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		less:  lessFunc,
		items: initialItems,
	}

	heap.Init(pq)
	return pq
}

type IdleServer struct {
	weight int
	idx    int
}

func NewIdleServer(weight, idx int) *IdleServer {
	return &IdleServer{
		weight: weight,
		idx:    idx,
	}
}

type Worker struct {
	serverIdx      int
	completionTime int
}

func NewWorker(serverIdx, completionTime int) *Worker {
	return &Worker{
		serverIdx:      serverIdx,
		completionTime: completionTime,
	}
}

func assignTasks(servers []int, tasks []int) []int {
	ans := make([]int, len(tasks))
	idleServers := make([]IdleServer, len(servers))
	for idx, weight := range servers {
		idleServers[idx] = *NewIdleServer(weight, idx)
	}
	idleServerQueue := NewPriorityQueue(func(i, j IdleServer) bool {
		if i.weight == j.weight {
			return i.idx < j.idx
		}
		return i.weight < j.weight
	}, idleServers)

	taskQueue := NewPriorityQueue(func(i, j Worker) bool {
		return i.completionTime < j.completionTime
	}, nil)

	readyQueue := NewPriorityQueue(func(i, j int) bool {
		return i < j
	}, nil)

	for currTime := 0; currTime < len(tasks) || readyQueue.Len() != 0; currTime++ {
		if currTime < len(tasks) {
			heap.Push(readyQueue, currTime)
		} else {
			if elem, exist := taskQueue.Peek(); exist {
				currTime = elem.completionTime
			}
		}

		for elem, exist := taskQueue.Peek(); exist && elem.completionTime <= currTime; elem, exist = taskQueue.Peek() {
			elem = heap.Pop(taskQueue).(Worker)
			heap.Push(idleServerQueue, *NewIdleServer(servers[elem.serverIdx], elem.serverIdx))
		}

		for idleServerQueue.Len() != 0 && readyQueue.Len() != 0 {
			taskIdx := heap.Pop(readyQueue).(int)
			serverIdx := heap.Pop(idleServerQueue).(IdleServer).idx
			heap.Push(taskQueue, *NewWorker(serverIdx, currTime+tasks[taskIdx]))
			ans[taskIdx] = serverIdx
		}
	}
	return ans
}

func AssignTasks() {
	servers := []int{10, 63, 95, 16, 85, 57, 83, 95, 6, 29, 71}
	tasks := []int{70, 31, 83, 15, 32, 67, 98, 65, 56, 48, 38, 90, 5}
	fmt.Println("assignTasks", assignTasks(servers, tasks))
}
