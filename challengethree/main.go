package main

import (
	"fmt"
)

// Optimal Network Routing
func main() {
	graph := Graph{
		"A": []Edge{{"B", 10}, {"C", 20}},
		"B": []Edge{{"D", 15}},
		"C": []Edge{{"D", 30}},
		"D": []Edge{},
	}

	start := "A"
	end := "D"
	compressionNodes := []string{"B", "C"}

	totalLatency := FindMinimumLatencyPath(graph, compressionNodes, start, end)
	if totalLatency != 0 {
		fmt.Printf("minimum total latency %d", totalLatency)
	} else {
		fmt.Printf("No path found from %s to %s\n", start, end)
	}
}
