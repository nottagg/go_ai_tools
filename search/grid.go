package search

import (
	"fmt"

	"main.go/helpers"
)

// node represents a node in the grid.
// Implements the Node interface from the helpers package.
type Node struct {
	X      int
	Y      int
	Neighbors *[]Neighbor
	Parent *Node
	Graph *Graph
}

type Neighbor struct {
	cost int
	node *Node
}

// Grid represents a 2D grid of nodes.
// It contains the width and height of the grid, a 2D slice of nodes,
// and a boolean indicating whether diagonal movement is allowed.
// AllowDiagonal is only used for grid like graphs
type Graph struct {
	Nodes			[]*Node
	Edges
	AllowDiagonal	bool
}

// NewGraphFromMatrix creates a new graph from a 2D matrix of weights.
// The matrix should be a slice of slices of integers, where each integer
// represents the weight of the corresponding edge in the grid.
// A weight of -1 indicates a non-traversable node.
func NewGraphFromMatrix(matrix [][]int, allowDiagonal bool) *Graph {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])
	nodes := make([]*Node, rows*cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			nodes[i*cols+j] = &Node{
				X:      i,
				Y:      j,
				Neighbors: &[]Neighbor{},
			}
		}
	}

	g := &Graph{
		Nodes: nodes,
		AllowDiagonal: allowDiagonal,
	}

	for _, node := range nodes {
		node.Graph = g
	}

	return g
}

// GetNeighbors returns the neighbors of the node in the grid.
// It returns a slice of nodes that are adjacent to the current node.
// The neighbors are determined based on the grid's dimensions and the
// AllowDiagonal property of the graph.
func (g *Graph) GetNeighbors(node *Node) []*Node {
	if node == nil {
		return nil
	}
	directions := helpers.GetGridDirections(g.AllowDiagonal)
	neighbors := make([]*Node, 0)

	for _, direction := range directions {
		newX := node.X + direction[0]
		newY := node.Y + direction[1]

		if newX >= 0 && newX < len(g.Nodes) && newY >= 0 && newY < len(g.Nodes[0]) {
			neighbor := g.Nodes[newX*len(g.Nodes[0])+newY]
			if neighbor.Weight != -1 {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}


// ExecuteSearch executes the specified search algorithm on the grid.
// It takes the start and end nodes, the search type as a string "BFS", "DFS", "Dijkstra", or "AStar",
// and a boolean indicating whether diagonal movement is allowed.
// It returns the path from startnode to endnode and the visited nodes.
// If no path is found, it returns an error.
func (g *Grid) ExecuteSearch(startnode, endnode *node, searchType string) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	switch searchType {
	case "BFS":
		return BFS(g, startnode, endnode)
	case "DFS":
		return DFS(g, startnode, endnode)
	case "Dijkstra":
		return Dijkstra(g, startnode, endnode)
	case "AStar":
		return AStar(g, startnode, endnode)
	default:
		return nil, nil, fmt.Errorf("unsupported search type: %s", searchType)
	}
}

// BFS performs a breadth-first search on the grid from startnode to endnode.
// It returns the path from startnode to endnode and the visited nodes.
// If no path is found, it returns an error.
func BFS(g *Grid, startnode, endnode *node) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	if startnode == nil || endnode == nil {
		return nil, nil, fmt.Errorf("start or end node is nil")
	}
	if startnode.Weight == -1 || endnode.Weight == -1 {
		return nil, nil, fmt.Errorf("start or end node is blocked")
	}
	if startnode == endnode {
		return []helpers.Node[*Grid]{startnode}, nil, nil
	}
	frontier := []helpers.Node[*Grid]{startnode}
	visited := make(map[helpers.Node[*Grid]]bool)

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1GetNeighbors(g)

		neighbors := current.Neighbors
		visited[current] = true
		for _, neiighbors {
			if !visited[neighbor] {
				neighbor.Parent = current
				if neighbor == endnode {
					return helpers.BuildReturnPath((neighbor)), helpers.MapKeysToSlice((visited)), nil
				}
				frontier = append(frontier, neighbor)
			}
		}
	}
	return nil, helpers.MapKeysToSlice(visited), fmt.Errorf("no path found")
}

func DFS(g *Grid, startnode, endnode *node) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	if startnode == nil || endnode == nil {
		return nil, nil, fmt.Errorf("start or end node is nil")
	}
	if startnode.Weight == -1 || endnode.Weight == -1 {
		return nil, nil, fmt.Errorf("start or end node is blocked")
	}
	if startnode == endnode {
		return []helpers.Node[*Grid]{startnode}, nil, nil
	}
	frontier := []helpers.Node[*Grid]{startnode}
	visited := make(map[helpers.Node[*Grid]]bool)

	for len(frontier) > 0 {
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		neighbors := current.GetNeighbors(g)
		visited[current] = true
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				neighbor.SetParent(current)
				if neighbor == endnode {
					return helpers.BuildReturnPath((neighbor)), helpers.MapKeysToSlice((visited)), nil
				}
				frontier = append(frontier, neighbor)
			}
		}
	}
	return nil, helpers.MapKeysToSlice(visited), fmt.Errorf("no path found")
}

func Dijkstra(g *Grid, startnode, endnode *node) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	return nil, nil, fmt.Errorf("Dijkstra search not implemented")
}

func AStar(g *Grid, startnode, endnode *node) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	return nil, nil, fmt.Errorf("A* search not implemented")
}
