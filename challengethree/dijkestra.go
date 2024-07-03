package main

import (
	"container/heap"
)

func Dijkstra(graph Graph, start, end string) []string {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	initialPriority := map[string]int{}
	initialPriority[start] = 0
	heap.Push(&pq, &Item{node: start, latency: 0, priority: 0})

	previous := map[string]*Item{}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		if current.node == end {
			var path []string
			for current != nil {
				path = append([]string{current.node}, path...)
				current = previous[current.node]
			}
			return path
		}

		for _, edge := range graph[current.node] {
			newLatency := current.latency + edge.Latency
			if priority, ok := initialPriority[edge.To]; !ok || newLatency < priority {
				initialPriority[edge.To] = newLatency
				item := &Item{node: edge.To, latency: newLatency, priority: newLatency}
				previous[edge.To] = current
				heap.Push(&pq, item)
			}
		}
	}

	return nil
}
