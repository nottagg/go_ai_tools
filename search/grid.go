package search

import (
	"fmt"
	"math/rand"

	"main.go/helpers"
)

// GridNode represents a GridNode in the grid.
// Implements the Node interface from the helpers package.
type GridNode struct {
	X      int
	Y      int
	Weight int
	Parent *GridNode
}

// GetParent returns the parent of the GridNode.
func (c *GridNode) GetParent() helpers.Node[*Grid] {
	return c.Parent
}

// SetParent sets the parent of the GridNode.
func (c *GridNode) SetParent(parent helpers.Node[*Grid]) {
	c.Parent = parent.(*GridNode)
}

// GetNeighbors returns the neighboring GridNodes of the specified GridNode.
// Otherwise, only orthogonal neighbors are included.
// Non-traversable GridNodes (weight -1) are excluded from the neighbors.
func (c *GridNode) GetNeighbors(g *Grid) []helpers.Node[*Grid] {
	neighbors := []helpers.Node[*Grid]{}
	var directions []struct {
		dx int
		dy int
	}
	if g.AllowDiagonal {
		directions = []struct {
			dx int
			dy int
		}{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}
	} else {
		directions = []struct {
			dx int
			dy int
		}{
			{-1, 0},
			{0, -1},
			{0, 1},
			{1, 0},
		}
	}
	for _, dir := range directions {
		nx, ny := c.X+dir.dx, c.Y+dir.dy
		if neighbor := g.GetGridNode(nx, ny); neighbor != nil && neighbor.Weight != -1 {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// Grid represents a 2D grid of GridNodes.
// It contains the width and height of the grid, a 2D slice of GridNodes,
// and a boolean indicating whether diagonal movement is allowed.
type Grid struct {
	Width         int
	Height        int
	GridNodes     [][]*GridNode
	AllowDiagonal bool
}

// NewGridFromMatrix creates a new grid from a 2D matrix of weights.
// The matrix should be a slice of slices of integers, where each integer
// represents the weight of the corresponding GridNode in the grid.
// A weight of -1 indicates a non-traversable GridNode.
func NewGrid(matrix [][]int) *Grid {
	width := len(matrix[0])
	height := len(matrix)
	GridNodes := make([][]*GridNode, height)
	for y := 0; y < height; y++ {
		GridNodes[y] = make([]*GridNode, width)
		for x := 0; x < width; x++ {
			GridNodes[y][x] = &GridNode{
				X:      x,
				Y:      y,
				Weight: matrix[y][x],
			}
		}
	}
	return &Grid{Width: width, Height: height, GridNodes: GridNodes}
}

// RandomizeWeights takes a grid and produces random weights between -1 and max_weights
func (g *Grid) RandomizeWeights(max_weight int) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			weight := rand.Intn(max_weight) - 1
			g.GridNodes[y][x].Weight = weight
		}
	}
}

// GetGridNode returns the GridNode at the specified coordinates.
func (g *Grid) GetGridNode(x, y int) *GridNode {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return nil
	}
	return g.GridNodes[y][x]
}

// SetWeight sets the weight of the GridNode at the specified coordinates.
// A weight of -1 indicates a non-traversable GridNode.
func (g *Grid) SetWeight(x, y, weight int) {
	GridNode := g.GetGridNode(x, y)
	if GridNode != nil {
		GridNode.Weight = weight
	}
}

// ExecuteSearch executes the specified search algorithm on the grid.
// It takes the start and end GridNodes, the search type as a string "BFS", "DFS", "Dijkstra", or "AStar",
// and a boolean indicating whether diagonal movement is allowed.
// It returns the path from startGridNode to endGridNode and the visited GridNodes.
// If no path is found, it returns an error.
func (g *Grid) ExecuteSearch(startGridNode, endGridNode *GridNode, searchType string) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	switch searchType {
	case "BFS":
		return BFS(g, startGridNode, endGridNode)
	case "DFS":
		return DFS(g, startGridNode, endGridNode)
	case "Dijkstra":
		return Dijkstra(g, startGridNode, endGridNode)
	case "AStar":
		return AStar(g, startGridNode, endGridNode)
	default:
		return nil, nil, fmt.Errorf("unsupported search type: %s", searchType)
	}
}

// BFS performs a breadth-first search on the grid from startGridNode to endGridNode.
// It returns the path from startGridNode to endGridNode and the visited GridNodes.
// If no path is found, it returns an error.
func BFS(g *Grid, startGridNode, endGridNode *GridNode) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	if startGridNode == nil || endGridNode == nil {
		return nil, nil, fmt.Errorf("start or end GridNode is nil")
	}
	if startGridNode.Weight == -1 || endGridNode.Weight == -1 {
		return nil, nil, fmt.Errorf("start or end GridNode is blocked")
	}
	if startGridNode == endGridNode {
		return []helpers.Node[*Grid]{startGridNode}, nil, nil
	}
	frontier := []helpers.Node[*Grid]{startGridNode}
	visited := make(map[helpers.Node[*Grid]]bool)

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		neighbors := current.GetNeighbors(g)
		visited[current] = true
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				neighbor.SetParent(current)
				frontier = append(frontier, neighbor)
				if neighbor == endGridNode {
					path := helpers.BuildReturnPath(neighbor)
					visisted_GridNodes := helpers.MapKeysToSlice(visited)

					return path, visisted_GridNodes, nil
				}
			}
		}
	}
	return nil, helpers.MapKeysToSlice(visited), fmt.Errorf("no path found")
}

func DFS(g *Grid, startGridNode, endGridNode *GridNode) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	if startGridNode == nil || endGridNode == nil {
		return nil, nil, fmt.Errorf("start or end GridNode is nil")
	}
	if startGridNode.Weight == -1 || endGridNode.Weight == -1 {
		return nil, nil, fmt.Errorf("start or end GridNode is blocked")
	}
	if startGridNode == endGridNode {
		return []helpers.Node[*Grid]{startGridNode}, nil, nil
	}
	frontier := []helpers.Node[*Grid]{startGridNode}
	visited := make(map[helpers.Node[*Grid]]bool)

	for len(frontier) > 0 {
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		neighbors := current.GetNeighbors(g)
		visited[current] = true
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				neighbor.SetParent(current)
				frontier = append(frontier, neighbor)
				if neighbor == endGridNode {
					path := helpers.BuildReturnPath(neighbor)
					visisted_GridNodes := helpers.MapKeysToSlice(visited)

					return path, visisted_GridNodes, nil
				}
			}
		}
	}
	return nil, helpers.MapKeysToSlice(visited), fmt.Errorf("no path found")
}

func Dijkstra(g *Grid, startGridNode, endGridNode *GridNode) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	return nil, nil, fmt.Errorf("Dijkstra search not implemented")
}

func AStar(g *Grid, startGridNode, endGridNode *GridNode) ([]helpers.Node[*Grid], []helpers.Node[*Grid], error) {
	return nil, nil, fmt.Errorf("A* search not implemented")
}
