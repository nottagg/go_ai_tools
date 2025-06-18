package graph

import (
	"errors"

	"main.go/helpers"
)

//go interpretation of https://networkx.org/documentation/stable/_modules/networkx/classes/graph.html#Graph

type Graph[K comparable, V any] struct {
	Nodes      map[K]*Node[K, V]
	Edges      map[K]map[K]float64
	Name       string
	IsDirected bool
}

type Node[K comparable, V any] struct {
	Key         K
	Value       V
	CurrentCost float64
	Parent      *Node[K, V]
}

// GraphType is a string that represents the type of graph
// It can be "grid", "directed", or "undirected"
// Returns a pointer to a new graph object
// Examples
// g := New("MyGraph", true))
func New[K comparable, V any](name string, isDirected bool) *Graph[K, V] {
	return &Graph[K, V]{
		Nodes:      make(map[K]*Node[K, V], 0),
		Edges:      make(map[K]map[K]float64),
		Name:       name,
		IsDirected: isDirected,
	}
}

// Resets the current cost and parent of all nodes in the graph
// This is useful for algorithms that need to reinitialize the graph state
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.ResetNodes()
// fmt.Println(g.Nodes["A"].CurrentCost) // Output: 0
// fmt.Println(g.Nodes["A"].Parent) // Output: <nil>
func (g *Graph[K, V]) ResetNodes() {
	for _, node := range g.Nodes {
		node.CurrentCost = 0
		node.Parent = nil
	}
}

// Returns the number of nodes in the graph
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// fmt.Println(g.LengthNodes()) // Output: 2
func (g *Graph[K, V]) LengthNodes() int {
	return len(g.Nodes)
}

// Checks if the graph contains a node with the given key
// Returns true if the node exists, false otherwise
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// fmt.Println(g.ContainsNode("A")) // Output: true
// fmt.Println(g.ContainsNode("B")) // Output: false
func (g *Graph[K, V]) ContainsNode(k K) bool {
	if _, exists := g.Nodes[k]; exists {
		return true
	}
	return false
}

// Creates a new node with the given key and value
// Adds the node to the graph's node list
// If the node already exists, it is not added again
// Returns a pointer to the newly created node
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// fmt.Println(g.ContainsNode("A")) // Output: true
func (g *Graph[K, V]) AddNode(k K, v V) {
	node := &Node[K, V]{
		Key:         k,
		Value:       v,
		CurrentCost: 0,
		Parent:      nil,
	}
	if _, exists := g.Nodes[k]; !exists {
		g.Nodes[k] = node
	}
}

// Removes a node from the graph
// If the node is not in the graph, do nothing
// Also removes all edges associated with the node
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// g.RemoveNode("A")
// fmt.Println(g.ContainsNode("A")) // Output: false
// fmt.Println(g.ContainsNode("B")) // Output: true
// fmt.Println(g.LengthEdges()) // Output: 0
func (g *Graph[K, V]) RemoveNode(k K) {
	if g.ContainsNode(k) {
		delete(g.Nodes, k)
		delete(g.Edges, k)
		for e := range g.Edges {
			delete(g.Edges[e], k)
		}
	}
}

// Returns the number of edges in the graph
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// fmt.Println(g.LengthEdges()) // Output: 1
func (g *Graph[K, V]) LengthEdges() int {
	count := 0
	for _, edges := range g.Edges {
		count += len(edges)
	}
	return count
}

// Checks if the graph contains an edge between two nodes
// If the graph is directed, it checks for the edge in one direction
// If the graph is undirected, it checks for the edge in both directions
// Returns true if the edge exists, false otherwise
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// fmt.Println(g.ContainsEdge("A", "B")) // Output: true
// fmt.Println(g.ContainsEdge("B", "A")) // Output: false (For directed graph)
func (g *Graph[K, V]) ContainsEdge(k1, k2 K) bool {
	if _, exists := g.Edges[k1]; exists {
		if _, exists := g.Edges[k1][k2]; exists {
			return true
		}
	}
	if !g.IsDirected {
		if _, exists := g.Edges[k2]; exists {
			if _, exists := g.Edges[k2][k1]; exists {
				return true
			}
		}
	}
	return false
}

// Adds an edge between two nodes
// If the nodes are not in the graph, do nothing
// If the edge already exists, do nothing
// If the graph is directed, the edge is added in one direction
// If the graph is undirected, the edge is added in both directions
// weight could be uniform for a graph wtihout edge weights
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// fmt.Println(g.ContainsEdge("A", "B")) // Output: true
func (g *Graph[K, V]) AddEdge(k1, k2 K, weight float64) {
	if _, exists := g.Nodes[k1]; exists {
		if _, exists := g.Nodes[k2]; exists {
			if _, exists := g.Edges[k1]; !exists {
				g.Edges[k1] = make(map[K]float64)
			}
			g.Edges[k1][k2] = weight
			if !g.IsDirected {
				if _, exists := g.Edges[k2]; !exists {
					g.Edges[k2] = make(map[K]float64)
				}
				g.Edges[k2][k1] = weight
			}
		}
	}
}

