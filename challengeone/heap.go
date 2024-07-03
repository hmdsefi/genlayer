package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	key   int
	value *DataHash
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].key > pq[j].key
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// InsertMapIntoHeap inserts a map[int]string into a heap ordered by map key
func InsertMapIntoHeap(m map[int]*DataHash) *PriorityQueue {
	pq := make(PriorityQueue, 0, len(m))
	heap.Init(&pq)
	for key, value := range m {
		heap.Push(&pq, &Item{key: key, value: value})
	}
	return &pq
}

// Helper function to print heap contents
func printHeap(pq *PriorityQueue) {
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("key: %d, value: %s\n", item.key, item.value)
	}
}
