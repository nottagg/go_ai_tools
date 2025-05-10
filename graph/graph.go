package graph

import (
	"errors"

	"main.go/helpers"
)

//go interepretation of https://networkx.org/documentation/stable/_modules/networkx/classes/graph.html#Graph

type Graph[K comparable, V any] struct {
	Nodes      map[K]*Node[K, V]
	Edges      map[K]map[*Node[K, V]]int
	Name       string
	IsDirected bool
}
type Node[K comparable, V any] struct {
	Id          K
	Value       V
	X           int
	Y           int
	CurrentCost float64
}

// GraphType is a string that represents the type of graph
// It can be "grid", "directed", or "undirected"
// Returns a pointer to a new graph object
// Examples
// g := New("MyGraph", true))
func New[K comparable, V any](n string, isDirected bool) *Graph[K, V] {
	return &Graph[K, V]{
		Nodes:      make(map[K]*Node[K, V]),
		Edges:      make(map[K]map[*Node[K, V]]int),
		Name:       n,
		IsDirected: isDirected,
	}
}

// Returns True if n is a node, False otherwise
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// fmt.Println(g.Contains("A")) // true
// fmt.Println(g.Contains("B")) // false
func (g *Graph[K, V]) Contains(n K) bool {
	_, exists := g.Nodes[n]
	return exists
}

// Returns the number of nodes in the graph
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// g.AddNode("B")
// fmt.Println(g.Length()) // 2
func (g *Graph[K, V]) LengthNodes() int {
	return len(g.Nodes)
}

// Returns the number of edges in the graph
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// g.AddNode("B")
// g.AddEdge("A", "B", 1)
// fmt.Println(g.LengthEdges()) // 1
func (g *Graph[K, V]) LengthEdges() int {
	count := 0
	for _, edges := range g.Edges {
		count += len(edges)
	}
	return count
}

// Adds a node to the graph. Must include an x,y coordinate for
// calculating euclidean/manhattan distance in A*
// K is a unique indentifier while V is the value
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", 0, 0)
// g.AddNode("B", 1, 1)
// fmt.Println(g.Length()) // 2
func (g *Graph[K, V]) AddNode(k K, v V, x, y int) {
	if _, exists := g.Nodes[k]; !exists {
		g.Nodes[k] = &Node[K, V]{
			Value: v,
			X:     x,
			Y:     y,
			Id:    k,
		}
	}
}

// Adds multiple nodes to the graph.
// For using later on with converting 2d grid to graph
// Examples
// g := New("MyGraph", false)
// g.AddNodesFrom([]Node{"A", "B", "C"})
// fmt.Println(g.Length()) // 3
func (g *Graph[K, V]) AddNodesFrom(nodes []Node[K, V]) {
	for _, node := range nodes {
		g.AddNode(node.Id, node.Value, node.X, node.Y)
	}
}

// Removes a node from the graph
// If key is simply the value or coordinate pair pass that
// If the node is not in the graph, do nothing
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// g.RemoveNode("A")
// fmt.Println(g.Length()) // 1
func (g *Graph[K, V]) RemoveNode(k K) {
	delete(g.Nodes, k)
}

// Removes multiple nodes from the graph
// Takes an array of node keys
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// g.RemoveNodesFrom([]string{"A", "B"})
// fmt.Println(g.Length()) // 0
func (g *Graph[K, V]) RemoveNodesFrom(nodes []K) {
	for _, node := range nodes {
		g.RemoveNode(node)
	}
}

// Gets a node from the graph
// If the node is not in the graph, return nil
// If the node is in the graph, return a pointer to the node object
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// fmt.Println(g.GetNode("A")) // &{A A 0 0}
func (g *Graph[K, V]) GetNode(n K) *Node[K, V] {
	if _, exists := g.Nodes[n]; exists {
		return g.Nodes[n]
	}
	return nil
}

// Adds an edge between two nodes
// If the nodes are not in the graph, do nothing
// If the edge already exists, do nothing
// If the graph is directed, the edge is added in one direction
// If the graph is undirected, the edge is added in both directions
// weight could be uniform for a graph wtihout edge weights
// Examples
// g := New("MyGraph", false)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("B", "C", 2)
// fmt.Println(g.HasEdge("B", "C")) // true
func (g *Graph[K, V]) AddEdge(n1, n2 K, weight int) {
	if node1, exists := g.Nodes[n1]; exists {
		if node2, exists := g.Nodes[n2]; exists {
			if _, exists := g.Edges[n1]; !exists {
				g.Edges[n1] = make(map[*Node[K, V]]int)
			}
			g.Edges[n1][node2] = weight
			if !g.IsDirected {
				if _, exists := g.Edges[n2]; !exists {
					g.Edges[n2] = make(map[*Node[K, V]]int)
				}
				g.Edges[n2][node1] = weight
			}
		}
	}
}

