package utils

import "container/heap"

type Item struct {
	Id     int
	Weight float64
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest weight Tile, so we use less than here.
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type hp struct {
	Items *PriorityQueue
}

func NewHeap() *hp {
	return &hp{Items: &PriorityQueue{}}
}

func (hp *hp) Push(item *Item) {
	heap.Push(hp.Items, item)
}

func (hp *hp) Pop() *Item {
	return heap.Pop(hp.Items).(*Item)
}