// Returns the edge weight between two nodes
// If the edge does not exist, return -1
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// fmt.Println(g.GetEdgeWeight("A", "B")) // Output: 1
func (g *Graph[K, V]) GetEdgeWeight(k1, k2 K) float64 {
	if _, exists := g.Edges[k1]; exists {
		if _, exists := g.Edges[k1][k2]; exists {
			return g.Edges[k1][k2]
		}
	}
	return -1
}

// Removes an edge between two nodes
// If the edge does not exist, do nothing
// If the graph is directed, the edge is removed in one direction
// If the graph is undirected, the edge is removed in both directions
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddEdge("A", "B", 1)
// g.RemoveEdge("A", "B")
// fmt.Println(g.ContainsEdge("A", "B")) // Output: false
func (g *Graph[K, V]) RemoveEdge(k1, k2 K) {
	if _, exists := g.Edges[k1]; exists {
		delete(g.Edges[k1], k2)
	}
	if !g.IsDirected {
		if _, exists := g.Edges[k2]; exists {
			delete(g.Edges[k2], k1)
		}
	}
}

// Returns a new undirected graph created from a 2D matrix with weights
// The matrix should be a slice of slices of integers, where each integer
// represents the weight of the corresponding edge in the grid.
// The graph is generated directed but every node is connected with the cost
// Being the value of the cell in the matrix
// A weight of -1 indicates a non-traversable node.
// Neighbors are determined based on the cells directly left, right, up, and down
// of the current cell.
// If allowDiagonal is true, the neighbors are determined based on the cells
// diagonally adjacent to the current cell as well.
// Examples
//
//	matrix := [][]int{
//		{0, 1, -1},
//		{1, 5, 3},
//		{0, 1, 0},
//	}
//
// g := NewGraphFromMatrix("grid", matrix, false)
// fmt.Println(g.LengthNodes()) // Output: 9
// fmt.Println(g.LengthEdges()) // Output: 18
// fmt.Println(g.ContainsEdge(helpers.Coordinate{X: 0, Y:
func NewGraphFromMatrix(n string, matrix [][]float64, allowDiagonal bool) *Graph[helpers.Coordinate, float64] {
	g := New[helpers.Coordinate, float64](n, false)
	directions := helpers.GetGridDirections(allowDiagonal)

	for i := range matrix {
		for j := range matrix[i] {
			k1 := helpers.Coordinate{X: float64(i), Y: float64(j)}
			g.AddNode(k1, matrix[i][j])
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[0]) && matrix[ni][nj] != -1 {
					k2 := helpers.Coordinate{X: float64(ni), Y: float64(nj)}
					if _, exists := g.Nodes[k2]; !exists {
						g.AddNode(k2, matrix[ni][nj])
					}
					// Don't add edge for value of -1
					if matrix[ni][nj] == -1 || matrix[i][j] == -1 {
						continue
					}
					g.AddEdge(k1, k2, matrix[ni][nj])
					g.AddEdge(k2, k1, matrix[i][j])
				}
			}
		}
	}
	return g
}

// Returns a path from start to end using BFS
// If no path exists, return nil
// If the start or end node is not in the graph, return nil
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddNode("C", 3)
// g.AddEdge("A", "B", 1)
// g.AddEdge("B", "C", 1)
// path, visited, err := g.BFS("A", "C")
//
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//
//	    fmt.Println("Path:", path) // Output: Path: [A B C]
//	    fmt.Println("Visited:", visited) // Output: Visited: [A B C]
//	}
func (g *Graph[K, V]) BFS(start, end K) ([]K, []K, error) {

	if start == end {
		return []K{start}, []K{start}, nil
	} else if !g.ContainsNode(start) || !g.ContainsNode(end) {
		return nil, nil, errors.New("start or end node not in graph")
	}

	visited := make(map[K]bool)
	visited[start] = true
	queue := []K{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for neighbor := range g.Edges[node] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			g.Nodes[neighbor].Parent = g.Nodes[node]
			if neighbor == end {
				path := g.constructPath(neighbor)
				return path, helpers.MapKeysToSlice(visited), nil
			}
			queue = append(queue, neighbor)
		}
	}
	return nil, nil, errors.New("no path found")
}