// Gets the edges of a node
// If the edge does not exist, return false
// Examples
// g := New("MyGraph", false)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("B", "C", 2)
// fmt.Println(g.GetEdge("B")) // map[C:2]
func (g *Graph[K, V]) GetNodeEdges(n1 K) map[*Node[K, V]]int {
	if _, exists := g.Edges[n1]; exists {
		return g.Edges[n1]
	}
	return nil
}

// Returns the edge weight between two nodes
// If the edge does not exist, return -1
// Examples
// g := New("MyGraph", false)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("B", "C", 2)
// fmt.Println(g.GetEdgeWeight("B", "C")) // 2
func (g *Graph[K, V]) GetEdgeWeight(n1, n2 K) int {
	if _, exists := g.Edges[n1]; exists {
		if _, exists := g.Edges[n1][g.Nodes[n2]]; exists {
			return g.Edges[n1][g.Nodes[n2]]
		}
	}
	return -1
}

// Removes an edge between two nodes
// If the edge does not exist, do nothing
// If the graph is directed, the edge is removed in one direction
// If the graph is undirected, the edge is removed in both directions
// Examples
// g := New("MyGraph", false)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("B", "C", 2)
// g.RemoveEdge("B", "C")
// fmt.Println(g.HasEdge("B", "C")) // false
func (g *Graph[K, V]) RemoveEdge(n1, n2 K) {
	if _, exists := g.Edges[n1]; exists {
		delete(g.Edges[n1], g.Nodes[n2])
	}
	if !g.IsDirected {
		if _, exists := g.Edges[n2]; exists {
			delete(g.Edges[n2], g.Nodes[n1])
		}
	}
}

// Returns a new undirected graph created from a 2D matrix with weights
// The matrix should be a slice of slices of integers, where each integer
// represents the weight of the corresponding edge in the grid.
// The edge weight is the absolute difference between the values of the
// two nodes.
// A weight of -1 indicates a non-traversable node.
// Neighbors are determined based on the cells directly left, right, up, and down
// of the current cell.
// If allowDiagonal is true, the neighbors are determined based on the cells
// diagonally adjacent to the current cell as well.
// Examples
//
//	matrix := [][]int{
//		{0, 1, 0},
//		{1, 0, 1},
//		{0, 1, 0},
//	}
//
// g.NewGraphFromMatrix("GridGraph", matrix, true)
// fmt.Println(g.GetEdges()) // map[0:map[1:1] 1:map[0:1 2:1] 2:map[1:1]]
func NewGraphFromMatrix(n string, matrix [][]int, allowDiagonal bool) *Graph[helpers.Coordinate, int] {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])
	nodes := make([]*Node[helpers.Coordinate, int], rows*cols)
	for i := range rows {
		for j := range cols {
			nodes[i*cols+j] = &Node[helpers.Coordinate, int]{
				X: i,
				Y: j,
				Id: helpers.Coordinate{
					X: i,
					Y: j,
				},
				Value: matrix[i][j],
			}
		}
	}
	g := &Graph[helpers.Coordinate, int]{
		Nodes:      make(map[helpers.Coordinate]*Node[helpers.Coordinate, int]),
		Edges:      make(map[helpers.Coordinate]map[*Node[helpers.Coordinate, int]]int),
		Name:       n,
		IsDirected: false,
	}
	for _, node := range nodes {
		g.Nodes[node.Id] = node
	}

	//Create edges
	for i := range rows {
		for j := range cols {
			if matrix[i][j] == -1 {
				continue
			}
			nodeID := helpers.Coordinate{X: i, Y: j}
			if _, exists := g.Edges[nodeID]; !exists {
				g.Edges[nodeID] = make(map[*Node[helpers.Coordinate, int]]int)
			}
			directions := helpers.GetGridDirections(allowDiagonal)
			for _, direction := range directions {
				newX := i + direction[0]
				newY := j + direction[1]

				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					if matrix[newX][newY] == -1 {
						continue
					}
					neighborID := helpers.Coordinate{X: newX, Y: newY}
					edge_difference := helpers.IntegerAbsoluteValue(matrix[i][j] - matrix[newX][newY])
					g.Edges[nodeID][g.Nodes[neighborID]] = edge_difference
					if _, exists := g.Edges[neighborID]; !exists {
						g.Edges[neighborID] = make(map[*Node[helpers.Coordinate, int]]int)
					}
					g.Edges[neighborID][g.Nodes[nodeID]] = edge_difference
				}
			}
		}
	}
	return g
}

