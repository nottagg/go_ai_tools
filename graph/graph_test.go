package graph

import (
	"testing"

	"main.go/helpers"
)

func TestGraphDirected(t *testing.T) {
	g := New[string, int]("testGraph", true)
	g.AddNode("A", 1, 1, 1)
	g.AddNode("B", 2, 2, 2)
	g.AddEdge("A", "B", 1)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(g.Edges))
	}

	if len(g.GetNodeEdges("A")) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(g.GetNodeEdges("A")))
	}

	if len(g.GetNodeEdges("B")) != 0 {
		t.Errorf("Expected 0 edges, got %d", len(g.GetNodeEdges("B")))
	}

	if g.GetEdgeWeight("A", "B") != 1 {
		t.Errorf("Expected edge weight 1, got %d", g.GetEdgeWeight("A", "B"))
	}

	if g.GetNodeEdges("C") != nil {
		t.Errorf("Expected nil for non-existent node, got %v", g.GetNodeEdges("C"))
	}
	g.RemoveEdge("A", "B")
	if len(g.GetNodeEdges("A")) != 0 {
		t.Errorf("Expected no edges after removal, got %d", len(g.GetNodeEdges("A")))
	}
}

func TestGraphUndirected(t *testing.T) {
	g := New[string, int]("testGraph", false)
	g.AddNode("A", 1, 1, 1)
	g.AddNode("B", 2, 2, 2)
	g.AddEdge("A", "B", 2)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 2 {
		t.Errorf("Expected 2 edge, got %d", len(g.Edges))
	}

	if len(g.GetNodeEdges("A")) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(g.GetNodeEdges("A")))
	}

	if len(g.GetNodeEdges("B")) != 1 {
		t.Errorf("Expected 1 edges, got %d", len(g.GetNodeEdges("B")))
	}

	if g.GetEdgeWeight("A", "B") != 2 {
		t.Errorf("Expected edge weight 2, got %d", g.GetEdgeWeight("A", "B"))
	}

	if g.GetNodeEdges("C") != nil {
		t.Errorf("Expected nil for non-existent node, got %v", g.GetNodeEdges("C"))
	}
	g.RemoveEdge("A", "B")
	if len(g.GetNodeEdges("A")) != 0 {
		t.Errorf("Expected no edges after removal, got %d", len(g.GetNodeEdges("A")))
	}
}

func TestGraphFromMatrixOrtho(t *testing.T) {
	matrix := [][]int{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, false)
	if g == nil {
		t.Fatal("Expected a graph, got nil")
	}
	if len(g.Nodes) != 9 {
		t.Errorf("Expected 9 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 8 {
		t.Errorf("Expected 8 edges, got %d", len(g.Edges))
	}

	pair := helpers.Coordinate{X: 0, Y: 1}
	if len(g.GetNodeEdges(pair)) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(g.GetNodeEdges(pair)))
	}
	pair2 := helpers.Coordinate{X: 1, Y: 1}
	if g.GetEdgeWeight(pair2, pair) != 4 {
		t.Errorf("Expected edge weight 4, got %d", g.GetEdgeWeight(pair, pair2))
	}
	if g.GetEdgeWeight(pair, pair2) != 4 {
		t.Errorf("Expected edge weight 4, got %d", g.GetEdgeWeight(pair, pair2))
	}
	if g.GetNodeEdges(helpers.Coordinate{X: 0, Y: 2}) != nil {
		t.Errorf("Expected nil for -1 node, got %v", g.GetNodeEdges(helpers.Coordinate{X: 2, Y: 2}))
	}
}

func TestGraphFromMatrixDiag(t *testing.T) {
	matrix := [][]int{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, true)
	if g == nil {
		t.Fatal("Expected a graph, got nil")
	}
	if len(g.Nodes) != 9 {
		t.Errorf("Expected 9 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 8 {
		t.Errorf("Expected 8 edges, got %d", len(g.Edges))
	}

	pair := helpers.Coordinate{X: 0, Y: 1}
	if len(g.GetNodeEdges(pair)) != 4 {
		t.Errorf("Expected 4 edges, got %d", len(g.GetNodeEdges(pair)))
	}
	pair2 := helpers.Coordinate{X: 1, Y: 1}
	if g.GetEdgeWeight(pair2, pair) != 4 {
		t.Errorf("Expected edge weight 4, got %d", g.GetEdgeWeight(pair, pair2))
	}
	if g.GetEdgeWeight(pair, pair2) != 4 {
		t.Errorf("Expected edge weight 4, got %d", g.GetEdgeWeight(pair, pair2))
	}
	if g.GetNodeEdges(helpers.Coordinate{X: 0, Y: 2}) != nil {
		t.Errorf("Expected nil for -1 node, got %v", g.GetNodeEdges(helpers.Coordinate{X: 2, Y: 2}))
	}
}

func TestAStarSearch(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, -1, -1},
		{1, 0, -1, 3, -1},
		{2, -1, 0, 1, 4},
		{-1, 3, 1, 0, 2},
		{-1, -1, 4, 2, 0},
	}

	// Create a graph from the matrix
	g := NewGraphFromMatrix("testGraph", matrix, false)
	if g == nil {
		t.Fatal("Expected a graph, got nil")
	}

	// Ensure the graph has at least 15 nodes
	if len(g.Nodes) < 15 {
		t.Fatalf("Expected at least 15 nodes, got %d", len(g.Nodes))
	}

	// Define start and goal nodes
	start := g.GetNode(helpers.Coordinate{X: 0, Y: 0})
	goal := g.GetNode(helpers.Coordinate{X: 4, Y: 4})

	// Run A* search
	path, cost, error := g.AStar(*start, *goal)

	if error != nil {
		t.Fatalf("Expected no error, got %v", error)
	}
	// Validate the results
	if path == nil {
		t.Fatalf("Expected a valid path, got nil")
	}

	if cost <= 0 {
		t.Errorf("Expected a positive cost, got %d", cost)
	}

	// Print the path for debugging
	t.Logf("Path: %v, Cost: %d", path, cost)

	// Additional checks
	if len(path) < 2 {
		t.Errorf("Expected a path with at least 2 nodes, got %d", len(path))
	}

	// Ensure the path starts and ends at the correct nodes
	if path[0] != start {
		t.Errorf("Expected path to start at %v, got %v", start, path[0])
	}
	if path[len(path)-1] != goal {
		t.Errorf("Expected path to end at %v, got %v", goal, path[len(path)-1])
	}
}
