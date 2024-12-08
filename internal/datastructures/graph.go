package datastructures

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

// Adds one or more directed edges to the graph
//
// edge: [source, dest]
func (g Graph[T]) AddEdges(edges ...[2]T) {
	for _, e := range edges {
		_, ok := g[e[0]]

		if !ok {
			g[e[0]] = NewSet(e[1])
		} else {
			g[e[0]].Add(e[1])
		}
	}
}

func (g Graph[T]) TopologicalOrder(vertices []T) []T {
	var order []T
	// count of directed edges (from node pointing to node)
	inDegree := map[T]int{}

	for _, from := range vertices {
		for _, to := range g[from].Members() {
			inDegree[to] += 1
		}
	}

	queue := NewDeque[T]([]T{})
	for _, node := range vertices {
		if inDegree[node] == 0 {
			queue.PushBack(node)
		}
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

	return order
}
