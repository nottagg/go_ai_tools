package graph

import (
	"testing"
)

func TestGraphDirected(t *testing.T) {
	g := New[string, int]("testGraph", true)
	g.AddNode("A", 1)
	g.AddNode("B", 2)
	g.AddEdge("A", "B", 1)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", g.LengthEdges())
	}

	if g.GetEdgeWeight("A", "B") != 1 {
		t.Errorf("Expected edge weight 1, got %f", g.GetEdgeWeight("A", "B"))
	}
	if g.GetEdgeWeight("B", "A") != -1 {
		t.Errorf("Expected edge weight -1 for non-existent edge, got %f", g.GetEdgeWeight("B", "A"))
	}
}

func TestGraphUndirected(t *testing.T) {
	g := New[string, int]("testGraph", false)
	g.AddNode("A", 1)
	g.AddNode("B", 2)
	g.AddEdge("A", "B", 2)

	if len(g.Nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 2 {
		t.Errorf("Expected 2 edge, got %d", g.LengthEdges())
	}

	if g.GetEdgeWeight("A", "B") != 2 {
		t.Errorf("Expected edge weight 2, got %f", g.GetEdgeWeight("A", "B"))
	}
	if g.GetEdgeWeight("B", "A") != 2 {
		t.Errorf("Expected edge weight 2 for undirected edge, got %f", g.GetEdgeWeight("B", "A"))
	}
	if g.GetEdgeWeight("A", "C") != -1 {
		t.Errorf("Expected edge weight -1 for non-existent edge, got %f", g.GetEdgeWeight("A", "C"))
	}
}

func TestGraphFromMatrixOrtho(t *testing.T) {
	matrix := [][]float64{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, false)
	if g == nil {
		t.Fatal("Expected a graph, got nil")
	}
	if len(g.Nodes) != 9 {
		t.Errorf("Expected 9 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 20 {
		t.Errorf("Expected 20 edges, got %d", g.LengthEdges())
	}
}

func TestGraphFromMatrixDiag(t *testing.T) {
	matrix := [][]float64{
		{0, 1, -1},
		{1, 5, 3},
		{0, 1, 0},
	}
	g := NewGraphFromMatrix("grid", matrix, true)
	if g == nil {
		t.Fatal("Expected a graph, got nil")
	}
	if len(g.Nodes) != 9 {
		t.Errorf("Expected 9 nodes, got %d", g.LengthNodes())
	}

	if len(g.Edges) != 34 {
		t.Errorf("Expected 34 edges, got %d", g.LengthEdges())
	}
}

func TestSearch(t *testing.T) {
	matrix := [][]float64{
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
}
