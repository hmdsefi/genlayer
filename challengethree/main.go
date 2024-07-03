package main

import (
	"fmt"
)

func main() {
	graph := Graph{
		"A": []Edge{{"B", 10}, {"C", 20}},
		"B": []Edge{{"D", 15}},
		"C": []Edge{{"D", 30}},
		"D": []Edge{},
	}

	start := "A"
	end := "D"

	path := Dijkstra(graph, start, end)
	if path != nil {
		fmt.Printf("fastest path from %s to %s: %v\n", start, end, path)
	} else {
		fmt.Printf("No path found from %s to %s\n", start, end)
	}
}
