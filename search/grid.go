package search

import (
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
	isBlock bool
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
		for x := range width {
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

// SetBlock sets the cell at the specified coordinates as a block.
func (g *Grid) SetBlock(x, y int) {
	cell := g.GetCell(x, y)
	if cell != nil {
		cell.isBlock = true
	}
}

// SetWeight sets the weight of the cell at the specified coordinates.
func (g *Grid) SetWeight(x, y, weight int) {
	cell := g.GetCell(x, y)
	if cell != nil {
		cell.Weight = weight
	}
}

// GetNeighbors returns the neighboring cells of the specified cell.
// If allowDiagonal is true, diagonal neighbors are included.
// Otherwise, only orthogonal neighbors are included.
func (g *Grid) GetNeighbors(cell *Cell, allowDiagonal bool) []*Cell {
	neighbors := []*Cell{}
	var directions []struct {
		dx int
		dy int
	}
	if allowDiagonal {
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
		nx, ny := cell.X+dir.dx, cell.Y+dir.dy
		if neighbor := g.GetCell(nx, ny); neighbor != nil && !neighbor.isBlock {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (g *Grid) ExecuteSearch(startCell, endCell *Cell, searchType string, allowDiagonal bool) ([]*Cell, error) {
	switch searchType {
	case "BFS":
		return BFS(g, startCell, endCell, allowDiagonal)
	case "DFS":
		return DFS(g, startCell, endCell, allowDiagonal)
	case "Dijkstra":
		return Dijkstra(g, startCell, endCell, allowDiagonal)
	case "AStar":
		return AStar(g, startCell, endCell, allowDiagonal)
	default:
		return nil, fmt.Errorf("unsupported search type: %s", searchType)
	}
}

// BFS performs a breadth-first search on the grid from startCell to endCell.
// It returns the path from startCell to endCell and the visited cells.
// If no path is found, it returns an error.
func BFS(g *Grid, startCell, endCell *Cell, allowDiagonal bool) ([]*Cell, []*Cell, error) {
	if startCell == nil || endCell == nil {
		return nil, fmt.Errorf("start or end cell is nil")
	}
	if startCell.isBlock || endCell.isBlock {
		return nil, fmt.Errorf("start or end cell is blocked")
	}
	if startCell == endCell {
		return []*Cell{startCell}, nil
	}
	frontier := []*Cell{startCell}
	visited := make(map[*Cell]bool)

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if visited[current] {
			continue
		}
		visited[current] = true
		neighbors := g.GetNeighbors(current, allowDiagonal)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				neighbor.Parent = current
				frontier = append(frontier, neighbor)
				if neighbor == endCell {
					return constructPath(neighbor), maptoslice(visited), nil
				}
			}
		}
	}
}

func maptoslice(m map[*Cell]bool) []*Cell {
	var slice []*Cell
	for k := range m {
		slice = append(slice, k)
	}
	return slice
}
