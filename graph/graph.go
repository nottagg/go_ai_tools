package graph

import "main.go/helpers"

//go interepretation of https://networkx.org/documentation/stable/_modules/networkx/classes/graph.html#Graph

type Graph[K comparable, V any] struct {
	nodes     map[K]*Node[K, V]
	edges     map[K]map[K]int
	name      string
	graphType GraphType
}
type Node[K comparable, V any] struct {
	id    K
	value V
	x     int
	y     int
}

type GraphType int

const (
	Directed GraphType = iota
	Undirected
)

// For grid type, graph uniqueness is based on x,y coordinates
// For directed and undirected, graph uniqueness is based on the value of the node
var graphTypes = map[string]GraphType{
	"directed":   Directed,
	"undirected": Undirected,
}

// GraphType is a string that represents the type of graph
// It can be "grid", "directed", or "undirected"
// Returns a pointer to a new graph object
// Examples
// g := New("MyGraph", "grid"))
func New[K comparable, V any](n string, gt GraphType) *Graph[K, V] {
	return &Graph[K, V]{
		nodes:     make(map[K]*Node[K, V]),
		edges:     make(map[K]map[K]int),
		name:      n,
		graphType: gt,
	}
}

// String identifier of the graph
func (g *Graph[K, V]) Name(s string) {
	g.name = s
}

// Returns True if n is a node, False otherwise
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// fmt.Println(g.Contains("A")) // true
// fmt.Println(g.Contains("B")) // false
func (g *Graph[K, V]) Contains(n K) bool {
	_, exists := g.nodes[n]
	return exists
}

// Returns the number of nodes in the graph
// Examples
// g := New("MyGraph", false)
// g.AddNode("A")
// g.AddNode("B")
// fmt.Println(g.Length()) // 2
func (g *Graph[K, V]) LengthNodes() int {
	return len(g.nodes)
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
	for _, edges := range g.edges {
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
	if _, exists := g.nodes[k]; !exists {
		g.nodes[k] = &Node[K, V]{
			value: v,
			x:     x,
			y:     y,
			id:    k,
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
		g.AddNode(node.id, node.value, node.x, node.y)
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
	delete(g.nodes, k)
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
	if _, exists := g.nodes[n]; exists {
		return g.nodes[n]
	}
	return nil
}

// Returns the map of nodes
// The map is keyed by the node id
// The value is a pointer to the node object
// Examples
// g := New("MyGraph", false)
// g.AddNode("A", "A",0, 0)
// g.AddNode("B", "B",1, 1)
// fmt.Println(g.GetNodes()) // map[A:{VALUE} B:{VALUE}]
func (g *Graph[K, V]) GetNodes() map[K]*Node[K, V] {
	return g.nodes
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
	if _, exists := g.nodes[n1]; exists {
		if _, exists := g.nodes[n2]; exists {
			if _, exists := g.edges[n1]; !exists {
				g.edges[n1] = make(map[K]int)
			}
			g.edges[n1][n2] = weight
			if g.graphType == Undirected {
				if _, exists := g.edges[n2]; !exists {
					g.edges[n2] = make(map[K]int)
				}
				g.edges[n2][n1] = weight
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
func (g *Graph[K, V]) GetEdge(n1 K) map[K]int {
	if _, exists := g.edges[n1]; exists {
		return g.edges[n1]
	}
	return nil
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
	if _, exists := g.edges[n1]; exists {
		delete(g.edges[n1], n2)
	}
	if g.graphType == Undirected {
		if _, exists := g.edges[n2]; exists {
			delete(g.edges[n2], n1)
		}
	}
}

// Returns the edges in the graph
// The edges are returned as a map of maps
// The outer map is keyed by the node id
// The inner map is keyed by the node id of the neighbor
// The value is the weight of the edge
// Examples
// g := New("MyGraph", false)
// g.AddNode("B", "B",1, 1)
// g.AddNode("C", "C",2, 2)
// g.AddEdge("B", "C", 2)
// fmt.Println(g.GetEdges()) // map[B:map[C:2]]
func (g *Graph[K, V]) GetEdges() map[K]map[K]int {
	return g.edges
}

// Returns the graph type
// The graph type is returned as a string
// The graph type can be "grid", "directed", or "undirected"
// Examples
// g := New("MyGraph", false)
// fmt.Println(g.GetGraphType()) // "grid"
func (g *Graph[K, V]) GetGraphType() GraphType {
	return g.graphType
}

// Returns a new undirected graph created from a 2D matrix with weights
// The matrix should be a slice of slices of integers, where each integer
// represents the weight of the corresponding edge in the grid.
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
				x: i,
				y: j,
				id: helpers.Coordinate{
					X: i,
					Y: j,
				},
				value: matrix[i][j],
			}
		}
	}
	g := &Graph[helpers.Coordinate, int]{
		nodes:     make(map[helpers.Coordinate]*Node[helpers.Coordinate, int]),
		edges:     make(map[helpers.Coordinate]map[helpers.Coordinate]int),
		name:      n,
		graphType: Undirected,
	}
	for _, node := range nodes {
		g.nodes[node.id] = node
	}

	//Create edges
	for i := range rows {
		for j := range cols {
			if matrix[i][j] == -1 {
				continue
			}
			nodeID := helpers.Coordinate{X: i, Y: j}
			if _, exists := g.edges[nodeID]; !exists {
				g.edges[nodeID] = make(map[helpers.Coordinate]int)
			}
			directions := helpers.GetGridDirections(allowDiagonal)
			for _, direction := range directions {
				newX := i + direction[0]
				newY := j + direction[1]
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
					neighborID := helpers.Coordinate{X: newX, Y: newY}
					if matrix[newX][newY] != -1 {
						g.edges[nodeID][neighborID] = matrix[newX][newY]
					}
				}
			}
		}
	}

	return g
}
