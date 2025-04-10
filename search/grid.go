package search

import (
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
