package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
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

	if len(g.GetEdge("A")) != 1 {
		t.Errorf("Expected edge weight 1, got %d", len(g.GetEdge("A")))
	}

	if len(g.GetEdge("B")) != 0 {
		t.Errorf("Expected edge weight 0, got %d", len(g.GetEdge("B")))
	}

	if g.GetEdge("C") != nil {
		t.Errorf("Expected nil for non-existent node, got %v", g.GetEdge("C"))
	}
}
