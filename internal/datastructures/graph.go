package datastructures

import "fmt"

type Graph[T comparable] map[T]Set[T]

// Creates a new directed graph
//
// edge: [source, dest]
func NewGraph[T comparable](edges [][2]T, vertices []T) Graph[T] {
	graph := Graph[T]{}

	for _, v := range vertices {
		graph[v] = NewSet[T]()
	}

	graph.AddEdges(edges...)

	return graph
}

// Creates a new undirected graph
//
// edges: [source, dest] and the other way around
func NewGraphUndirected[T comparable](edges [][2]T, vertices []T) Graph[T] {
	graph := Graph[T]{}

	for _, v := range vertices {
		graph[v] = NewSet[T]()
	}

	graph.AddEdgesUndirected(edges...)

	return graph
}

// Adds one or more directed edges to the graph
//
// edge: [source, dest]
func (g Graph[T]) AddEdges(edges ...[2]T) {
	for _, e := range edges {
		if _, ok := g[e[0]]; !ok {
			g[e[0]] = NewSet(e[1])
		} else {
			g[e[0]].Add(e[1])
		}
	}
}

// Adds one or more directed edges to the graph
//
// edge: [source, dest]
func (g Graph[T]) AddEdgesUndirected(edges ...[2]T) {
	for _, e := range edges {
		if _, ok := g[e[0]]; !ok {
			g[e[0]] = NewSet[T]()
		}

		if _, ok := g[e[1]]; !ok {
			g[e[1]] = NewSet[T]()
		}

		g[e[0]].Add(e[1])
		g[e[1]].Add(e[0])
	}
}

// returns topoligical order for provided vertices
func (g Graph[T]) TopologicalOrder(vertices []T) ([]T, error) {
	var order []T
	// count of directed edges (from node pointing to node)
	inDegree := map[T]int{}

	for _, from := range vertices {
		for _, to := range g[from].Members() {
			inDegree[to] += 1
		}
	}

	queue := NewDeque([]T{})
	for _, node := range vertices {
		if inDegree[node] == 0 {
			queue.PushBack(node)
		}
	}

	if queue.Len() == 0 {
		return order, fmt.Errorf("no vertices with in degree of 0")
	}

	for queue.Len() > 0 {
		node, ok := queue.PopFront()
		if ok {
			order = append(order, node)

			for _, neighbor := range g[node].Members() {
				inDegree[neighbor] -= 1
				if inDegree[neighbor] == 0 {
					queue.PushBack(neighbor)
				}
			}
		}
	}

	return order, nil
}
