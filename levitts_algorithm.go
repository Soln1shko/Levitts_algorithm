package main

import (
	"container/list"
	"fmt"
	"math"
)

// Edge представляет ребро графа
type Edge struct {
	to     int
	weight int
}

// Graph представляет неориентированный граф на списке смежности
type Graph struct {
	vertices int
	edges    [][]Edge
}

// NewGraph создает новый граф с заданным количеством вершин
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]Edge, vertices),
	}
}

// AddEdge добавляет новое ребро в граф
func (g *Graph) AddEdge(from, to, weight int) {
	g.edges[from] = append(g.edges[from], Edge{to, weight})
	g.edges[to] = append(g.edges[to], Edge{from, weight})
}

// Levit реализует алгоритм Левита для поиска кратчайших путей от единственного источника
func (g *Graph) Levit(source int) []int {
	distances := make([]int, g.vertices)
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[source] = 0

	// Множества Q и M
	inQueue := make([]bool, g.vertices)
	isProcessed := make([]bool, g.vertices)

	// Дек для хранения вершин
	deque := list.New()
	deque.PushBack(source)

	for deque.Len() > 0 {
		// Вытаскиваем вершину из дека
		vertex := deque.Remove(deque.Front()).(int)

		// Обработка всех соседей
		for _, edge := range g.edges[vertex] {
			newDist := distances[vertex] + edge.weight
			if newDist < distances[edge.to] {
				distances[edge.to] = newDist
				if !inQueue[edge.to] {
					// Не в очереди
					inQueue[edge.to] = true
					if isProcessed[edge.to] {
						// Вершина обработана
						deque.PushFront(edge.to)
					} else {
						deque.PushBack(edge.to)
					}
				}
			}
		}
		isProcessed[vertex] = true
		inQueue[vertex] = false
	}

	return distances
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 2)
	g.AddEdge(0, 3, 9)
	g.AddEdge(3, 4, 1)
	g.AddEdge(2, 4, 4)

	source := 0
	distances := g.Levit(source)
	fmt.Printf("Кратчайшие расстояния от вершины %d: %v\n", source, distances)
}
