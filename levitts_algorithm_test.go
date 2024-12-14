package main

import (
	"math"
	"reflect"
	"testing"
)

func TestLevit(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][3]int
		source   int
		expected []int
	}{
		{
			name:     "small graph",
			vertices: 5,
			edges: [][3]int{
				{0, 1, 10},
				{0, 2, 5},
				{1, 3, 1},
				{2, 1, 3},
				{2, 3, 8},
				{2, 4, 2},
				{3, 4, 4},
				{4, 3, 6},
			},
			source:   0,
			expected: []int{0, 8, 5, 9, 7},
		},
		{
			name:     "no path graph",
			vertices: 3,
			edges:    [][3]int{},
			source:   0,
			expected: []int{0, math.MaxInt32, math.MaxInt32},
		},
		{
			name:     "single node graph",
			vertices: 1,
			edges:    [][3]int{},
			source:   0,
			expected: []int{0},
		},
		{
			name:     "disconnected graph",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 1},
				{2, 3, 1},
			},
			source:   0,
			expected: []int{0, 1, math.MaxInt32, math.MaxInt32},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			graph := NewGraph(test.vertices)
			for _, edge := range test.edges {
				graph.AddEdge(edge[0], edge[1], edge[2])
			}
			result := graph.Levit(test.source)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected distances %v, got %v", test.expected, result)
			}
		})
	}
}
