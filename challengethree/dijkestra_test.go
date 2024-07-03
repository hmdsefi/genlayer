package main

import (
	"reflect"
	"testing"
)

func TestFindMinimumLatencyPath_CompressionNodes(t *testing.T) {
	graph := Graph{
		"A": []Edge{{"B", 15}, {"C", 12}},
		"B": []Edge{{"D", 15}},
		"C": []Edge{{"D", 22}},
		"D": []Edge{},
	}

	tests := []struct {
		name         string
		start, end   string
		compressions []string
		expectedPath int
	}{
		{
			name:         "Path through compressed node B",
			start:        "A",
			end:          "D",
			compressions: []string{"B"},
			expectedPath: 22,
		},
		{
			name:         "Path through compressed node C",
			start:        "A",
			end:          "D",
			compressions: []string{"C"},
			expectedPath: 28,
		},
		{
			name:         "No path between A and D without compression",
			start:        "A",
			end:          "D",
			expectedPath: 30,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				path := FindMinimumLatencyPath(graph, tt.compressions, tt.start, tt.end)
				if !reflect.DeepEqual(path, tt.expectedPath) {
					t.Errorf("Test %s failed: expected path %v, got %v", tt.name, tt.expectedPath, path)
				}
			},
		)
	}
}
