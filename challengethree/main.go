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

	path := FindMinimumLatencyPath(graph, compressionNodes, start, end)
	if path != nil {
		fmt.Printf("fastest path from %s to %s: %v\n", start, end, path)
	} else {
		fmt.Printf("No path found from %s to %s\n", start, end)
	}
}