// Returns a path from start to end using DFS
// If no path exists, return nil
// If the start or end node is not in the graph, return nil
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddNode("C", 3)
// g.AddEdge("A", "B", 1)
// g.AddEdge("B", "C", 1)
// path, visited, err := g.DFS("A", "C")
//
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//	    fmt.Println("Path:", path) // Output: Path: [A B C]
//	    fmt.Println("Visited:", visited) // Output: Visited: [A B C]
//	}
func (g *Graph[K, V]) DFS(start, end K) ([]K, []K, error) {
	if start == end {
		return []K{start}, []K{start}, nil
	} else if !g.ContainsNode(start) || !g.ContainsNode(end) {
		return nil, nil, errors.New("start or end node not in graph")
	}

	visited := make(map[K]bool)
	visited[start] = true
	stack := []K{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for neighbor := range g.Edges[node] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			g.Nodes[neighbor].Parent = g.Nodes[node]
			if neighbor == end {
				path := g.constructPath(neighbor)
				return path, helpers.MapKeysToSlice(visited), nil
			}
			stack = append(stack, neighbor)
		}
	}
	return nil, nil, errors.New("no path found")
}

// Returns a path from start to end using Dijkstra's algorithm (UCS)
// If no path exists, return nil
// If the start or end node is not in the graph, return nil
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddNode("C", 3)
// g.AddEdge("A", "B", 2)
// g.AddEdge("B", "C", 2)
// g.AddEdge("A", "C", 2)
// path, visited, err := g.Dijkstra("A", "C")
//
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//	    fmt.Println("Path:", path) // Output: Path: [A C]
//	    fmt.Println("Visited:", visited) // Output: Visited: [A B C]
//	}
func (g *Graph[K, V]) Dijkstra(start, end K) ([]K, []K, error) {
	if start == end {
		return []K{start}, []K{start}, nil
	} else if !g.ContainsNode(start) || !g.ContainsNode(end) {
		return nil, nil, errors.New("start or end node not in graph")
	}
	visited := make(map[K]bool)
	pq := make(helpers.PriorityQueue[K], 0)
	pq.PushItem(start, 0)
	g.Nodes[start].CurrentCost = 0

	for pq.Len() > 0 {
		node := pq.PopItem()
		if visited[node] {
			continue
		}
		visited[node] = true
		if node == end {
			path := g.constructPath(node)
			return path, helpers.MapKeysToSlice(visited), nil
		}
		for neighbor := range g.Edges[node] {
			if visited[neighbor] {
				continue
			}
			newCost := g.Nodes[node].CurrentCost + g.GetEdgeWeight(node, neighbor)
			if g.Nodes[neighbor].Parent == nil || newCost < g.Nodes[neighbor].CurrentCost {
				g.Nodes[neighbor].CurrentCost = newCost
				g.Nodes[neighbor].Parent = g.Nodes[node]
				pq.PushItem(neighbor, newCost)
			}
		}
	}
	return nil, nil, errors.New("no path found")
}

// Returns a path from start to end using A* algorithm
// If no path exists, return nil
// If the start or end node is not in the graph, return nil
// The heuristic function should return the estimated cost from node a to node b
// Examples
// g := New("MyGraph", true)
// g.AddNode("A", 1)
// g.AddNode("B", 2)
// g.AddNode("C", 3)
// g.AddEdge("A", "B", 2)
// g.AddEdge("B", "C", 2)
// g.AddEdge("A", "C", 2)
// heuristic := func(a, b helpers.Coordinate) float64 {
//     return 1 // Example heuristic, replace with actual heuristic logic
// }
// path, visited, err := g.AStar("A", "C", heuristic)
//
//	if err != nil {
//	    fmt.Println(err)
//	} else {
//	    fmt.Println("Path:", path) // Output: Path: [A C]
//	    fmt.Println("Visited:", visited) // Output: Visited: [A B C]
//	}

func (g *Graph[K, V]) AStar(start, end K, heuristic func(a, b K) float64) ([]K, []K, error) {
	if start == end {
		return []K{start}, []K{start}, nil
	} else if !g.ContainsNode(start) || !g.ContainsNode(end) {
		return nil, nil, errors.New("start or end node not in graph")
	}
	visited := make(map[K]bool)
	pq := make(helpers.PriorityQueue[K], 0)
	pq.PushItem(start, 0)
	g.Nodes[start].CurrentCost = 0

	for pq.Len() > 0 {
		node := pq.PopItem()
		if visited[node] {
			continue
		}
		visited[node] = true
		if node == end {
			path := g.constructPath(node)
			return path, helpers.MapKeysToSlice(visited), nil
		}
		for neighbor := range g.Edges[node] {
			if visited[neighbor] {
				continue
			}
			newCost := g.Nodes[node].CurrentCost + g.GetEdgeWeight(node, neighbor) + heuristic(neighbor, end)
			if g.Nodes[neighbor].Parent == nil || newCost < g.Nodes[neighbor].CurrentCost {
				g.Nodes[neighbor].CurrentCost = newCost
				g.Nodes[neighbor].Parent = g.Nodes[node]
				pq.PushItem(neighbor, newCost)
			}
		}
	}
	return nil, nil, errors.New("no path found")
}

func (g *Graph[K, V]) constructPath(k K) []K {
	path := make([]K, 0)
	for g.Nodes[k].Parent != nil {
		path = append([]K{k}, path...)
		k = g.Nodes[k].Parent.Key
	}
	path = append([]K{k}, path...)
	return path
}
