package main

import (
	"container/heap"
	"math"
)

func FindMinimumLatencyPath(graph Graph, compressionNodes []string, start, end string) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	initialPriority := map[string]int{}
	initialPriority[start] = 0
	heap.Push(&pq, &Item{node: start, latency: 0, priority: 0, usedCompression: false})

	previous := map[string]*Item{}

	// Convert compression nodes to a map for quick lookup
	compressionSet := make(map[string]bool)
	for _, node := range compressionNodes {
		compressionSet[node] = true
	}

	var totalLatency int

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		// Check if the current node is the end node
		if current.node == end {
			var path []string
			for current != nil {
				path = append([]string{current.node}, path...)
				current = previous[current.node]
			}
			return totalLatency
		}

		// Explore neighbors
		for _, edge := range graph[current.node] {
			// Skip compression nodes
			if compressionSet[edge.To] && !current.usedCompression {
				newLatency := current.latency + int(math.Floor(float64(edge.Latency)/2.0))
				if priority, ok := initialPriority[edge.To]; !ok || newLatency < priority {
					initialPriority[edge.To] = newLatency
					item := &Item{
						node:            edge.To,
						latency:         newLatency,
						priority:        newLatency,
						usedCompression: true,
					}
					previous[edge.To] = current
					heap.Push(&pq, item)
					totalLatency = newLatency
				}
			} else {
				newLatency := current.latency + edge.Latency
				if priority, ok := initialPriority[edge.To]; !ok || newLatency < priority {
					initialPriority[edge.To] = newLatency
					item := &Item{
						node:            edge.To,
						latency:         newLatency,
						priority:        newLatency,
						usedCompression: current.usedCompression,
					}
					previous[edge.To] = current
					heap.Push(&pq, item)
					totalLatency = newLatency
				}
			}
		}
	}

	return totalLatency
}
