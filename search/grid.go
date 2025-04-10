package search

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// Grid represents a 2D grid of cells.
type Grid struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

// Cell represents a cell in the grid.
type Cell struct {
	X       int
	Y       int
	Weight  int
	Visited bool
	Parent  *Cell
}

// NewGrid creates a new grid with the specified width and height.
// If random is true, the grid will be filled with random weights.
// Otherwise, all cells will have a weight of 1.
func NewGrid(width, height, max_weight int, random bool) *Grid {
	cells := make([][]*Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]*Cell, width)
		for x := 0; x < width; x++ {
			weight := 1
			if random {
				weight = rand.Intn(max_weight) // Random weight between 0 and max_weight
			}
			cells[y][x] = &Cell{
				X:       x,
				Y:       y,
				Weight:  weight,
				Visited: false,
			}
		}
	}
	return &Grid{Width: width, Height: height, Cells: cells}
}

// GetCell returns the cell at the specified coordinates.
func (g *Grid) GetCell(x, y int) *Cell {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return nil
	}
	return g.Cells[y][x]
}

// SetWeight sets the weight of the cell at the specified coordinates.
func (g *Grid) SetWeight(x, y, weight int) {
	cell := g.GetCell(x, y)
	if cell != nil {
		cell.Weight = weight
	}
}

// GetNeighbors returns the neighboring cells of the specified cell.
func (g *Grid) GetNeighbors(cell *Cell) []*Cell {
	neighbors := []*Cell{}
	directions := []struct {
		dx int
		dy int
	}{
		{-1, 0}, // left
		{1, 0},  // right
		{0, -1}, // up
		{0, 1},  // down
	}
	for _, dir := range directions {
		nx, ny := cell.X+dir.dx, cell.Y+dir.dy
		if neighbor := g.GetCell(nx, ny); neighbor != nil {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// GenerateRandom returns a grid with random weights given a width and height
func GenerateRandomGrid(x, y int) *Grid {
	grid := NewGrid(x, y)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			grid.SetWeight(j, i, rand.Intn(10))
		}
	}
}

// A* algorithm to find the shortest path from start to goal
func (g *Grid) AStar(startX, startY, goalX, goalY int) ([]*Cell, error) {
	start := g.GetCell(startX, startY)
	goal := g.GetCell(goalX, goalY)
	if start == nil || goal == nil {
		return nil, fmt.Errorf("invalid start or goal cell")
	}

	openSet := &PriorityQueue{}
	heap.Push(openSet, &Item{Cell: start, Priority: 0})

	cameFrom := make(map[*Cell]*Cell)
	gScore := make(map[*Cell]int)
	gScore[start] = 0

	fScore := make(map[*Cell]int)
	fScore[start] = heuristic(start, goal)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Item).Cell

		if current == goal {
			return reconstructPath(cameFrom, current), nil
		}

		for _, neighbor := range g.GetNeighbors(current) {
			if neighbor.Visited {
				continue
			}
			tentativeGScore := gScore[current] + neighbor.Weight
			if _, ok := gScore[neighbor]; !ok || tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + heuristic(neighbor, goal)
				if !contains(openSet, neighbor) {
					heap.Push(openSet, &Item{Cell: neighbor, Priority: fScore[neighbor]})
				}
			}
		}
		current.Visited = true
	}

	return nil, fmt.Errorf("no path found")
}

// heuristic function to estimate the distance between two cells
func heuristic(a, b *Cell) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// reconstructPath reconstructs the path from the start to the goal.
func reconstructPath(cameFrom map[*Cell]*Cell, current *Cell) []*Cell {
	path := []*Cell{current}
	for current != nil {
		path = append(path, current)
		current = cameFrom[current]
	}
	reverse(path)
	return path
}

// reverse reverses the order of the cells in the path.
func reverse(path []*Cell) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