//TODO: Maybe enable more extensability by allowing the user to determine the edge weights as maybe diff(0,value)?

//Returns a path from start to end using BFS
// If no path exists, return nil
// If the start or end node is not in the graph, return nil
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("A", "B", 1)
// g.AddEdge("B", "C", 1)
// path, _, err := g.BFS("A", "C")
// if err != nil {

func (g *Graph[K, V]) BFS(start, end K) ([]K, int, error) {
	visited := make(map[K]bool)
	queue := []K{start}
	visited[start] = true
	parent := make(map[K]K)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			path := []*Node[K, V]{}
			for node != start {
				path = append(path, node)
				node = parent[node]
			}
			path = append(path, start)
			return path, len(visited), nil
		}

		for neighbor := range g.Edges[node] {
			if !visited[neighbor.Id] {
				visited[neighbor.Id] = true
				queue = append(queue, neighbor.Id)
				parent[neighbor.Id] = node
			}
		}
	}
	return nil, -1, errors.New("No path found")
}

func (g *Graph[K, V]) DFS(start, end Node[K, V]) ([]*Node[K, V], int, error) {
	visited := make(map[K]bool)
	stack := []*Node[K, V]{&start}
	visited[start.Id] = true
	parent := make(map[K]*Node[K, V])

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Id == end.Id {
			path := []*Node[K, V]{}
			for node.Id != start.Id {
				path = append(path, node)
				node = parent[node.Id]
			}
			path = append(path, node)
			return path, len(visited), nil
		}

		for neighbor := range g.Edges[node] {
			if !visited[neighbor.Id] {
				visited[neighbor.Id] = true
				stack = append(stack, neighbor.Id)
				parent[neighbor.Id] = node
			}
		}
	}

	return nil, 0, errors.New("No path found")
}

func (g *Graph[K, V]) Dijkstra(start, end Node[K, V]) ([]*Node[K, V], []*Node[K, V], error) {
	visited := make(map[K]bool)
	visited[start.Id] = true
	queue := helpers.NewPriorityQueue[*Node[K, V]]()
	parent := make(map[K]*Node[K, V])
	start.CurrentCost = 0
	queue.Enqueue(&start, 0)

	for queue.Len() > 0 {
		node := queue.Dequeue()

		if node.Id == end.Id {
			path := []*Node[K, V]{}
			for node.Id != start.Id {
				path = append(path, node)
				node = parent[node.Id]
			}
			return path, helpers.MapValuesToSlice(parent), nil
		}
		neighbors := g.Edges[node.Id]
		for neighbor := range neighbors {
			if visited[neighbor.Id] {
				continue
			}
			cost := g.Edges[node.Id][g.Nodes[neighbor.Id]]
			neighbor.CurrentCost = node.CurrentCost + float64(cost)
			parent[neighbor.Id] = node
			queue.Enqueue(neighbor, neighbor.CurrentCost)
		}
	}

	return nil, nil, nil
}

func (g *Graph[K, V]) AStar(start, end Node[K, V]) ([]*Node[K, V], []*Node[K, V], error) {
	visited := make(map[K]bool)
	visited[start.Id] = true
	queue := helpers.NewPriorityQueue[*Node[K, V]]()
	parent := make(map[K]*Node[K, V])
	start.CurrentCost = 0
	queue.Enqueue(&start, 0)

	for queue.Len() > 0 {
		node := queue.Dequeue()

		if node.Id == end.Id {
			path := []*Node[K, V]{}
			for node.Id != start.Id {
				path = append(path, node)
				node = parent[node.Id]
			}
			return path, helpers.MapValuesToSlice(parent), nil
		}
		neighbors := g.Edges[node.Id]
		for neighbor := range neighbors {
			if visited[neighbor.Id] {
				continue
			}
			cost := g.Edges[node.Id][g.Nodes[neighbor.Id]]
			neighbor.CurrentCost = node.CurrentCost + float64(cost)
			parent[neighbor.Id] = node

			heuristic := helpers.EuclideanDistance(neighbor.X, neighbor.Y, end.X, end.Y)
			queue.Enqueue(neighbor, neighbor.CurrentCost+heuristic)
		}
	}

	return nil, nil, nil
}
